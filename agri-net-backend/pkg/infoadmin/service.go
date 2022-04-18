package infoadmin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IInfoadminService interface {
	// GetInfoadminByEmail  uses "infoadmin_email" type string
	GetInfoadminByEmail(ctx context.Context) (*model.Infoadmin, error)
	// CreateInfoadmin uses "info_admin" of type *model.Infoadmin
	CreateInfoadmin(ctx context.Context) (*model.Infoadmin, error)
}

type InfoadminService struct {
	Repo IInfoadminRepo
}

func NewInfoadminService(repo IInfoadminRepo) IInfoadminService {
	return &InfoadminService{
		Repo: repo,
	}
}

func (service InfoadminService) GetInfoadminByEmail(ctx context.Context) (*model.Infoadmin, error) {
	return service.Repo.GetInfoadminByEmail(ctx)
}

func (service InfoadminService) CreateInfoadmin(ctx context.Context) (*model.Infoadmin, error) {
	return service.Repo.CreateInfoadmin(ctx)
}
