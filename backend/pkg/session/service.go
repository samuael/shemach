package session

import (
	"context"

	"github.com/samuael/shemach/backend/pkg/constants/types"
)

type ISessionService interface {
	DeleteSessionBeforeTimestamp(ctx context.Context, timestamp int64) (int64, error)

	GetSessionByUserID(ctx context.Context, id uint) (*types.SaveSession, error)
	SaveSession(ctx context.Context, session *types.SaveSession) error
	DeleteSesssion(ctx context.Context, id, userid uint) error
	DeleteSubscriberSession(ctx context.Context, id, userid uint) error
}

type SessionService struct {
	Repo ISessionRepo
}

func NewSessionService(repo ISessionService) ISessionService {
	return &SessionService{
		Repo: repo,
	}
}

// DeleteSessionBeforeTimestamp deletes session instances before the specified timestamp.
func (service *SessionService) DeleteSessionBeforeTimestamp(ctx context.Context, timestamp int64) (int64, error) {
	return service.Repo.DeleteSessionBeforeTimestamp(ctx, timestamp)
}

func (sservice *SessionService) GetSessionByUserID(ctx context.Context, id uint) (*types.SaveSession, error) {
	return sservice.Repo.GetSessionByUserID(ctx, id)
}
func (sservice *SessionService) SaveSession(ctx context.Context, session *types.SaveSession) error {
	return sservice.Repo.SaveSession(ctx, session)
}
func (sservice *SessionService) DeleteSesssion(ctx context.Context, id, userid uint) error {
	return sservice.Repo.DeleteSesssion(ctx, id, userid)
}
func (sservice *SessionService) DeleteSubscriberSession(ctx context.Context, id, userid uint) error {
	return sservice.Repo.DeleteSubscriberSession(ctx, id, userid)
}
