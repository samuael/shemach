package transaction

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ITransactionService interface {
	CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error
	GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error)
	DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error)
}

type TransactionService struct {
	Repo ITransactionRepo
}

func NewTransactionService(repo ITransactionRepo) ITransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (service TransactionService) CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error {
	return service.Repo.CreateNewTransaction(ctx, transaction)
}

func (service TransactionService) GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error) {
	return service.Repo.GetMyActiveTransactions(ctx, userid)
}

func (service TransactionService) DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error) {
	return service.Repo.DeclineTransaction(ctx, transactionid, userid)
}
