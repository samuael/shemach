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
	SaveEmailConfirmation(ctx context.Context, emailc *model.EmailConfirmation) (int, error)
	// UpdateUser used *model.User
	UpdateUser(ctx context.Context, user *model.User) (int, error)

	// GetUserByPhone
	GetUserByPhone(ctx context.Context, phone string) (user *model.User, role int, status int, er error)
	// RegisterTempoCXP ...
	RegisterTempoCXP(ctx context.Context, tempo *model.TempoCXP) error
	GetTempoCXP(ctx context.Context, phone string, response model.TempoCXP) error
	RemoveTempoCXP(ctx context.Context, phone string) error
	RemoveExpiredCXPConfirmations(timestamp uint64) (count int, er error)
	ConfirmUserEmailUpdate(ctx context.Context, id uint64, newemail, oldemail string) error
	GetEmailInConfirmationByID(ctx context.Context, id uint64) (*model.EmailConfirmation, error)
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

func (service *UserService) SaveEmailConfirmation(ctx context.Context, emailc *model.EmailConfirmation) (int, error) {
	return service.Repo.SaveEmailConfirmation(ctx, emailc)
}

// UpdateUser
func (service *UserService) UpdateUser(ctx context.Context, user *model.User) (int, error) {
	return service.Repo.UpdateUser(ctx, user)
}

// GetUserByPhone
func (service *UserService) GetUserByPhone(ctx context.Context, phone string) (user *model.User, role int, status int, er error) {
	return service.Repo.GetUserByPhone(ctx, phone)
}

// RegisterTempoCXP
func (service *UserService) RegisterTempoCXP(ctx context.Context, tempo *model.TempoCXP) error {
	return service.Repo.RegisterTempoCXP(ctx, tempo)
}

// GetTempoCXP
func (service *UserService) GetTempoCXP(ctx context.Context, phone string, response model.TempoCXP) error {
	return service.Repo.GetTempoCXP(ctx, phone, response)
}

// RemoveTempoCXP  ... hanldser service method
func (service *UserService) RemoveTempoCXP(ctx context.Context, phone string) error {
	return service.Repo.RemoveTempoCXP(ctx, phone)
}

// RemoveExpiredCXPConfirmations deleting expired confirmaiton messages of CXPs and their data from the database.
func (service *UserService) RemoveExpiredCXPConfirmations(timestamp uint64) (count int, er error) {
	return service.Repo.RemoveExpiredCXPConfirmations(timestamp)
}

func (service *UserService) ConfirmUserEmailUpdate(ctx context.Context, id uint64, newemail, oldemail string) error {
	return service.Repo.ConfirmUserEmailUpdate(ctx, id, newemail, oldemail)
}

// GetEmailInConfirmationByID(ctx context.Context, id uint64) (*model.EmailConfirmation, error)
func (service *UserService) GetEmailInConfirmationByID(ctx context.Context, id uint64) (*model.EmailConfirmation, error) {
	return service.Repo.GetEmailInConfirmationByID(ctx, id)
}
