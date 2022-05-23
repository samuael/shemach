package transaction

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ITransactionRepo interface {
	CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error
	GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error)
	DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error)
}
