package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
)

type CropRepo struct {
	DB *pgxpool.Pool
}

func NewCropRepo(conn *pgxpool.Pool) crop.ICropRepo {
	return &CropRepo{
		DB: conn,
	}
}

func (repo *CropRepo) CreateCrop(ctx context.Context, crop *model.Crop) (int, error) {
	return 0, nil
}
