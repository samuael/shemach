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
// in this method if the account is new then not only the email in confirmation will be deleted
// but also the user account in the email will also be deleted so this functionality will be implemented withe the trigger i am going to write.
func (repo *UserRepo) DeletePendingEmailConfirmation(timestamp uint64) error {
	deleted := 0
	er := repo.DB.QueryRow(context.Background(), "delete from emailInConfirmation where created_at<$1  and is_new_account=$2", timestamp, false).Scan(&deleted)
	if er != nil || deleted == 0 {
		return errors.New("no row deleted ")
	}
	var ids []uint64

	rows, err := repo.DB.Query(context.Background(), "select userid from emailInConfirmation where created_at<$1  and is_new_account=$2", timestamp, true)
	if err != nil {
		return err
	}
	for rows.Next() {
		var id uint64
		r := rows.Scan(&id)
		if r != nil {
			continue
		}
		ids = append(ids, id)
	}

	var count uint64
	er = repo.DB.QueryRow(context.Background(), "select * from deleteUnconfirmedAdmins($1)", ids).Scan(&count)
	if er != nil || count == 0 {
		return er
	}
	return nil
}

// SaveEmailConfirmation
func (repo *UserRepo) SaveEmailConfirmation(ctx context.Context, ec *model.EmailConfirmation) (int, error) {
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

// UpdateUser ...
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
	er = repo.DB.QueryRow(ctx, "select id,firstname,lastname,phone,email,imageurl,created_at,password,lang from users where phone=$1", phone).
		Scan(&(user.ID), &(user.Firstname), &(user.Lastname), &(user.Phone), &(user.Email), &(user.Imgurl), &(user.CreatedAt), &(user.Password), &(user.Lang))
	if er != nil {
		return nil, role, state.STATUS_DBQUERY_ERROR, er
	}
	return user, role, state.STATUS_OK, nil
}

// RegisterTempoCXP ...
func (repo *UserRepo) RegisterTempoCXP(ctx context.Context, tempo *model.TempoCXP) error {
	er := repo.DB.QueryRow(ctx, `insert into tempo_cxp(phone ,confirmation ,role ,created_at)  values($1,$2,$3,$4) returning id`, tempo.Phone, tempo.Confirmation, tempo.Role, tempo.CreatedAt).Scan(&(tempo.ID))
	if er != nil {
		return er
	}
	return nil
}

// GetTempoCXP ...
func (repo *UserRepo) GetTempoCXP(ctx context.Context, phone string, response model.TempoCXP) error {
	ere := repo.DB.QueryRow(ctx, "select id, phone ,confirmation ,role ,created_at from tempo_cxp where phone=$1", phone).Scan(
		&(response.ID), &(response.Phone), &(response.Confirmation), &(response.Role), &(response.CreatedAt),
	)
	return ere
}

// RemoveTempoCXP ...
func (repo *UserRepo) RemoveTempoCXP(ctx context.Context, phone string) error {
	count := 0
	ere := repo.DB.QueryRow(ctx, "delete from tempo_cxp where phone=$1 returning id", phone).Scan(&count)
	if count <= 0 {
		return errors.New("now row is deleted")
	}
	return ere
}

func (repo *UserRepo) RemoveExpiredCXPConfirmations(timestamp uint64) (count int, er error) {
	phones := []string{}
	recs, er := repo.DB.Query(context.Background(), "select phone from tempo_cxp where created_at<$1", timestamp)
	if er != nil {
		return -1, er
	}
	for recs.Next() {
		var phone string
		er := recs.Scan(&phone)
		if er == nil {
			phones = append(phones, phone)
		}
	}
	var deletedCount int
	erf := repo.DB.QueryRow(context.Background(), "select * from deleteExpredCXPAccount($1)", phones).Scan(&deletedCount)
	if erf != nil {
		return 0, erf
	}
	return count, nil
}

func (repo *UserRepo) ConfirmUserEmailUpdate(ctx context.Context, id uint64, newemail, oldemail string) error {
	confirm, er := repo.GetEmailInConfirmationByID(ctx, id)
	if er != nil {
		return er
	}
	if !(confirm.IsNewAccount) && confirm.Email == newemail && confirm.OldEmail == oldemail {
		uc, er := repo.DB.Exec(ctx, "update users set email=$1 where email=$2 returning id", newemail, oldemail)
		if uc.RowsAffected() == 0 || er != nil {
			if er != nil {
				return er
			}
			return errors.New("no rows affected")
		}
	} else if (confirm.Email == newemail) && (confirm.IsNewAccount) {
		raff, erf := repo.DB.Exec(ctx, "delete from emailInConfirmation where id=$1", id)
		if erf != nil || raff.RowsAffected() == 0 {
			if erf != nil {
				return erf
			}
			return errors.New("no row was deleted")
		}
	}
	return errors.New("unauthorized")
}

func (repo *UserRepo) GetEmailInConfirmationByID(ctx context.Context, id uint64) (*model.EmailConfirmation, error) {
	confirm := &model.EmailConfirmation{}
	er := repo.DB.QueryRow(ctx, "select  id ,userid ,new_email, is_new_account, old_email,created_at from emailInConfirmation where id=$1", id).
		Scan(&(confirm.ID), &(confirm.UserID), &(confirm.Email), &(confirm.IsNewAccount), &(confirm.OldEmail), &(confirm.CreatedAt))
	if er != nil {
		return nil, er
	}
	return confirm, nil
}
