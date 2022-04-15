package pgx_storage

import (
	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
)

type InfoadminRepo struct {
	DB *pgxpool.Pool
}

func NewInfoadminRepo(db *pgxpool.Pool) infoadmin.IInfoadminRepo {
	return &InfoadminRepo{
		DB: db,
	}
}
