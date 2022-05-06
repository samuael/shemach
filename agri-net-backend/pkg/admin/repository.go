package admin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAdminRepo interface {
	GetAdminByEmail(ctx context.Context, email string) (*model.Admin, error)
	CreateAdmin(ctx context.Context, admin *model.Admin) (int, int, error)
	GetAdmins(ctx context.Context, offset, limit int) ([]*model.Admin, error)
	DeleteAdminByID(ctx context.Context, id uint64) (int, error)
}
