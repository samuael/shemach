package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/session"
)

type SessionRepo struct {
	DB *pgxpool.Pool
}

func NewSessionRepo(conn *pgxpool.Pool) session.ISessionRepo {
	return &SessionRepo{
		DB: conn,
	}
}

func (repo *SessionRepo) GetSessionByUserID(ctx context.Context, id uint) (*model.SaveSession, error) {
	session := &model.SaveSession{}
	er := repo.DB.QueryRow(ctx, "select id,userid,token from session where userid=$1", id).Scan(&(session.ID), &(session.UserID), &(session.Token))
	if er != nil {
		return nil, er
	}
	return session, nil
}
func (repo *SessionRepo) GetSubscriberSessionByUserID(ctx context.Context, id uint) (*model.SaveSubscriberSession, error) {
	session := &model.SaveSubscriberSession{}
	er := repo.DB.QueryRow(ctx, "select id ,subscriberid ,token from subscriber_session where subscriberid=$1", id).Scan(&(session.ID), &(session.SubscriberID), &(session.Token))
	if er != nil {
		return nil, er
	}
	return session, nil
}
func (repo *SessionRepo) SaveSession(ctx context.Context, session *model.SaveSession) error {
	se, er := repo.GetSessionByUserID(ctx, uint(session.UserID))
	if se != nil && er == nil {
		er = repo.DB.QueryRow(ctx, "update session set token=$1 where userid=$2 returning id", session.Token, session.UserID).Scan(&(session.ID))
	} else {
		er = repo.DB.QueryRow(ctx, "insert into session(userid,token) values( $1 , $2 ) returning id", session.UserID, session.Token).Scan(&(session.ID))
	}
	return er
}
func (repo *SessionRepo) SaveSubscriberSession(ctx context.Context, session *model.SaveSubscriberSession) error {
	se, er := repo.GetSubscriberSessionByUserID(ctx, uint(session.SubscriberID))
	if se != nil && er == nil {
		er = repo.DB.QueryRow(ctx, "update subscriber_session set token=$1 where subscriberid=$2 returning id", session.Token, session.SubscriberID).Scan(&(session.ID))
	} else {
		er = repo.DB.QueryRow(ctx, "insert into subscriber_session(subscriberid,token) values( $1 , $2 ) returning id", session.SubscriberID, session.Token).Scan(&(session.ID))
	}
	return er
}
func (repo *SessionRepo) DeleteSesssion(ctx context.Context, id, userid uint) error {
	com, er := repo.DB.Exec(ctx, "DELETE FROM session WHERE id=$1  or userid=$2", id, userid)
	println(com.RowsAffected())
	if er != nil {
		println(er.Error())
	}
	if com.RowsAffected() == 0 || er != nil {
		return errors.New(" no rows affected")
	}
	return nil
}
func (repo *SessionRepo) DeleteSubscriberSession(ctx context.Context, id, subscriberid uint) error {
	com, er := repo.DB.Exec(ctx, "DELETE FROM subscriber_session WHERE id=$1 or subscriberid=$2", id, subscriberid)
	println(com.RowsAffected())
	if er != nil {
		println(er.Error())
	}
	if com.RowsAffected() == 0 || er != nil {
		return errors.New(" no rows affected")
	}
	return nil
}
