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
	// GetInfoadmins
	GetInfoadmins(ctx context.Context) ([]*model.Infoadmin, error)
	DeleteInfoadminByID(ctx context.Context) (int, error)
	GetInfoadminByID(ctx context.Context, id uint64) (*model.Infoadmin, error)
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
func (service InfoadminService) GetInfoadmins(ctx context.Context) ([]*model.Infoadmin, error) {
	return service.Repo.GetInfoadmins(ctx)
}
func (service InfoadminService) DeleteInfoadminByID(ctx context.Context) (int, error) {
	return service.Repo.DeleteInfoadminByID(ctx)
}

func (service InfoadminService) GetInfoadminByID(ctx context.Context, id uint64) (*model.Infoadmin, error) {
	return service.Repo.GetInfoadminByID(ctx, id)
}
