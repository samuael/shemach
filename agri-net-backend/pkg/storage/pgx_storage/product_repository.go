package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
)

type ProductRepo struct {
	DB *pgxpool.Pool
}

func NewProductRepo(conn *pgxpool.Pool) product.IProductRepo {
	return &ProductRepo{
		DB: conn,
	}
}

func (prohan *ProductRepo) CreateNewProduct(ctx context.Context) (*model.Product, int, error) {
	return nil, 0, nil
}
