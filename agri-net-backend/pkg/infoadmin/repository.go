package infoadmin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IInfoadminRepo interface {
	GetInfoadminByEmail(ctx context.Context) (*model.Infoadmin, error)
	CreateInfoadmin(ctx context.Context) (*model.Infoadmin, error)
}
