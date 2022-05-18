package merchant

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMerchantService interface {
	RegisterMerchant(ctx context.Context, Merchant *model.Merchant) (int, error)
	GetMerchantByID(ctx context.Context, id int) (*model.Merchant, error)
	CreateSubscriptions(ctx context.Context, productid uint8, merchantid uint64) (status int)
	UnsubscribeProduct(ctx context.Context, productid uint8, merchantid uint64) (status int)
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
