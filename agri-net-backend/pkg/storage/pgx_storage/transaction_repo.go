package pgx_storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/transaction"
)

type TransactionRepo struct {
	DB *pgxpool.Pool
}

func NewTransactionRepo(conn *pgxpool.Pool) transaction.ITransactionRepo {
	return &TransactionRepo{
		DB: conn,
	}
}
