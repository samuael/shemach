package pgx_storage

import (
	"context"
	"errors"
	"strconv"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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

// SaveMessage ...
func (repo MessageRepo) SaveMessage(ctx context.Context) (*model.Message, int, error) {
	message := ctx.Value("message").(*model.Message)
	er := repo.DB.QueryRow(ctx, `insert into messages(targets,lang,data,created_by) values($1,$2,$3,$4) returning id`,
		message.Targets, message.Lang, message.Data, message.CreatedBy).Scan(&(message.ID))
	if er != nil {
		return message, state.STATUS_DBQUERY_ERROR, er
	}
	return message, state.STATUS_OK, nil
}

// GetMessagedBy ... ||
func (repo *MessageRepo) GetRecentMessages(ctx context.Context) ([]*model.Message, int, error) {
	offset := ctx.Value("offset").(uint)
	limit := ctx.Value("limit").(uint)
	unixTime := ctx.Value("unix_time").(uint64)
	subscriptions := ctx.Value("subscriptions").([]int)
	lang := ctx.Value("lang").(string)
	// ------------------------------
	var rows pgx.Rows
	var er error

	messages := []*model.Message{}
	if unixTime == 0 {
		tobeappend := " "
		if len(subscriptions) > 0 {
			for _, val := range subscriptions {
				tobeappend += "or " + strconv.Itoa(val) + " = any( targets ) "
			}
		}
		tobeappend += ")"
		// println(`select id,targets,lang,data,created_by,created_at from messages  where ( lang='all' or lang=$3 ) and ( -1 = any(targets) ` + tobeappend + ` offset $1 limit $2`)
		rows, er = repo.DB.Query(ctx, `select id,targets,lang,data,created_by,created_at from messages  where ( lang='all' or lang=$3 ) and ( -1 = any(targets) `+tobeappend+`  ORDER BY id DESC offset $1 limit $2`, offset, limit, lang)
	} else {
		tobeappend := " "
		if len(subscriptions) > 0 {
			for _, val := range subscriptions {
				tobeappend += " or " + strconv.Itoa(val) + "= any( targets ) "
			}
		}
		tobeappend += ")"
		// println(`select id,targets,lang,data,created_by,created_at from messages where ( lang='all' or lang=$4 ) and created_at > $3 and ( -1 = any(targets) ` + tobeappend + `  offset $1 limit $2`)
		rows, er = repo.DB.Query(ctx, `select id,targets,lang,data,created_by,created_at from messages where ( lang='all' or lang=$4 ) and created_at > $3 and ( -1 = any(targets) `+tobeappend+` ORDER BY id DESC  offset $1 limit $2 `, offset, limit, unixTime, lang)
	}
	if er != nil || rows == nil {
		if er != nil {
			println(er.Error())
		}
		return messages, state.STATUS_DBQUERY_ERROR, er
	}
	for rows.Next() {
		message := &model.Message{}
		erf := rows.Scan(&(message.ID), &(message.Targets), &(message.Lang), &(message.Data), &(message.CreatedBy), &(message.CreatedAt))
		if erf != nil {
			continue
		}
		messages = append(messages, message)
	}
	if len(messages) == 0 {
		return messages, state.STATUS_NO_RECORD_FOUND, errors.New("no message instance was fetched")
	}
	return messages, state.STATUS_OK, nil
}
