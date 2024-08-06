package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/superadmin"
)

type SuperadminRepo struct {
	DB *pgxpool.Pool
}

func NewSuperadminRepo(conn *pgxpool.Pool) superadmin.ISuperadminRepo {
	return &SuperadminRepo{
		DB: conn,
	}
}

func (repo *SuperadminRepo) CheckTheExistanceOfSuperadmin(ctx context.Context) int {
	email := ctx.Value("user_email").(string)
	var status int
	er := repo.DB.QueryRow(ctx, "SELECT * FROM checkTheExistanceOfUser($1)", email).Scan(&status)
	if er != nil {
		return -1
	}
	return status
}

func (repo *SuperadminRepo) GetSuperadminByEmail(ctx context.Context) (*model.Superadmin, int, error) {
	email := ctx.Value("user_email").(string)
	superadmin := &model.Superadmin{}
	er := repo.DB.QueryRow(ctx, "select id ,firstname ,lastname ,phone ,email ,created_at ,password,registered_admins,registered_products from superadmin where email=$1", email).
		Scan(&(superadmin.ID), &(superadmin.Firstname), &(superadmin.Lastname), &(superadmin.Phone), &(superadmin.Email), &(superadmin.CreatedAtUnix),
			&(superadmin.Password), &(superadmin.RegisteredAdmins), &(superadmin.RegisteredProducts))
	if er != nil {
		return superadmin, state.STATUS_DBQUERY_ERROR, er
	}
	return superadmin, state.STATUS_OK, nil
}

func (repo *SuperadminRepo) GetSuperadminByID(ctx context.Context, id int) (*model.Superadmin, error) {
	superadmin := &model.Superadmin{}
	er := repo.DB.QueryRow(ctx, "select id ,firstname ,lastname ,phone ,email ,created_at ,password,registered_admins,registered_products from superadmin where id=$1", id).
		Scan(&(superadmin.ID), &(superadmin.Firstname), &(superadmin.Lastname), &(superadmin.Phone), &(superadmin.Email), &(superadmin.CreatedAtUnix),
			&(superadmin.Password), &(superadmin.RegisteredAdmins), &(superadmin.RegisteredProducts))
	if er != nil {
		return superadmin, er
	}
	return superadmin, nil
}

func (repo *SuperadminRepo) GetAllSuperadmins(ctx context.Context) ([]*model.Superadmin, error) {
	superadmins := []*model.Superadmin{}
	rows, er := repo.DB.Query(ctx, "select id ,firstname ,lastname ,phone ,email ,created_at ,password,registered_admins,registered_products from superadmin")
	if er != nil {
		return nil, er
	}
	for rows.Next() {
		superadmin := &model.Superadmin{}
		er = rows.Scan(&(superadmin.ID), &(superadmin.Firstname), &(superadmin.Lastname), &(superadmin.Phone), &(superadmin.Email), &(superadmin.CreatedAtUnix),
			&(superadmin.Password), &(superadmin.RegisteredAdmins), &(superadmin.RegisteredProducts))
		if er != nil {
			println(er.Error())
			continue
		}
		superadmins = append(superadmins, superadmin)
	}
	return superadmins, er
}
