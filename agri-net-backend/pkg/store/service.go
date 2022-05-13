package store

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IStoreService interface {
	CreateStore(ctx context.Context, store *model.Store) (int, error)
	GetMerchantStores(ctx context.Context, merchantID uint64) ([]*model.Store, error)
	GetStoreByID(ctx context.Context, id uint64) (*model.Store, error)
}

type StoreService struct {
	Repo IStoreRepo
}

func NewStoreService(repo IStoreRepo) IStoreService {
	return &StoreService{
		Repo: repo,
	}
}

// CreateStore create a store instance
func (service *StoreService) CreateStore(ctx context.Context, store *model.Store) (int, error) {
	return service.Repo.CreateStore(ctx, store)
}

// GetMerchantStores ...
func (service *StoreService) GetMerchantStores(ctx context.Context, merchantID uint64) ([]*model.Store, error) {
	return service.Repo.GetMerchantStores(ctx, merchantID)
}

// GetStoreByID ...
func (service *StoreService) GetStoreByID(ctx context.Context, id uint64) (*model.Store, error) {
	return service.Repo.GetStoreByID(ctx, id)
}
