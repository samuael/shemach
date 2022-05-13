package store

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IStoreRepo interface {
	CreateStore(ctx context.Context, store *model.Store) (int, error)
	GetMerchantStores(ctx context.Context, merchantID uint64) ([]*model.Store, error)
	GetStoreByID(ctx context.Context, id uint64) (*model.Store, error)
}
