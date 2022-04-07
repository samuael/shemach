package admin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

// Interfaces to be implemented by the admin service instances
type IAdminService interface {
	AdminByEmail(ctx context.Context) (*model.Admin, error)
	ChangePassword(ctx context.Context) (bool, error)
	// DeleteAccountByEmail uses "email"  type string
	DeleteAccountByEmail(context.Context) (bool, int)
	// DeleteAccountByID users "admin_id" type "uint64"
	DeleteAccountByID(ctx context.Context) (bool, int)
	CreateAdmin(context.Context) (*model.Admin, error)
	// AdminByID uses "admin_id"
	AdminByID(ctx context.Context) (*model.Admin, error)
	UpdateAdmin(ctx context.Context) (*model.Admin, error)

	// GetImageUrl  uses session in the context of the application to retrieve the user informationa
	GetImageUrl(ctx context.Context) string
	// ChangeImageUrl uses 'image_url' and 'user_id' to modify the user's profile Picture.
	ChangeImageUrl(ctx context.Context) bool
	// DeleteProfilePicture uses the session to delete the imgurl
	DeleteProfilePicture(ctx context.Context) bool
	// GetAdmins ....
	GetAdmins(ctx context.Context) ([]*model.Admin, error)
}

// AdminService struct representing a admin service
type AdminService struct {
	Repo IAdminRepo
}

// NewAdminService function returninng an admin service  instance
func NewAdminService(repo IAdminRepo) IAdminService {
	return &AdminService{
		Repo: repo,
	}
}

func (adminser *AdminService) AdminByEmail(ctx context.Context) (*model.Admin, error) {
	return adminser.Repo.AdminByEmail(ctx)
}

// ChangePassword (ctx context.Context) (bool, error)
func (adminser *AdminService) ChangePassword(ctx context.Context) (bool, error) {
	return adminser.Repo.ChangePassword(ctx)
}

func (adminser *AdminService) DeleteAccountByEmail(ctx context.Context) (bool, int) {
	return adminser.Repo.DeleteAccountByEmail(ctx)
}

// CreateAdmin(context.Context) (*model.Admin, error)
func (adminser *AdminService) CreateAdmin(ctx context.Context) (*model.Admin, error) {
	return adminser.Repo.CreateAdmin(ctx)
}

func (adminser *AdminService) AdminByID(ctx context.Context) (*model.Admin, error) {
	return adminser.Repo.AdminByID(ctx)
}
func (adminser *AdminService) UpdateAdmin(ctx context.Context) (*model.Admin, error) {
	return adminser.Repo.UpdateAdmin(ctx)
}

func (adminser *AdminService) GetImageUrl(ctx context.Context) string {
	img, er := adminser.Repo.GetImageUrl(ctx)
	if er != nil {
		return ""
	}
	return img
}
func (adminser *AdminService) ChangeImageUrl(ctx context.Context) bool {
	return adminser.Repo.ChangeImageUrl(ctx) == nil
}

// DeleteProfilePicture(ctx context.Context) error
func (adminser *AdminService) DeleteProfilePicture(ctx context.Context) bool {
	return adminser.Repo.DeleteProfilePicture(ctx) == nil
}

func (adminser *AdminService) DeleteAccountByID(ctx context.Context) (bool, int) {
	return adminser.Repo.DeleteAccountByID(ctx)
}
func (adminser *AdminService) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	return adminser.Repo.GetAdmins(ctx)
}
