package subscriber

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type ISubscriberService interface {
	// RegisterSubscriber uses "tempo_subscriber"  of type *mode.TempoSubscription to save and return the subscription instance.
	// with the information of status code , error : ---
	RegisterTempoSubscriber(ctx context.Context) (int, error)
	// RegisterSubscriber "subscriber"  of type *model.Subscriber
	// returns *model.Subscriber with it's new ID and int status code and error ( error message )
	RegisterSubscriber(ctx context.Context) (*model.Subscriber, int, error)
	// CheckTheExistanceOfPhone uses 'phone'  string
	// returns int status 0 == database query error , 1== already registerd ,  2== already in a pending state ofconfirmation , or 3== the phoen is naither in the subscribers nor in the tempo_subscriber
	// error == returning the internal problems which has happened qhile querying.
	CheckTheExistanceOfPhone(ctx context.Context) (int, error)
	// RemoveExpiredTempoSubscription takes a unix time stamp to delete the tempo_confirmations
	RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error)
	// GetSubscriberByPhone uses "subscriber_phone"  type string
	GetSubscriberByPhone(ctx context.Context) (*model.Subscriber, int, error)
	// GetSubscriberByID uses "subscriber_id"  type uint64
	GetSubscriberByID(ctx context.Context) (*model.Subscriber, int, error)
	// DeleteTempoLoginSubscriber returns the number of rows and error if exist.
	DeleteTempoLoginSubscriber(unix uint64) (int, error)
	// RegisterTempoLoginSubcriber uses "login_tempo_subscriber" of type *mdoel.TempoLoginSubscriber
	RegisterTempoLoginSubcriber(ctx context.Context) (int, error)
	// GetPendingRegistrationSubscriptionByPhone uses "subscriber_phone" type string
	GetPendingRegistrationSubscriptionByPhone(ctx context.Context) (*model.TempoSubscriber, int, error)
	// DeletePendingRegistrationSubscriptionByID  uses "subscription_id" type uint64
	DeletePendingRegistrationSubscriptionByID(ctx context.Context) (int, error)
	// DeletePendingLoginSubscriptionByID uses "subscription_id" type uint64
	DeletePendingLoginSubscriptionByID(ctx context.Context) (int, error)
	// GetPendingLoginSubscriptionByPhone uses "subscriber_phone" type string
	GetPendingLoginSubscriptionByPhone(ctx context.Context) (*model.TempoLoginSubscriber, int, error)
}

type SubscriberService struct {
	Repo ISubscriberRepo
}

func NewSubscriberService(repo ISubscriberRepo) ISubscriberService {
	return &SubscriberService{
		Repo: repo,
	}
}

func (subscriptionService *SubscriberService) RegisterTempoSubscriber(ctx context.Context) (int, error) {
	return subscriptionService.Repo.RegisterTempoSubscriber(ctx)
}

func (subscriptionService *SubscriberService) CheckTheExistanceOfPhone(ctx context.Context) (int, error) {
	return subscriptionService.Repo.CheckTheExistanceOfPhone(ctx)
}

func (subscriptionService *SubscriberService) RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error) {
	return subscriptionService.Repo.RemoveExpiredTempoSubscription(unixTime)
}

func (subscriptionService *SubscriberService) GetSubscriberByPhone(ctx context.Context) (*model.Subscriber, int, error) {
	return subscriptionService.Repo.GetSubscriberByPhone(ctx)
}

func (subscriptionService *SubscriberService) DeleteTempoLoginSubscriber(unix uint64) (int, error) {
	return subscriptionService.Repo.DeleteTempoLoginSubscriber(unix)
}

func (subscriptionService *SubscriberService) RegisterSubscriber(ctx context.Context) (*model.Subscriber, int, error) {
	return subscriptionService.Repo.RegisterSubscriber(ctx)
}
func (subscriptionService *SubscriberService) RegisterTempoLoginSubcriber(ctx context.Context) (int, error) {
	return subscriptionService.Repo.RegisterTempoLoginSubcriber(ctx)
}

func (subscriptionService *SubscriberService) GetPendingRegistrationSubscriptionByPhone(ctx context.Context) (*model.TempoSubscriber, int, error) {
	return subscriptionService.Repo.GetPendingRegistrationSubscriptionByPhone(ctx)
}
func (subscriptionService *SubscriberService) DeletePendingRegistrationSubscriptionByID(ctx context.Context) (int, error) {
	return subscriptionService.Repo.DeletePendingRegistrationSubscriptionByID(ctx)
}

// DeletePendingLoginSubscriptionByID
func (subscriptionService *SubscriberService) DeletePendingLoginSubscriptionByID(ctx context.Context) (int, error) {
	return subscriptionService.Repo.DeletePendingLoginSubscriptionByID(ctx)
}

// GetPendingLoginSubscriptionByPhone
func (subscriptionService *SubscriberService) GetPendingLoginSubscriptionByPhone(ctx context.Context) (*model.TempoLoginSubscriber, int, error) {
	return subscriptionService.Repo.GetPendingLoginSubscriptionByPhone(ctx)
}
func (subscriptionService *SubscriberService) GetSubscriberByID(ctx context.Context) (*model.Subscriber, int, error) {
	return subscriptionService.Repo.GetSubscriberByID(ctx)
}
