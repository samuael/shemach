package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
)

// AdminRepo ...
type AdminRepo struct {
	DB *pgxpool.Pool
}

// NewAdminRepo returning admin repo implementing al  the
// interfaces specified by the cruds
func NewAdminRepo(db *pgxpool.Pool) admin.IAdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (adminr *AdminRepo) AdminByEmail(ctx context.Context) (*model.Admin, error) {
	email := ctx.Value("email").(string)
	admin := &model.Admin{}
	if er := adminr.DB.QueryRow(ctx, "SELECT id,fullname,email,password,imgurl,superadmin,created_at FROM admins WHERE email=$1", email).
		Scan(&(admin.ID), &(admin.Fullname), &(admin.Email), &(admin.Password), &(admin.Imgurl), &(admin.Superadmin), &(admin.CreatedAtRef)); er == nil {
		date := &model.Date{}
		if er := adminr.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", admin.CreatedAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
			return admin, er
		}
		admin.CreatedAt = date
		return admin, nil
	} else {
		return nil, er
	}
}

// ChangePassword ...
func (adminr *AdminRepo) ChangePassword(ctx context.Context) (bool, error) {
	id := ctx.Value("user_id").(uint)
	password := ctx.Value("password").(string)
	cmd, err := adminr.DB.Exec(ctx, "UPDATE admins SET password=$1 WHERE id=$2", password, id)
	if err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (adminr *AdminRepo) DeleteAccountByEmail(ctx context.Context) (bool, int) {
	email := ctx.Value("email").(string)
	var success bool
	if err := adminr.DB.QueryRow(ctx, "select deleteadminbyemail from deleteAdminByEmail($1)", email).Scan(&success); err == nil {
		return success, state.DT_STATUS_OK
	} else {
		log.Println(err.Error())
	}
	return false, state.DT_STATUS_DBQUERY_ERROR
}

// CreateAdmin(ctx context.Context) (*model.Admin, error)
func (adminr *AdminRepo) CreateAdmin(ctx context.Context) (*model.Admin, error) {
	admin := ctx.Value("admin").(*model.Admin)
	row := adminr.DB.QueryRow(ctx, "select id,fullname,email,password,superadmin,imgurl,created_at from createAdmin( $1,$2,$3,$4,$5,$6,$7,$8,$9,$10, $11 );",
		admin.Fullname, admin.Email, admin.Password, admin.Superadmin, admin.Imgurl, admin.CreatedAt.Years, admin.CreatedAt.Months, admin.CreatedAt.Days, admin.CreatedAt.Hours, admin.CreatedAt.Minutes, admin.CreatedAt.Seconds)
	// duplicate key value violates unique constraint "admins_email_key"

	if row != nil {
		println(string(helper.MarshalThis(row)))
		era := row.Scan(&(admin.ID), &(admin.Fullname), &(admin.Email), &(admin.Password), &(admin.Superadmin), &(admin.Imgurl), &(admin.CreatedAtRef))
		return admin, era
	} else {
		return nil, errors.New("internal error while creating an admin")
	}
}

// AdminByID(ctx context.Context) (*model.Admin, error)
func (adminr *AdminRepo) AdminByID(ctx context.Context) (*model.Admin, error) {
	ID := ctx.Value("user_id").(uint64)
	admin := &model.Admin{}
	if er := adminr.DB.QueryRow(ctx, "SELECT id,fullname,email,password,imgurl,superadmin,created_at FROM admins WHERE id=$1", ID).
		Scan(&(admin.ID), &(admin.Fullname), &(admin.Email), &(admin.Password), &(admin.Imgurl), &(admin.Superadmin), &(admin.CreatedAtRef)); er == nil {
		date := &model.Date{}
		if er := adminr.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", admin.CreatedAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
			return admin, er
		}
		admin.CreatedAt = date
		return admin, nil
	} else {
		return nil, er
	}
}

func (adminr *AdminRepo) UpdateAdmin(ctx context.Context) (*model.Admin, error) {
	admin := ctx.Value("admin").(*model.Admin)
	if cmd, er := adminr.DB.Exec(ctx, "UPDATE admins SET email=$1 , password=$2 , imgurl=$3 , fullname=$4 ,superadmin=$5, WHERE id=$6",
		admin.Email, admin.Password, admin.Imgurl, admin.Fullname, admin.Superadmin, admin.ID,
	); cmd.Update() && er == nil {
		return admin, nil
	} else {
		return nil, errors.New("update was not successful")
	}
}

// GetImageUrl
func (adminr *AdminRepo) GetImageUrl(ctx context.Context) (string, error) {
	session := ctx.Value("session").(*model.Session)
	imgurl := ""
	if err := adminr.DB.QueryRow(ctx, "SELECT imgurl from admins where id=$1", session.ID).Scan(&imgurl); err != nil {
		return imgurl, err
	}
	return imgurl, nil
}

func (adminr *AdminRepo) ChangeImageUrl(ctx context.Context) error {
	imgurl := ctx.Value("image_url").(string)
	userid := ctx.Value("user_id").(uint64)
	if cmd, err := adminr.DB.Exec(ctx, "UPDATE admins SET imgurl=$1  WHERE id=$2", imgurl, userid); cmd.RowsAffected() == 0 || err != nil {
		return errors.New("internal server error while updating the image url")
	}
	return nil
}

// DeleteProfilePicture(ctx context.Context) error
func (adminr *AdminRepo) DeleteProfilePicture(ctx context.Context) error {
	session := ctx.Value("session").(*model.Session)
	if cmd, err := adminr.DB.Exec(ctx, "UPDATE admins SET imgurl=$1 WHERE id=$2", "", session.ID); cmd.RowsAffected() == 0 || err != nil {
		return errors.New("internal server error while updating the image url")
	}
	return nil
}

func (adminr *AdminRepo) DeleteAccountByID(ctx context.Context) (bool, int) {
	adminID := ctx.Value("admin_id").(uint64)
	var success bool
	if err := adminr.DB.QueryRow(ctx, "select deleteadmin from deleteadmin($1,$2)", "", int(adminID)).Scan(&(success)); err == nil {
		println(adminID, "  ", success)
		return success, state.DT_STATUS_OK
	} else {
		println(err.Error(), "  ---  ")
		return false, state.DT_STATUS_DBQUERY_ERROR
	}
}
func (adminr *AdminRepo) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	admins := []*model.Admin{}
	rows, gerr := adminr.DB.Query(ctx, "SELECT id,fullname,email,password,imgurl,superadmin,created_at FROM admins")
	if rows == nil || gerr != nil {
		if gerr != nil {
			println(gerr.Error())
			return admins, gerr
		}
		return admins, fmt.Errorf("invalid row data")
	}
	for rows.Next() {
		admin := &model.Admin{}
		er := rows.Scan(&(admin.ID), &(admin.Fullname), &(admin.Email), &(admin.Password), &(admin.Imgurl), &(admin.Superadmin), &(admin.CreatedAtRef))
		if er != nil {
			continue
		}
		date := &model.Date{}
		if er := adminr.DB.QueryRow(ctx, "SELECT id, year,month,day,hour,minute,second from eth_date where id=$1", admin.CreatedAtRef).
			Scan(&(date.ID), &(date.Years), &(date.Months), &(date.Days), &(date.Hours), &(date.Minutes), &(date.Seconds)); er != nil {
			continue
		}
		admin.CreatedAt = date
		admins = append(admins, admin)
	}
	if len(admins) == 0 {
		return admins, fmt.Errorf("no record found")
	}
	return admins, nil
}
