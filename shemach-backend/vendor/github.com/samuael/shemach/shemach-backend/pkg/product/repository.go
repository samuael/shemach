package product

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IProductRepo interface {
	CreateNewProduct(ctx context.Context) (*model.Product, int, error)
	// This is the products repository interface.
	CheckTheExistanceOfProductInformation(ctx context.Context) bool
	GetProducts(ctx context.Context) ([]*model.Product, int, error)
	GetProductByID(ctx context.Context) (*model.Product, int, error)
	CreateSubscriptions(ctx context.Context) (status int)
	UnsubscribeProduct(ctx context.Context) (status int)
	UpdateProductPrice(ctx context.Context) (int, int, error)
	SearchProductsByText(ctx context.Context) ([]*model.Product, int, error)
}
