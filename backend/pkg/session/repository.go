package session

import (
	"context"

	"github.com/samuael/shemach/backend/pkg/constants/types"
)

type ISessionRepo interface {
	DeleteSessionBeforeTimestamp(ctx context.Context, timestamp int64) (int64, error)

	GetSessionByUserID(ctx context.Context, id uint) (*types.SaveSession, error)
	SaveSession(ctx context.Context, session *types.SaveSession) error

	DeleteSesssion(ctx context.Context, id, userid uint) error
	DeleteSubscriberSession(ctx context.Context, id, subscriberid uint) error
}
