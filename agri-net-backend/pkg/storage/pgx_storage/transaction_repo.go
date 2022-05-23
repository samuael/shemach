package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
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

func (repo *TransactionRepo) CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error {
	return repo.DB.QueryRow(ctx, `insert into transaction(
		price,
		quantity,
		state,
		description,
		crop_id,
		requester_id,
		requester_store_ref,
		seller_id,
		seller_store_ref
	) values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning transaction_id`, transaction.RequestingPrice, transaction.Quantity, transaction.State,
		transaction.Description, transaction.ProductID, transaction.RequesterID,
		transaction.RequesterStoreRef, transaction.SellerID, transaction.SellerStoreRef).Scan(&(transaction.ID))
}

// GetMyActiveTransactions
func (repo *TransactionRepo) GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error) {
	transactions := []*model.Transaction{}
	rows, er := repo.DB.Query(ctx, `select transaction_id,price,quantity,state,deadline,description,crop_id,requester_id,
	requester_store_ref,seller_id,seller_store_ref,created_at,kebd_amount,guarantee_amount from transaction where requester_id=$1 or seller_id=$1`, userid)
	if er != nil {
		return nil, er
	}
	for rows.Next() {
		t := &model.Transaction{}
		er := rows.Scan(
			&(t.ID), &(t.RequestingPrice), &(t.Quantity), &(t.State), &(t.Deadline), &(t.Description), &(t.ProductID), &(t.RequesterID),
			&(t.RequesterStoreRef), &(t.SellerID), &(t.SellerStoreRef), &(t.CreatedAt), &(t.KebdAmount),
			&(t.GuaranteeAmount))
		if er != nil {
			continue
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// DeclineTransaction
func (repo *TransactionRepo) DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error) {
	var state int
	er := repo.DB.QueryRow(ctx, "select * from declineTransaction($1,$2)", transactionid, userid).Scan(&state)
	return state, er
}
