package admin

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IAdminRepo interface {
	GetAdminByEmail(ctx context.Context, email string) (*model.Admin, error)
	CreateAdmin(ctx context.Context, admin *model.Admin) (int, int, error)
	GetAdmins(ctx context.Context, offset, limit int) ([]*model.Admin, error)
	DeleteAdminByID(ctx context.Context, id uint64) (int, error)
	GetAdminByID(ctx context.Context, id uint64) (*model.Admin, error)
}
