package pgx_storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/resource"
)

type ResourceRepo struct {
	DB *pgxpool.Pool
}

func NewResourceRepo(conn *pgxpool.Pool) resource.IResourceRepo {
	return &ResourceRepo{
		DB: conn,
	}
}

func (repo *ResourceRepo) SaveImagesResources(ctx context.Context, resources []*types.PostImg) error {
	for x := range resources {
		img := resources[x]
		img.CreatedAt = uint64(time.Now().Unix())
		er := repo.DB.QueryRow(ctx, `insert into img(resource,created_by,authorizations) 
		values($1,$2,$3) returning id`,
			img.Resource, img.CreatedBy, img.Role).Scan(&(img.ID))
		if er != nil {
			return er
		}
	}
	return nil
}

func (repo *ResourceRepo) GetImageDetailByID(ctx context.Context, imgid uint64) (*types.PostImg, error) {
	img := &types.PostImg{}
	if err := repo.DB.QueryRow(ctx, "select getImageDetail($1)", imgid).
		Scan(&(img.ID), &(img.Resource), &(img.Resource), &(img.CreatedBy), &(img.CreatedAt), &(img.BlurredPath), &(img.CreatedAt)); err != nil {
		return nil, err
	}
	return img, nil
}

func (repo *ResourceRepo) GetImagePath(ctx context.Context, imgid uint64) (string, error) {
	var path string
	return path, repo.DB.QueryRow(ctx, "select resource from img where id = $1", imgid).Scan(&path)
}

// CheckPropertyImageUploadEligibility checks if the property detail infomration already exists and checks if it is eligible for new upload.
// 0: for eligible
// 1: property not found
// 2: you are not the owner (the one specified by the createdBy ID)
// 3: max images exceeded
// bigint:  cover image by the specified ID already exists.
func (repo *ResourceRepo) CheckPropertyImageUploadEligibility(ctx context.Context, propertyDetailID, createdBy uint64, isCoverPicture bool, maxImages uint8) (uint64, error) {
	var statusID uint64
	return statusID, repo.DB.QueryRow(ctx, "select checkPropertyImageUploadEligibility($1, $2, $3, $4)", propertyDetailID, createdBy, isCoverPicture, maxImages).Scan(&statusID)
}

// SavePropertyImage saves an image information into the property detail
// -- 1 if the property not found
// -- 2 owner of the property mismatch mismatch
// -- 3 insertion error
// -- 4 cover picture update error
// -- 5 pictures update error
// -- bigint successful insertion.
func (repo ResourceRepo) SavePropertyImage(ctx context.Context, path string, propertyID, createdBy uint64, isCoverPicture bool, role uint8) (uint64, error) {
	var imgID uint64
	return imgID, repo.DB.QueryRow(ctx, "select savePropertyImage($1, $2, $3, $4, $5)", path, propertyID, isCoverPicture, createdBy, role).Scan(&imgID)
}
