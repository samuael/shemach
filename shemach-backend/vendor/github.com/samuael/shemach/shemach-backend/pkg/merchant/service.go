package merchant

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IMerchantService interface {
	RegisterMerchant(ctx context.Context, Merchant *model.Merchant) (int, error)
	GetMerchantByID(ctx context.Context, id int) (*model.Merchant, error)
	CreateSubscriptions(ctx context.Context, productid uint8, merchantid uint64) (status int)
	UnsubscribeProduct(ctx context.Context, productid uint8, merchantid uint64) (status int)
	SearchMerchants(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Merchant, error)
	DeleteMerchantByID(ctx context.Context, merchantid uint64) error
}

type MerchantService struct {
	Repo IMerchantRepo
}

func NewMerchantService(repo IMerchantRepo) IMerchantService {
	return &MerchantService{
		Repo: repo,
	}
}

// RegisterMerchant ...
func (service *MerchantService) RegisterMerchant(ctx context.Context, Merchant *model.Merchant) (int, error) {
	return service.Repo.RegisterMerchant(ctx, Merchant)
}
func (service *MerchantService) GetMerchantByID(ctx context.Context, id int) (*model.Merchant, error) {
	return service.Repo.GetMerchantByID(ctx, id)
}
func (service *MerchantService) CreateSubscriptions(ctx context.Context, productid uint8, merchantid uint64) (status int) {
	return service.Repo.CreateSubscriptions(ctx, productid, merchantid)
}
func (service *MerchantService) UnsubscribeProduct(ctx context.Context, productid uint8, merchantid uint64) (status int) {
	return service.Repo.UnsubscribeProduct(ctx, productid, merchantid)
}

func (service *MerchantService) SearchMerchants(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Merchant, error) {
	return service.Repo.SearchMerchants(ctx, phone, name, createdBy, offset, limit)
}

func (service *MerchantService) DeleteMerchantByID(ctx context.Context, merchantid uint64) error {
	return service.Repo.DeleteMerchantByID(ctx, merchantid)
}
