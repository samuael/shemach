package payment

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IPaymentRepository interface {
	Authenticate(ctx context.Context) error
	ValidateInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error)
	SendAnInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error)
	DeleteAnInvoiceByID(ctx context.Context, invoiceID string) error
	SendAuthHTTPRequest(ctx context.Context, path, method string, input interface{}, response interface{}, authenticated bool) (int, error)
	GetInvoiceByID(ctx context.Context, invoiceID string) (*model.HelloCashInvoice, error)
	GetInvoices(ctx context.Context, status model.InvoiceStatus) ([]*model.HelloCashInvoice, error)
	ValidateMoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error)
	MoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error)
	GetTransferPaymentByID(ctx context.Context, transferID string) (*model.HellocashTransffer, error)
	GetTransferPayments(ctx context.Context) ([]*model.HellocashTransffer, error)

	//
	CreateTransactionPayment(ctx context.Context, tp *model.TransactionPayment) (int, error)
	GetTransactionPaymentByTransactionID(ctx context.Context, transactionID uint64) (*model.TransactionPayment, error)
	UpdateTransactionPaymentStateByTransactionID(ctx context.Context, transactionID uint, state uint) error
	GetPendingPayment(ctx context.Context) ([]*model.TransactionPayment, error)
	DeletePaymentByTransactionID(ctx context.Context, transactionID uint64) error
	//
	UpdatePaymentState(ctx context.Context, state uint8, transactionid uint64) error
}
