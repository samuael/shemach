package transaction

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ITransactionService interface {
	CreateNewTransaction(ctx context.Context, transaction *model.Transaction) error
	GetMyActiveTransactions(ctx context.Context, userid uint64) ([]*model.Transaction, error)
	DeclineTransaction(ctx context.Context, transactionid uint64, userid uint64) (int, error)
	GetTransactionByID(ctx context.Context, transactionid uint64) (*model.Transaction, error)
	SaveTransactionAmendmentRequest(ctx context.Context, transaction *model.TransactionRequest) error
	GetTransactionRequestByID(ctx context.Context, id uint64) (*model.TransactionRequest, error)
	AcceptTransactionAmendmentRequest(ctx context.Context, merchantid, requestid uint64) (int, error)
	SetTransactionAmendment(ctx context.Context, userid uint64, req *model.TransactionRequest) (int, error)
	CreateKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error)
	CreateGuaranteeRequest(ctx context.Context, cxpid uint64, input *model.GuaranteeAmountRequest) (int, error)
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
func (service *TransactionService) GetTransactionByID(ctx context.Context, transactionid uint64) (*model.Transaction, error) {
	return service.Repo.GetTransactionByID(ctx, transactionid)
}
func (service *TransactionService) SaveTransactionAmendmentRequest(ctx context.Context, transaction *model.TransactionRequest) error {
	return service.Repo.SaveTransactionAmendmentRequest(ctx, transaction)
}

func (service *TransactionService) GetTransactionRequestByID(ctx context.Context, id uint64) (*model.TransactionRequest, error) {
	return service.Repo.GetTransactionRequestByID(ctx, id)
}
func (service *TransactionService) AcceptTransactionAmendmentRequest(ctx context.Context, merchantid, requestid uint64) (int, error) {
	return service.Repo.AcceptTransactionAmendmentRequest(ctx, merchantid, requestid)
}

func (service *TransactionService) SetTransactionAmendment(ctx context.Context, userid uint64, req *model.TransactionRequest) (int, error) {
	return service.Repo.SetTransactionAmendment(ctx, userid, req)
}

func (service *TransactionService) CreateKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error) {
	return service.Repo.CreateKebdRequest(ctx, cxpid, input)
}
func (service *TransactionService) CreateGuaranteeRequest(ctx context.Context, cxpid uint64, input *model.GuaranteeAmountRequest) (int, error) {
	return service.Repo.CreateGuaranteeRequest(ctx, cxpid, input)
}
