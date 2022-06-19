package message

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMessageRepo interface {
	SaveMessage(ctx context.Context) (*model.Message, int, error)
	GetRecentMessages(ctx context.Context) ([]*model.Message, int, error)
	GetMessages(ctx context.Context, offset, limit int) ([]*model.Message, error)
	DeleteMessageBYID(ctx context.Context, messageid uint) error
}
