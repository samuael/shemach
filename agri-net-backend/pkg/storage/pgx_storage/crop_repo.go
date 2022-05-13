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
	result := 0
	er := repo.DB.QueryRow(ctx, `select * from createProductPost( $1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		crop.TypeID, crop.Description, crop.Negotiable, crop.RemainingQuantity, crop.SellingPrice,
		crop.AddressRef, crop.StoreID, crop.AgentID, crop.StoreOwned,
	).Scan(&result)
	if er != nil || result <= 0 {
		if er != nil {
			println(er.Error())
		}
		return result, er
	}
	crop.ID = uint64(result)
	return result, nil
}
