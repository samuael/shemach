package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
)

type InfoadminRepo struct {
	DB *pgxpool.Pool
}

func NewInfoadminRepo(db *pgxpool.Pool) infoadmin.IInfoadminRepo {
	return &InfoadminRepo{
		DB: db,
	}
}

func (repo *InfoadminRepo) GetInfoadminByEmail(ctx context.Context) (*model.Infoadmin, error) {
	email := ctx.Value("infoadmin_email").(string)
	admin := &model.Infoadmin{}
	er := repo.DB.QueryRow(ctx, `select id,firstname ,lastname ,phone ,email ,imageurl ,created_at ,password, messages_count, created_by from infoadmin where email=$1`, email).Scan(
		&(admin.ID), &(admin.Firstname), &(admin.Lastname), &(admin.Phone), &(admin.Email), &(admin.Imgurl), &(admin.CreatedAt), &(admin.Password), &(admin.BroadcastedMessagesCount), &(admin.Createdby),
	)
	if er != nil {
		println(er.Error())
		return admin, er
	}
	return admin, nil
}

func (repo *InfoadminRepo) CreateInfoadmin(ctx context.Context) (*model.Infoadmin, error) {
	infoadmin := ctx.Value("info_admin").(*model.Infoadmin)
	er := repo.DB.QueryRow(ctx, "insert into infoadmin( firstname, lastname, phone, email, password, messages_count, created_by) values($1,$2,$3,$4,$5,$6,$7) returning id",
		infoadmin.Firstname, infoadmin.Lastname, infoadmin.Phone, infoadmin.Email, infoadmin.Password, infoadmin.BroadcastedMessagesCount, infoadmin.Createdby,
	).Scan(&(infoadmin.ID))
	return infoadmin, er
}
func (repo *InfoadminRepo) GetInfoadmins(ctx context.Context) ([]*model.Infoadmin, error) {
	infoadmins := []*model.Infoadmin{}
	row, er := repo.DB.Query(ctx, "select id,firstname,lastname,phone,email,imageurl,created_at,lang,messages_count,created_by from infoadmin")
	if er != nil {
		return nil, er
	}
	for row.Next() {
		infoadmin := &model.Infoadmin{}
		er := row.Scan(&(infoadmin.ID), &(infoadmin.Firstname), &(infoadmin.Lastname), &(infoadmin.Phone), &(infoadmin.Email),
			&(infoadmin.Imgurl), &(infoadmin.CreatedAt), &(infoadmin.Lang), &(infoadmin.BroadcastedMessagesCount), &(infoadmin.Createdby))
		if er != nil {
			continue
		}
		infoadmins = append(infoadmins, infoadmin)
	}
	return infoadmins, nil
}

func (repo *InfoadminRepo) DeleteInfoadminByID(ctx context.Context) (int, error) {
	infoadminID := ctx.Value("infoadmin_id").(uint64)
	status := 0
	er := repo.DB.QueryRow(ctx, "select * from deleteinfoadminById($1)", infoadminID).Scan(&status)
	if er != nil {
		return status, er
	}
	return status, nil
}

func (repo *InfoadminRepo) GetInfoadminByID(ctx context.Context, id uint64) (*model.Infoadmin, error) {
	admin := &model.Infoadmin{}
	er := repo.DB.QueryRow(ctx, `select id,firstname ,lastname ,phone ,email ,imageurl ,created_at ,password, messages_count, created_by from infoadmin where id=$1`, id).Scan(
		&(admin.ID), &(admin.Firstname), &(admin.Lastname), &(admin.Phone), &(admin.Email), &(admin.Imgurl), &(admin.CreatedAt), &(admin.Password), &(admin.BroadcastedMessagesCount), &(admin.Createdby),
	)
	if er != nil {
		println(er.Error())
		return admin, er
	}
	return admin, nil
}
