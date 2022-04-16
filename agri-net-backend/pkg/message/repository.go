package message

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type IMessageRepo interface {
	SaveMessage(ctx context.Context) (*model.Message, int, error)
}
