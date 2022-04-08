package subscriber

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISubscriberService interface {
	// RegisterSubscriber uses "tempo_subscription"  of type *mode.TempoSubscription to save and return the subscription instance.
	// with the information of status code , error : ---
	RegisterSubscriber(ctx context.Context) (*model.TempoSubscriber, int, error)
	// CheckTheExistanceOfPhone uses 'phone'  string
	// returns int status 0 == database query error , 1== already registerd ,  2== already in a pending state ofconfirmation , or 3== the phoen is naither in the subscribers nor in the tempo_subscriber
	// error == returning the internal problems which has happened qhile querying.
	CheckTheExistanceOfPhone(ctx context.Context) (int, error)
	// RemoveExpiredTempoSubscription takes a unix time stamp to delete the tempo_confirmations
	RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error)
}

type SubscriberService struct {
	Repo ISubscriberRepo
}

func NewSubscriberService(repo ISubscriberRepo) ISubscriberService {
	return &SubscriberService{
		Repo: repo,
	}
}

func (subscriptionService *SubscriberService) RegisterSubscriber(ctx context.Context) (*model.TempoSubscriber, int, error) {
	return subscriptionService.Repo.RegisterSubscriber(ctx)
}

func (subscriptionService *SubscriberService) CheckTheExistanceOfPhone(ctx context.Context) (int, error) {
	return subscriptionService.Repo.CheckTheExistanceOfPhone(ctx)
}

func (subscriptionService *SubscriberService) RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error) {
	return subscriptionService.Repo.RemoveExpiredTempoSubscription(unixTime)
}
