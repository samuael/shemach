package subscriber

import (
	"context"
)

type ISubscriberRepo interface {
	RegisterSubscriber(ctx context.Context) (int, error)
	CheckTheExistanceOfPhone(ctx context.Context) (int, error)
	RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error)
}
