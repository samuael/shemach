package pgx_storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/resource"
)

type ResourceRepo struct {
	DB *pgxpool.Pool
}

func NewResourceRepo(conn *pgxpool.Pool) resource.IResourceRepo {
	return &ResourceRepo{
		DB: conn,
	}
}

func (repo *ResourceRepo) SaveImagesResources(ctx context.Context, resources []*model.PostImg) error {
	for x := range resources {
		img := resources[x]
		img.CreatedAt = uint64(time.Now().Unix())
		er := repo.DB.QueryRow(ctx, `insert into img(resource,owner_id,owner_role,authorized,authorizations,blurred_res) 
		values($1,$2,$3,$4,$5,$6) returning img_id`,
			img.Resource, img.OwnerID, img.OwnerRole, img.Authorized, img.Authorizations, img.BlurredRe,
		).Scan(&(img.ID))
		if er != nil {
			return er
		}
	}
	return nil
}

func (repo *ResourceRepo) GetImageByID(ctx context.Context, imgid uint64) (*model.PostImg, error) {
	img := &model.PostImg{}
	if er := repo.DB.QueryRow(ctx, "select  img_id,resource,owner_id,owner_role,authorized,authorizations,created_at,blurred_res from img where img_id=$1", imgid).
		Scan(
			&(img.ID), &(img.Resource), &(img.OwnerID), &(img.OwnerRole), &(img.Authorized), &(img.Authorizations), &(img.CreatedAt), &(img.BlurredRe),
		); er != nil {
		return nil, er
	}
	return img, nil
}
