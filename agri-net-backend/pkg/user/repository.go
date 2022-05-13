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
	GetUserByPhone(ctx context.Context, phone string) (user *model.User, role int, status int, er error)
	RegisterTempoCXP(ctx context.Context, tempo *model.TempoCXP) error
	GetTempoCXP(ctx context.Context, phone string, response model.TempoCXP) error
	RemoveTempoCXP(ctx context.Context, phone string) error
	RemoveExpiredCXPConfirmations(timestamp uint64) (count int, er error)
	ConfirmUserEmailUpdate(ctx context.Context, id uint64, newemail, oldemail string) error
	GetEmailInConfirmationByID(ctx context.Context, id uint64) (*model.EmailConfirmation, error)
}
