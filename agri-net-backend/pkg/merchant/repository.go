package merchant

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMerchantRepo interface {
	RegisterMerchant(ctx context.Context, agent *model.Merchant) (int, error)
}
