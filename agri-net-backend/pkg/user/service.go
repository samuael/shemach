package user

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IUserService interface {
	// GetUserByEmailOrID uses "user_id" uint64 and "user_email" string to returns
	// user *mdoel.User ,  role int, statsuode int , er error
	GetUserByEmailOrID(ctx context.Context) (user *model.User, role int, status int, ers error)
	// UpdatePassword uses 'user_id':uint64 and 'new_password':string to update the password of a user.
	UpdatePassword(ctx context.Context) error
}

type UserService struct {
	Repo IUserRepo
}

func NewUserService(repo IUserRepo) IUserService {
	return &UserService{
		Repo: repo,
	}
}

// GetUserByEmailOrID uses "id" and "email" returns *mdoel.User , role int
func (service *UserService) GetUserByEmailOrID(ctx context.Context) (user *model.User, role int, status int, ers error) {
	return service.Repo.GetUserByEmailOrID(ctx)
}
func (service *UserService) UpdatePassword(ctx context.Context) error {
	return service.Repo.UpdatePassword(ctx)
}
