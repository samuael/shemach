package pgx_storage

import (
	"context"

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
	idinterface := ctx.Value("user_id")
	var id uint64
	id = 0
	if idinterface != nil {
		id = idinterface.(uint64)
	}
	email := ctx.Value("user_email").(string)
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
