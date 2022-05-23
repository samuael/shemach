package pgx_storage

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/dictionary"
)

type DictionaryRepo struct {
	DB *pgxpool.Pool
}

func NewDictionaryRepo(conn *pgxpool.Pool) dictionary.IDictionaryRepo {
	return &DictionaryRepo{
		DB: conn,
	}
}

func (repo *DictionaryRepo) NewDictionary(ctx context.Context, dict *model.Dictionary) error {
	// filtering
	dict.Lang = strings.Trim(strings.ToLower(dict.Lang), " ")
	dict.Text = strings.Trim(strings.ToLower(dict.Text), " ")
	dict.Translation = strings.Trim(strings.ToLower(dict.Translation), " ")
	if dict.Lang == "" || dict.Text == "" || dict.Translation == "" {
		return errors.New("missing important data")
	}
	er := repo.DB.QueryRow(ctx, "select * from createDictionary($1,$2,$3)", dict.Lang, dict.Text, dict.Translation).Scan(&(dict.ID))
	if er != nil {
		return er
	}
	return nil
}
func (repo *DictionaryRepo) Translate(ctx context.Context, dict *model.Dictionary) error {
	dict.Lang = strings.Trim(strings.ToLower(dict.Lang), " ")
	dict.Text = strings.Trim(strings.ToLower(dict.Text), " ")
	dict.Translation = strings.Trim(strings.ToLower(dict.Translation), " ")
	if dict.Lang == "" || dict.Text == "" {
		return errors.New("missing important data")
	}
	println(dict.Lang, dict.Text)
	er := repo.DB.QueryRow(ctx, "select id , translation from dictionary where lang	=$1 and sentence=$2", dict.Lang, dict.Text).Scan(&(dict.ID), &(dict.Translation))
	if er != nil {
		return er
	}
	return nil
}
func (repo *DictionaryRepo) DeleteTranslation(ctx context.Context, dict *model.Dictionary) (int, error) {
	dict.Lang = strings.Trim(strings.ToLower(dict.Lang), " ")
	dict.Text = strings.Trim(strings.ToLower(dict.Text), " ")
	dict.Translation = strings.Trim(strings.ToLower(dict.Translation), " ")
	raff, er := repo.DB.Exec(ctx, "delete from dictionary where id=$1", dict.ID)
	dict.ID = 0
	var erf error
	var raff2 pgconn.CommandTag
	if dict.Lang != "" && dict.Text != "" {
		erf = repo.DB.QueryRow(ctx, "delete from dictionary where lang=$1 and sentence=$2 returning id", dict.Lang, dict.Text).Scan(&(dict.ID))
	}
	if er != nil && erf != nil {
		return 0, erf
	}
	return int(raff.RowsAffected() + raff2.RowsAffected()), nil
}
func (repo *DictionaryRepo) UpdateTranslation(ctx context.Context, dict *model.Dictionary) error {
	dict.Lang = strings.Trim(strings.ToLower(dict.Lang), " ")
	dict.Text = strings.Trim(strings.ToLower(dict.Text), " ")
	dict.Translation = strings.Trim(strings.ToLower(dict.Translation), " ")
	uc, er := repo.DB.Exec(ctx, "update dictionary set lang=$1, sentence=$2, translation=$3 where id=$4", dict.Lang, dict.Text, dict.Translation, dict.ID)
	if er != nil || uc.RowsAffected() == 0 {
		if er != nil {
			return er
		}
		return errors.New("no row is affected")
	}
	return nil
}

func (repo *DictionaryRepo) GetDictionaries(ctx context.Context, offset, limit uint) ([]*model.Dictionary, error) {
	dictionaries := []*model.Dictionary{}
	rows, er := repo.DB.Query(ctx, "select id,sentence ,lang ,translation from dictionary ORDER BY id DESC OFFSET $1 LIMIT $2 ", offset, limit)
	if er != nil {
		return nil, er
	}
	for rows.Next() {
		dict := &model.Dictionary{}
		era := rows.Scan(&(dict.ID), &(dict.Text), &(dict.Lang), &(dict.Translation))
		if era != nil {
			continue
		}
		dictionaries = append(dictionaries, dict)
	}
	return dictionaries, nil
}
