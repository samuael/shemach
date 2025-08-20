package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/session"
)

type SessionRepo struct {
	DB *pgxpool.Pool
}

func NewSessionRepo(conn *pgxpool.Pool) session.ISessionRepo {
	return &SessionRepo{
		DB: conn,
	}
}

func (repo *SessionRepo) GetSessionByUserID(ctx context.Context, id uint) (*types.SaveSession, error) {
	session := &types.SaveSession{}
	er := repo.DB.QueryRow(ctx, "select id,userid,token from session where userid=$1", id).Scan(&(session.ID), &(session.UserID), &(session.Token))
	if er != nil {
		return nil, er
	}
	return session, nil
}
func (repo *SessionRepo) SaveSession(ctx context.Context, session *types.SaveSession) error {
	se, err := repo.GetSessionByUserID(ctx, uint(session.UserID))
	if se != nil && err == nil {
		err = repo.DB.QueryRow(ctx, "update session set token=$1 where userid=$2 returning id", session.Token, session.UserID).Scan(&(session.ID))
	} else {
		err = repo.DB.QueryRow(ctx, "insert into session(userid,token) values( $1 , $2) returning id", session.UserID, session.Token).Scan(&(session.ID))
	}
	return err
}

func (repo *SessionRepo) DeleteSesssion(ctx context.Context, id, userid uint) error {
	com, er := repo.DB.Exec(ctx, "DELETE FROM session WHERE id=$1  or userid=$2", id, userid)
	if er != nil {
		println(er.Error())
	}
	if com.RowsAffected() == 0 || er != nil {
		return errors.New(" no rows affected")
	}
	return nil
}

// DeleteSessionBeforeTimestamp deletes session instances that are before the specified timestamp.
func (repo *SessionRepo) DeleteSessionBeforeTimestamp(ctx context.Context, timestamp int64) (int64, error) {
	var userID int64
	return userID, repo.DB.QueryRow(ctx, "DELETE FROM session WHERE created_at <=$1 returning id").Scan(&userID)
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
