package subscriber

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISubscriberRepo interface {
	RegisterSubscriber(ctx context.Context) (*model.TempoSubscriber, int, error)
	CheckTheExistanceOfPhone(ctx context.Context) (int, error)
	RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error)
}
