package pgx_storage

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) user.IUserRepo {
	return &UserRepo{
		DB: db,
	}
}

// GetUserByEmailOrID ...
func (repo *UserRepo) GetUserByEmailOrID(ctx context.Context) (user *model.User, role int, status int, ers error) {
	id, ok := ctx.Value("user_id").(uint64)
	if !ok {
		id = 0
	}
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		if id == 0 {
			return nil, 0, state.STATUS_RECORD_NOT_FOUND, errors.New("missing important parameter")
		}
		email = ""
	}
	role = 0
	er := repo.DB.QueryRow(ctx, "select * from getTheRoleOfUserByIdOrEmail( $1,$2);", id, email).Scan(&role)
	if er != nil {
		println(er.Error())
		return nil, 0, state.STATUS_DBQUERY_ERROR, er
	}
	user = &model.User{}
	er = repo.DB.QueryRow(ctx, "select id,firstname,lastname,phone,email,imageurl,created_at,password,lang from users where id=$1 or email=$2", id, email).
		Scan(&(user.ID), &(user.Firstname), &(user.Lastname), &(user.Phone), &(user.Email), &(user.Imgurl), &(user.CreatedAt), &(user.Password), &(user.Lang))
	if er != nil {
		return nil, role, state.STATUS_DBQUERY_ERROR, er
	}
	return user, role, state.STATUS_OK, nil
}

// UpdatePassword ...
func (repo *UserRepo) UpdatePassword(ctx context.Context) error {
	userid := ctx.Value("user_id").(uint64)
	password := ctx.Value("new_password").(string)
	uc, er := repo.DB.Exec(ctx, "Update users set password=$1 where id=$2", password, userid)
	if er != nil {
		return er
	} else if uc.RowsAffected() == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// GetImageUrl ...  |
func (repo *UserRepo) GetImageUrl(ctx context.Context) string {
	userId := ctx.Value("user_id").(uint64)
	var imgurl string
	repo.DB.QueryRow(ctx, "select imageurl from users where id=$1", userId).Scan(&imgurl)
	return imgurl
}
func (repo *UserRepo) ChangeImageUrl(ctx context.Context) error {
	userId := ctx.Value("user_id").(uint64)
	imgurl := ctx.Value("image_url").(string)

	if er := repo.DB.QueryRow(ctx, "update users set imageurl=$1 where id=$2 returning id", imgurl, userId).Scan(&userId); er != nil {
		return er
	}
	return nil
}

// DeletePedingEmailConfirmation
func (repo *UserRepo) DeletePendingEmailConfirmation(timestamp uint64) error {
	deleted := 0
	er := repo.DB.QueryRow(context.Background(), "delete from emailInConfirmation where created_at<$1", timestamp).Scan(&deleted)
	if er != nil || deleted == 0 {
		// if er != nil {
		// 	println("ERROR:  ", er.Error())
		// }
		return errors.New("no row deleted ")
	}
	return nil
}

// SaveEmailConfirmation
func (repo *UserRepo) SaveEmailConfirmation(ctx context.Context, ec *model.EmailConfirmation) (int, error) {
	println(ec.UserID, ec.Email, ec.IsNewAccount, ec.OldEmail)
	did := 0
	er := repo.DB.QueryRow(ctx, "select * from insertEmailInConfirmation($1, $2,$3,$4)", ec.UserID, ec.Email, ec.IsNewAccount, ec.OldEmail).Scan(&did)
	if er != nil {
		if strings.Contains(er.Error(), "duplicate key value violates unique constrain") {
			return state.STATUS_RECORD_NOT_FOUND, errors.New("email with same account is in confirmation")
		} else if strings.Contains(er.Error(), "failed to connect to ") {
			return state.STATUS_DBQUERY_ERROR, errors.New("internal database problem, please try again later")
		}
		return state.STATUS_DBQUERY_ERROR, er
	} else if did == -1 {
		return did, errors.New("user with this id doesn't exist")
	} else if did == -2 {
		return did, errors.New("invalid")
	}
	ec.ID = uint64(did)
	return did, nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, user *model.User) (int, error) {
	er := repo.DB.QueryRow(ctx, "UPDATE users set firstname=$1,lastname=$2,phone=$3, email=$4, lang=$5 where id=$6 returning id", user.Firstname, user.Lastname, user.Phone, user.Email, user.Lang, user.ID).Scan(&(user.ID))
	if er != nil {
		return state.STATUS_DBQUERY_ERROR, er
	} else if user.ID == 0 {
		return state.STATUS_NO_RECORD_UPDATED, errors.New("user instance was not updated")
	}
	return state.STATUS_OK, nil
}

// GetUserByPhone(ctx context.Context, phone string) (*model.User, error)
func (repo *UserRepo) GetUserByPhone(ctx context.Context, phone string) (user *model.User, role int, status int, er error) {
	role = 0
	er = repo.DB.QueryRow(ctx, "select * from getTheRoleOfUserByPhone( $1 );", phone).Scan(&role)
	if er != nil {
		return nil, 0, state.STATUS_DBQUERY_ERROR, er
	}
	user = &model.User{}
	er = repo.DB.QueryRow(ctx, "select id,firstname,lastname,phone,email,imageurl,created_at,password,lang from users where id=$1 or phone=$2", phone).
		Scan(&(user.ID), &(user.Firstname), &(user.Lastname), &(user.Phone), &(user.Email), &(user.Imgurl), &(user.CreatedAt), &(user.Password), &(user.Lang))
	if er != nil {
		return nil, role, state.STATUS_DBQUERY_ERROR, er
	}
	return user, role, state.STATUS_OK, nil
}

func (repo *UserRepo) RegisterTempoCXP(ctx context.Context, tempo *model.TempoCXP) error {
	er := repo.DB.QueryRow(ctx, `insert into tempo_cxp(phone ,confirmation ,role ,created_at)  values($1,$2,$3,$4) returning id`, tempo.Phone, tempo.Confirmation, tempo.Role, tempo.CreatedAt).Scan(&(tempo.ID))
	if er != nil {
		println(er.Error())
		return er
	}
	return nil
}
