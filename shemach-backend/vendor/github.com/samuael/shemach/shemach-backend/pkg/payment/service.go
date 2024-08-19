package payment

import (
	"context"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type IPaymentService interface {
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

	CreateTransactionPayment(ctx context.Context, tp *model.TransactionPayment) (int, error)
	GetTransactionPaymentByTransactionID(ctx context.Context, transactionID uint64) (*model.TransactionPayment, error)
	UpdateTransactionPaymentStateByTransactionID(ctx context.Context, transactionID uint, state uint) error
	GetPendingPayment(ctx context.Context) ([]*model.TransactionPayment, error)
	DeletePaymentByTransactionID(ctx context.Context, transactionID uint64) error
	//
	UpdatePaymentState(ctx context.Context, state uint8, transactionid uint64) error
}
type PaymentService struct {
	Repo IPaymentRepository
}

func NewPaymentService(repo IPaymentRepository) IPaymentService {
	return &PaymentService{
		Repo: repo,
	}
}
func (ser *PaymentService) Authenticate(ctx context.Context) error {
	return ser.Repo.Authenticate(ctx)
}
func (ser *PaymentService) ValidateInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error) {
	return ser.Repo.ValidateInvoice(ctx, invoice)
}
func (ser *PaymentService) SendAnInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error) {
	return ser.Repo.SendAnInvoice(ctx, invoice)
}
func (ser *PaymentService) DeleteAnInvoiceByID(ctx context.Context, invoiceID string) error {
	return ser.Repo.DeleteAnInvoiceByID(ctx, invoiceID)
}
func (ser *PaymentService) SendAuthHTTPRequest(ctx context.Context, path, method string, input interface{}, response interface{}, authenticated bool) (int, error) {
	return ser.Repo.SendAuthHTTPRequest(ctx, path, method, input, response, authenticated)
}
func (ser *PaymentService) GetInvoiceByID(ctx context.Context, invoiceID string) (*model.HelloCashInvoice, error) {
	return ser.Repo.GetInvoiceByID(ctx, invoiceID)
}
func (ser *PaymentService) GetInvoices(ctx context.Context, status model.InvoiceStatus) ([]*model.HelloCashInvoice, error) {
	return ser.Repo.GetInvoices(ctx, status)
}
func (ser *PaymentService) ValidateMoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error) {
	return ser.Repo.ValidateMoneyTransffer(ctx, transfer)
}
func (ser *PaymentService) MoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error) {
	return ser.Repo.MoneyTransffer(ctx, transfer)
}
func (ser *PaymentService) GetTransferPaymentByID(ctx context.Context, transferID string) (*model.HellocashTransffer, error) {
	return ser.Repo.GetTransferPaymentByID(ctx, transferID)
}
func (ser *PaymentService) GetTransferPayments(ctx context.Context) ([]*model.HellocashTransffer, error) {
	return ser.Repo.GetTransferPayments(ctx)
}

func (ser *PaymentService) CreateTransactionPayment(ctx context.Context, tp *model.TransactionPayment) (int, error) {
	return ser.Repo.CreateTransactionPayment(ctx, tp)
}

func (ser *PaymentService) GetTransactionPaymentByTransactionID(ctx context.Context, transactionID uint64) (*model.TransactionPayment, error) {
	return ser.Repo.GetTransactionPaymentByTransactionID(ctx, transactionID)
}

func (ser *PaymentService) UpdateTransactionPaymentStateByTransactionID(ctx context.Context, transactionID uint, state uint) error {
	return ser.Repo.UpdateTransactionPaymentStateByTransactionID(ctx, transactionID, state)
}

func (ser *PaymentService) GetPendingPayment(ctx context.Context) ([]*model.TransactionPayment, error) {
	return ser.Repo.GetPendingPayment(ctx)
}
func (ser *PaymentService) DeletePaymentByTransactionID(ctx context.Context, transactionID uint64) error {
	return ser.Repo.DeletePaymentByTransactionID(ctx, transactionID)
}
func (ser *PaymentService) UpdatePaymentState(ctx context.Context, state uint8, transactionid uint64) error {
	return ser.Repo.UpdatePaymentState(ctx, state, transactionid)
}
