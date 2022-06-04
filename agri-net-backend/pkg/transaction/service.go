package transaction

import (
	"context"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/contract"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
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
	SellerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error
	BuyerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error
	GetTransactionNotifications(ctx context.Context, role string, userid uint64) ([]*model.TransactionNotification, int, error)
	GetTransactionNotificationByTransactionID(ctx context.Context, trid uint64) (*model.TransactionRequest, error)
	GetKebdNotificationByTransactionID(ctx context.Context, trid uint64) (*model.KebdAmountRequest, error)
	GetGuaranteeNotificationByTransactionID(ctx context.Context, trid uint64) (*model.GuaranteeAmountRequest, error)
	AmendKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error)
	CreateKebdAmendmentRequest(ctx context.Context, merchantid uint64, input *model.KebdAmountRequest) (int, error)
}

type TransactionService struct {
	Repo         ITransactionRepo
	PaymentRepo  payment.IPaymentRepository
	ContractRepo contract.IContractRepo
}

func NewTransactionService(
	repo ITransactionRepo,
	payRepo payment.IPaymentRepository,
	contractRepo contract.IContractRepo,
) ITransactionService {
	return &TransactionService{
		Repo:         repo,
		PaymentRepo:  payRepo,
		ContractRepo: contractRepo,
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

func (service *TransactionService) SellerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error {
	return service.Repo.SellerAcceptTransaction(ctx, sellerid, transactionID)
}

func (service *TransactionService) BuyerAcceptTransaction(ctx context.Context, sellerid, transactionID uint64) error {
	return service.Repo.BuyerAcceptTransaction(ctx, sellerid, transactionID)
}

func (service *TransactionService) GetTransactionNotifications(ctx context.Context, role string, userid uint64) ([]*model.TransactionNotification, int, error) {
	notifications := []*model.TransactionNotification{}
	transactions, er := service.Repo.GetMyActiveTransactions(ctx, userid)
	if er != nil {
		return nil, -1, er
	}
	for x := range transactions {
		notif := &model.TransactionNotification{
			TransactionID: uint64(transactions[x].ID),
		}
		switch transactions[x].State {
		case state.TS_AMENDMENT_REQUESTED, state.TS_AMENDED:
			notif.TransactionNotification, er = service.Repo.GetTransactionNotificationByTransactionID(ctx, uint64(transactions[x].ID))
			if er != nil {
				println(er.Error())
			}
		case state.TS_KEBD_REQUESTED, state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT, state.TS_KEBD_AMENDED:
			notif.KebdNotification, er = service.Repo.GetKebdNotificationByTransactionID(ctx, uint64(transactions[x].ID))
			if er != nil {
				println(er.Error())
			}
		case state.TS_GUARANTEE_AMOUNT_REQUEST_SENT, state.TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT, state.TS_GUARANTEE_AMOUNT_AMENDED:
			notif.GuaranteeNotification, er = service.Repo.GetGuaranteeNotificationByTransactionID(ctx, uint64(transactions[x].ID))
			if er != nil {
				println(er.Error())
			}
		case state.TS_PAYMENT_INSTANTIATED, state.TS_SELLER_PAYMENT_COMPLETED,
			state.TS_BUYER_PAYMENT_COMPLETED, state.TS_ERROR:
			notif.PaymentNotification, er = service.PaymentRepo.GetTransactionPaymentByTransactionID(ctx, uint64(transactions[x].ID))
			if er != nil {
				println(er.Error())
			}
		case state.TS_CONTRACT_CREATED, state.TS_CONTRACT_CREATED_ACTIVATED,
			state.TS_CONTRACT_FAILED, state.TS_CONTRACT_SUCCEED:
			notif.ContractNotification, er = service.ContractRepo.GetContractByTransactionID(ctx, uint64(transactions[x].ID))
		}
		notifications = append(notifications, notif)
	}
	return notifications, 0, nil
}

func (service *TransactionService) GetTransactionNotificationByTransactionID(ctx context.Context, trid uint64) (*model.TransactionRequest, error) {
	return service.Repo.GetTransactionNotificationByTransactionID(ctx, trid)
}
func (service *TransactionService) GetKebdNotificationByTransactionID(ctx context.Context, trid uint64) (*model.KebdAmountRequest, error) {
	return service.Repo.GetKebdNotificationByTransactionID(ctx, trid)
}
func (service *TransactionService) GetGuaranteeNotificationByTransactionID(ctx context.Context, trid uint64) (*model.GuaranteeAmountRequest, error) {
	return service.Repo.GetGuaranteeNotificationByTransactionID(ctx, trid)
}
func (service *TransactionService) AmendKebdRequest(ctx context.Context, cxpid uint64, input *model.KebdAmountRequest) (int, error) {
	return service.Repo.AmendKebdRequest(ctx, cxpid, input)
}

func (service *TransactionService) CreateKebdAmendmentRequest(ctx context.Context, merchantid uint64, input *model.KebdAmountRequest) (int, error) {
	return service.Repo.CreateKebdAmendmentRequest(ctx, merchantid, input)
}
