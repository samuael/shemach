package infoadmin

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IInfoadminRepo interface {
	GetInfoadminByEmail(ctx context.Context) (*model.Infoadmin, error)
	CreateInfoadmin(ctx context.Context) (*model.Infoadmin, error)
	GetInfoadmins(ctx context.Context) ([]*model.Infoadmin, error)
	DeleteInfoadminByID(ctx context.Context) (int, error)
	GetInfoadminByID(ctx context.Context, id uint64) (*model.Infoadmin, error)
}
