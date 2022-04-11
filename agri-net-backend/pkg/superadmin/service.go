package superadmin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISuperadminService interface {
	// CheckTheExistanceOfSuperadmin uses "user_email"  of type string
	CheckTheExistanceOfSuperadmin(ctx context.Context) int
	GetSuperadminByEmail(ctx context.Context) (*model.Superadmin, int, error)
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
