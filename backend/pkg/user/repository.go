package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/storage/pgx_storage"
)

type IUserRepo interface {
	GetTempoUserInfo(ctx context.Context, phone string) (*types.UserAuthData, error)
	TempoRegisterNewUser(ctx context.Context, userData *types.UserAuthData) error
	DeleteExpiredTempoUsers(ctx context.Context, timestamp int64) (int, error)
	DeleteExpiredVerifiedPhones(ctx context.Context, timestamp int64) (int, error)
	ConfirmPhoneShortCode(ctx context.Context, phone, code string) (int64, error)
	CheckVerifiedInfoAndRegisterUser(ctx context.Context, session *types.TempoRegistrationSession, userData *types.RegistrationData) (int64, error)
	GetUserByID(ctx context.Context, id uint64) (*types.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*types.User, error)
	InsertReplaceUserProfileImage(ctx context.Context, path string, userid uint64) (status int8, oldPath string, imgID uint64, err error)
	// Forgot password related endpoints
	SaveForgotPasswordDetail(ctx context.Context, phone, code string, phoneUsed, emailUsed, useTelegram bool, maxTrials uint8) (statusCode uint64, shortCode string, err error)
	DeleteExpiredForgotPasswordSession(ctx context.Context, timestamp uint64) (int64, error)
	ConfirmForgotPasswordShortcode(ctx context.Context, phone, shortCode string) (int64, error)
	ChangeUserPassword(ctx context.Context, userID uint64, newPassword string) error
}

func NewUserRepo(db *pgxpool.Pool) IUserRepo {
	return &pgx_storage.UserRepo{
		DB: db,
	}
}
