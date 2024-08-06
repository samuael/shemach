package subscriber

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type ISubscriberRepo interface {
	RegisterTempoSubscriber(ctx context.Context) (int, error)
	RegisterTempoLoginSubcriber(ctx context.Context) (int, error)
	CheckTheExistanceOfPhone(ctx context.Context) (int, error)
	RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error)
	GetSubscriberByPhone(ctx context.Context) (*model.Subscriber, int, error)
	DeleteTempoLoginSubscriber(unix uint64) (int, error)
	RegisterSubscriber(ctx context.Context) (*model.Subscriber, int, error)
	GetPendingRegistrationSubscriptionByPhone(ctx context.Context) (*model.TempoSubscriber, int, error)

	DeletePendingRegistrationSubscriptionByID(ctx context.Context) (int, error)
	DeletePendingLoginSubscriptionByID(ctx context.Context) (int, error)
	GetPendingLoginSubscriptionByPhone(ctx context.Context) (*model.TempoLoginSubscriber, int, error)
	GetSubscriberByID(ctx context.Context) (*model.Subscriber, int, error)
}
