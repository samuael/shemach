package session

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISessionRepo interface {
	GetSessionByUserID(ctx context.Context, id uint) (*model.SaveSession, error)
	GetSubscriberSessionByUserID(ctx context.Context, id uint) (*model.SaveSubscriberSession, error)
	SaveSession(ctx context.Context, session *model.SaveSession) error
	SaveSubscriberSession(ctx context.Context, session *model.SaveSubscriberSession) error

	DeleteSesssion(ctx context.Context, id, userid uint) error
	DeleteSubscriberSession(ctx context.Context, id, subscriberid uint) error
}
