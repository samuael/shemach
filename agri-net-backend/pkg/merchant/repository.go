package merchant

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMerchantRepo interface {
	RegisterMerchant(ctx context.Context, agent *model.Merchant) (int, error)
	GetMerchantByID(ctx context.Context, id int) (*model.Merchant, error)
	CreateSubscriptions(ctx context.Context, productid uint8, merchantid uint64) (status int)
	UnsubscribeProduct(ctx context.Context, productid uint8, merchantid uint64) (status int)
	SearchMerchants(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Merchant, error)
}
