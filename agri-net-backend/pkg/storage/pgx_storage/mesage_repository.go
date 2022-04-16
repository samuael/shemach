package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
)

type MessageRepo struct {
	DB *pgxpool.Pool
}

func NewMessageRepo(conn *pgxpool.Pool) message.IMessageRepo {
	return &MessageRepo{
		DB: conn,
	}
}

func (repo MessageRepo) SaveMessage(ctx context.Context) (*model.Message, int, error) {
	message := ctx.Value("message").(*model.Message)
	er := repo.DB.QueryRow(ctx, `insert into message(targets,lang,data,created_by) values($1,$2,$3,$4) returning id`,
		message.Targets, message.Lang, message.Data, message.CreatedBy).Scan(&(message.ID))
	if er != nil {
		return message, state.STATUS_DBQUERY_ERROR, er
	}
	return message, state.STATUS_OK, nil
}
