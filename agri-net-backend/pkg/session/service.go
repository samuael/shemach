package session

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ISessionService interface {
	GetSessionByUserID(ctx context.Context, id uint) (*model.SaveSession, error)
	GetSubscriberSessionByUserID(ctx context.Context, id uint) (*model.SaveSubscriberSession, error)
	SaveSession(ctx context.Context, session *model.SaveSession) error
	SaveSubscriberSession(ctx context.Context, session *model.SaveSubscriberSession) error
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

func (sservice *SessionService) GetSessionByUserID(ctx context.Context, id uint) (*model.SaveSession, error) {
	return sservice.Repo.GetSessionByUserID(ctx, id)
}
func (sservice *SessionService) GetSubscriberSessionByUserID(ctx context.Context, id uint) (*model.SaveSubscriberSession, error) {
	return sservice.Repo.GetSubscriberSessionByUserID(ctx, id)
}
func (sservice *SessionService) SaveSession(ctx context.Context, session *model.SaveSession) error {
	return sservice.Repo.SaveSession(ctx, session)
}
func (sservice *SessionService) SaveSubscriberSession(ctx context.Context, session *model.SaveSubscriberSession) error {
	return sservice.Repo.SaveSubscriberSession(ctx, session)
}
func (sservice *SessionService) DeleteSesssion(ctx context.Context, id, userid uint) error {
	return sservice.Repo.DeleteSesssion(ctx, id, userid)
}
func (sservice *SessionService) DeleteSubscriberSession(ctx context.Context, id, userid uint) error {
	return sservice.Repo.DeleteSubscriberSession(ctx, id, userid)
}
