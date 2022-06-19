package message

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMessageService interface {

	// SaveMessage uses 'message' of type *model.Message
	SaveMessage(ctx context.Context) (*model.Message, int, error)
	// GetMessages
	GetMessages(ctx context.Context, offset, limit int) ([]*model.Message, error)
	// GetRecentMessages uses "offset" type uint , "limit" type uint, and "unix_time" type uint64
	// "lang" : string,  "subscriptions" : []uint8
	GetRecentMessages(ctx context.Context) ([]*model.Message, int, error)

	// DeleteMessageBYID(ctx , messageid)
	DeleteMessageBYID(ctx context.Context, messageid uint) error
}

type MessageService struct {
	Repo IMessageRepo
}

func NewMessageService(repo IMessageRepo) IMessageService {
	return &MessageService{
		Repo: repo,
	}
}

func (service MessageService) SaveMessage(ctx context.Context) (*model.Message, int, error) {
	return service.Repo.SaveMessage(ctx)
}

func (service MessageService) GetRecentMessages(ctx context.Context) ([]*model.Message, int, error) {
	return service.Repo.GetRecentMessages(ctx)
}

func (service MessageService) GetMessages(ctx context.Context, offset, limit int) ([]*model.Message, error) {
	return service.Repo.GetMessages(ctx, offset, limit)
}

func (service MessageService) DeleteMessageBYID(ctx context.Context, messageid uint) error {
	return service.Repo.DeleteMessageBYID(ctx, messageid)
}
