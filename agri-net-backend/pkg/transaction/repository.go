package transaction

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type ITransactionRepo interface {
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
	SellerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error
	BuyerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error
	GetTransactionNotificationByTransactionID(ctx context.Context, trid uint64) (*model.TransactionRequest, error)
	GetKebdNotificationByTransactionID(ctx context.Context, trid uint64) (*model.KebdAmountRequest, error)
	GetGuaranteeNotificationByTransactionID(ctx context.Context, trid uint64) (*model.GuaranteeAmountRequest, error)
}
