package superadmin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISuperadminRepo interface {
	CheckTheExistanceOfSuperadmin(ctx context.Context) int
	GetSuperadminByEmail(ctx context.Context) (*model.Superadmin, int, error)
	GetSuperadminByID(ctx context.Context, id int) (*model.Superadmin, error)
	GetAllSuperadmins(ctx context.Context) ([]*model.Superadmin, error)
}
