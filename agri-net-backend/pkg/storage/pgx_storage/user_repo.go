package pgx_storage

import (
	"context"
	"errors"

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
func (repo *UserRepo) DeletePendingEmailConfirmation(timestamp uint64) error {
	deleted := 0
	er := repo.DB.QueryRow(context.Background(), "delete from emailInConfirmation where createdat >$1", timestamp).Scan(&deleted)
	if er != nil || deleted == 0 {
		return errors.New("no row deleted ")
	}
	return nil
}

// SaveEmailConfirmation
func (repo *UserRepo) SaveEmailConfirmation(ctx context.Context, ec *model.EmailConfirmation) error {
	er := repo.DB.QueryRow(ctx, "select * from insertEmailInConfirmation($1, $2,$3,$4)", ec.UserID, ec.Email, ec.IsNewAccount, ec.OldEmail).Scan(&ec.ID)
	return er
}
