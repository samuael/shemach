package superadmin

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type ISuperadminService interface {
	// CheckTheExistanceOfSuperadmin uses "user_email"  of type string
	CheckTheExistanceOfSuperadmin(ctx context.Context) int
	GetSuperadminByEmail(ctx context.Context) (*model.Superadmin, int, error)
	GetSuperadminByID(ctx context.Context, id int) (*model.Superadmin, error)
	GetAllSuperadmins(ctx context.Context) ([]*model.Superadmin, error)
}

type SuperadminService struct {
	Repo ISuperadminRepo
}

func NewSuperadminService(repo ISuperadminRepo) ISuperadminService {
	return &SuperadminService{
		Repo: repo,
	}
}

func (service *SuperadminService) CheckTheExistanceOfSuperadmin(ctx context.Context) int {
	return service.Repo.CheckTheExistanceOfSuperadmin(ctx)
}

func (service *SuperadminService) GetSuperadminByEmail(ctx context.Context) (*model.Superadmin, int, error) {
	return service.Repo.GetSuperadminByEmail(ctx)
}

func (service *SuperadminService) GetSuperadminByID(ctx context.Context, id int) (*model.Superadmin, error) {
	return service.Repo.GetSuperadminByID(ctx, id)
}

func (service *SuperadminService) GetAllSuperadmins(ctx context.Context) ([]*model.Superadmin, error) {
	return service.Repo.GetAllSuperadmins(ctx)
}
