package user

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IUserRepo interface {
	GetUserByEmailOrID(ctx context.Context) (user *model.User, role int, status int, ers error)
	UpdatePassword(ctx context.Context) error
}
