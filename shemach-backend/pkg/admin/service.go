package admin

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

// IAdminService  ... admin service
type IAdminService interface {
	// GetAdminByEmail ...
	GetAdminByEmail(ctx context.Context, email string) (*model.Admin, error)
	// CreateAdmin ...
	CreateAdmin(ctx context.Context, admin *model.Admin) (int, int, error)
	// GetAdmins ...
	GetAdmins(ctx context.Context, offset, limit int) ([]*model.Admin, error)
	// DeleteAdminByID
	DeleteAdminByID(ctx context.Context, id uint64) (int, error)
	// GetAdminByID ...
	GetAdminByID(ctx context.Context, id uint64) (*model.Admin, error)
}

// AdminService ... admin service.
type AdminService struct {
	Repo IAdminRepo
}

func NewAdminService(repo IAdminRepo) IAdminService {
	return &AdminService{
		Repo: repo,
	}
}

// GetAdminByEmail
func (service *AdminService) GetAdminByEmail(ctx context.Context, email string) (*model.Admin, error) {
	return service.Repo.GetAdminByEmail(ctx, email)
}

// CreateAdmin
func (service *AdminService) CreateAdmin(ctx context.Context, admin *model.Admin) (int, int, error) {
	return service.Repo.CreateAdmin(ctx, admin)
}

// GetAdmins
func (service *AdminService) GetAdmins(ctx context.Context, offset, limit int) ([]*model.Admin, error) {
	return service.Repo.GetAdmins(ctx, offset, limit)
}

// DeleteAdminByID
func (service *AdminService) DeleteAdminByID(ctx context.Context, id uint64) (int, error) {
	return service.Repo.DeleteAdminByID(ctx, id)
}
func (service *AdminService) GetAdminByID(ctx context.Context, id uint64) (*model.Admin, error) {
	return service.Repo.GetAdminByID(ctx, id)
}
