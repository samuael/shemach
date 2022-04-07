package admin

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IAdminRepo interface {
	AdminByEmail(ctx context.Context) (*model.Admin, error)
	ChangePassword(ctx context.Context) (bool, error)
	// DeleteAccountByEmail ..
	DeleteAccountByEmail(context.Context) (bool, int)
	// DeleteAccountByID ..
	DeleteAccountByID(ctx context.Context) (bool, int)
	CreateAdmin(ctx context.Context) (*model.Admin, error)
	AdminByID(ctx context.Context) (*model.Admin, error)
	UpdateAdmin(ctx context.Context) (*model.Admin, error)
	// GetImageUrl  uses session in the context of the application to retrieve the user informationa
	GetImageUrl(ctx context.Context) (string, error)
	// ChangeImageUrl uses 'image_url' and 'user_id' to modify the user's profile Picture.
	ChangeImageUrl(ctx context.Context) error
	// DeleteProfilePicture(ctx context.Context) error  uses the session in the context to get the user ID and profile
	// and delete the profile picture.
	DeleteProfilePicture(ctx context.Context) error
	GetAdmins(ctx context.Context) ([]*model.Admin, error)
}
