package pgx_storage

import (
	"context"
	"errors"

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
func (repo *SubscriberRepo) RegisterSubscriber(ctx context.Context) (*model.TempoSubscriber, int, error) {
	tempo := ctx.Value("tempo_subscriber").(*model.TempoSubscriber)
	er := repo.DB.QueryRow(ctx, "Insert into tempo_subscriber(fullname,Phone, lang, role,confirmation ) values( $1,$2,$3,$4,$5) returning id", tempo.Fullname, tempo.Phone, tempo.Lang, tempo.Role, tempo.ConfirmationCode, tempo.Unix).Scan(&(tempo.ID))
	if er != nil {
		return tempo, state.DT_STATUS_DBQUERY_ERROR, er
	}
	return tempo, state.DT_STATUS_OK, nil
}

// CheckTheExistanceOfPhone
func (repo *SubscriberRepo) CheckTheExistanceOfPhone(ctx context.Context) (int, error) {
	phone := ctx.Value("phone").(string)
	status := 0
	er := repo.DB.QueryRow(ctx, "select * from checkTheExistanceOfSubscriberByPhone($1)", phone).Scan(&status)
	if er != nil {
		return 0, errors.New("query error")
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
