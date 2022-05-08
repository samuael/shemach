package user

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IUserRepo interface {
	GetUserByEmailOrID(ctx context.Context) (user *model.User, role int, status int, ers error)
	UpdatePassword(ctx context.Context) error
	GetImageUrl(ctx context.Context) string
	ChangeImageUrl(ctx context.Context) error
	DeletePendingEmailConfirmation(timestamp uint64) error
	SaveEmailConfirmation(ctx context.Context, emailc *model.EmailConfirmation) (int, error)
	UpdateUser(ctx context.Context, user *model.User) (int, error)
}
