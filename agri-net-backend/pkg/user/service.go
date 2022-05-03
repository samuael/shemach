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
	// GetImageUrl uses "user_id": uint64
	GetImageUrl(ctx context.Context) string
	// ChangeImageUrl uses "user_id": uint64 and "image_url":string
	ChangeImageUrl(ctx context.Context) bool
	// DeletePendingEmailConfirmation usess a positional argument
	DeletePendingEmailConfirmation(timestamp uint64) error
	// SaveEmailConfirmation ...
	SaveEmailConfirmation(ctx context.Context, emailc *model.EmailConfirmation) error
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

func (service *UserService) GetImageUrl(ctx context.Context) string {
	return service.Repo.GetImageUrl(ctx)
}

func (service *UserService) ChangeImageUrl(ctx context.Context) bool {
	return (service.Repo.ChangeImageUrl(ctx) == nil)
}
func (service *UserService) DeletePendingEmailConfirmation(timestamp uint64) error {
	return service.Repo.DeletePendingEmailConfirmation(timestamp)
}

func (service *UserService) SaveEmailConfirmation(ctx context.Context, emailc *model.EmailConfirmation) error {
	return service.Repo.SaveEmailConfirmation(ctx, emailc)
}
