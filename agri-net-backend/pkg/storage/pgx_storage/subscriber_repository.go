package pgx_storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
)

// AdminRepo ...
type SubscriberRepo struct {
	DB *pgxpool.Pool
}

func NewSubscriberRepo(db *pgxpool.Pool) subscriber.ISubscriberRepo {
	return &SubscriberRepo{
		DB: db,
	}
}

// RegisterSubscriber
func (repo *SubscriberRepo) RegisterTempoSubscriber(ctx context.Context) (int, error) {
	tempo := ctx.Value("tempo_subscriber").(*model.TempoSubscriber)
	er := repo.DB.QueryRow(ctx, "insert into tempo_subscriber( fullname, phone, lang, role, confirmation, unix) values( $1,$2,$3,$4,$5,$6) returning id",
		tempo.Fullname, tempo.Phone, tempo.Lang, tempo.Role, tempo.ConfirmationCode, tempo.Unix).Scan(&(tempo.ID))
	if er != nil {
		return state.DT_STATUS_DBQUERY_ERROR, er
	}
	return state.DT_STATUS_OK, nil
}

// CheckTheExistanceOfPhone
func (repo *SubscriberRepo) CheckTheExistanceOfPhone(ctx context.Context) (int, error) {
	phone := ctx.Value("phone").(string)
	status := 0
	er := repo.DB.QueryRow(ctx, "select * from checkTheExistanceOfSubscriberByPhone($1)", phone).Scan(&status)
	if er != nil {
		return 0, er
	}
	return status, nil
}

// RemoveExpiredTempoSubscription
func (repo *SubscriberRepo) RemoveExpiredTempoSubscription(unixTime uint64) (numberOfDeletedConfirmations int, issue error) {
	uc, er := repo.DB.Exec(context.TODO(), "delete from tempo_subscriber where unix <=$1", unixTime)
	if er != nil {
		return int(uc.RowsAffected()), er
	}
	return int(uc.RowsAffected()), nil
}

func (repo *SubscriberRepo) GetSubscriberByPhone(ctx context.Context) (*model.Subscriber, int, error) {
	subscriber := &model.Subscriber{}
	phone := ctx.Value("subscriber_phone").(string)
	if er := repo.DB.QueryRow(ctx, "select id,fullname,lang,role,subscriptions from subscriber where phone=$1", phone).
		Scan(&(subscriber.ID), &(subscriber.Fullname), &(subscriber.Lang), &(subscriber.Role), &(subscriber.Subscriptions)); er != nil {
		return subscriber, state.DT_STATUS_DBQUERY_ERROR, er
	}
	subscriber.Phone = phone
	return subscriber, state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) DeleteTempoLoginSubscriber(unix uint64) (int, error) {
	uc, er := repo.DB.Exec(context.TODO(), "delete from tempo_subscribers_login where unix <=$1", unix)
	if er != nil || uc.RowsAffected() == 0 {
		if er != nil {
			return 0, er
		}
		return 0, fmt.Errorf("no rows affected")
	}
	return int(uc.RowsAffected()), nil
}
func (repo *SubscriberRepo) RegisterSubscriber(ctx context.Context) (*model.Subscriber, int, error) {
	subscriber := ctx.Value("subscriber").(*model.Subscriber)
	er := repo.DB.QueryRow(ctx, "insert into subscriber( fullname,Phone,lang,role) values( $1,$2,$3,$4) returning id",
		subscriber.Fullname, subscriber.Phone, subscriber.Lang, subscriber.Role).Scan(&(subscriber.ID))
	if er != nil {
		return subscriber, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return subscriber, state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) RegisterTempoLoginSubcriber(ctx context.Context) (int, error) {
	tempo := ctx.Value("login_tempo_subscriber").(*model.TempoLoginSubscriber)
	er := repo.DB.QueryRow(ctx, "insert into tempo_subscribers_login( phone, confirmation, unix) values( $1,$2,$3 ) returning id",
		tempo.Phone, tempo.Confirmation, tempo.Unix).Scan(&(tempo.ID))
	if er != nil {
		return state.DT_STATUS_DBQUERY_ERROR, er
	}
	return state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) GetPendingRegistrationSubscriptionByPhone(ctx context.Context) (*model.TempoSubscriber, int, error) {
	phone := ctx.Value("subscriber_phone").(string)
	subscriber := &model.TempoSubscriber{}
	if er := repo.DB.QueryRow(ctx, "select id,fullname,Phone,lang,role,confirmation,unix,trials from selectTempoSubscriberWithPhoneAndUpdatedTrials($1)", phone).
		Scan(&(subscriber.ID), &(subscriber.Fullname), &(subscriber.Phone), &(subscriber.Lang), &(subscriber.Role), &(subscriber.ConfirmationCode), &(subscriber.Unix), &(subscriber.Trials)); er != nil {
		println(er.Error())
		return subscriber, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return subscriber, state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) DeletePendingRegistrationSubscriptionByID(ctx context.Context) (int, error) {
	id := ctx.Value("subscription_id").(uint64)
	if uc, er := repo.DB.Exec(ctx, "delete from tempo_subscriber where id=$1", id); uc.RowsAffected() == 0 || er != nil {
		if er != nil {
			return state.DT_STATUS_DBQUERY_ERROR, er
		}
		return state.DT_STATUS_RECORD_NOT_FOUND, fmt.Errorf("no data to delete")
	}
	return state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) DeletePendingLoginSubscriptionByID(ctx context.Context) (int, error) {
	id := ctx.Value("subscription_id").(uint64)
	if uc, er := repo.DB.Exec(ctx, "delete from tempo_subscribers_login where id=$1", id); uc.RowsAffected() == 0 || er != nil {
		if er != nil {
			return state.DT_STATUS_DBQUERY_ERROR, er
		}
		return state.DT_STATUS_RECORD_NOT_FOUND, fmt.Errorf("no data to delete")
	}
	return state.DT_STATUS_OK, nil
}

func (repo *SubscriberRepo) GetPendingLoginSubscriptionByPhone(ctx context.Context) (*model.TempoLoginSubscriber, int, error) {
	phone := ctx.Value("subscriber_phone").(string)
	subscriber := &model.TempoLoginSubscriber{}
	if er := repo.DB.QueryRow(ctx, "select id,phone,confirmation,unix from tempo_subscribers_login where phone=$1", phone).
		Scan(&(subscriber.ID), &(subscriber.Phone), &(subscriber.Confirmation), &(subscriber.Unix)); er != nil {
		return subscriber, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return subscriber, state.DT_STATUS_OK, nil
}
