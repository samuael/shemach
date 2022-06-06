package pgx_storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
)

// PaymentRepo
type PaymentRepo struct {
	DB          *pgxpool.Pool
	PaymentURL  string
	Credentials *model.Credentials
	Client      *http.Client
	AuthToken   string
}

type InvoiceStatus string

const (
	Initializing = InvoiceStatus("INITIALIZING")
	PENDING      = InvoiceStatus("PENDING")
	EXPIRED      = InvoiceStatus("EXPIRED")
	PROCESSED    = InvoiceStatus("PROCESSED")
	CANCELED     = InvoiceStatus("CANCELED")
	FAILED       = InvoiceStatus("FAILED")
	NONE         = InvoiceStatus("")
)

const (
	validateInvoiceRoute = "/invoices/validate"
	invoices             = "/invoices"
	transfers            = "/transfers"
	authentication       = "/authenticate"
)

// AuthenticationReponse
type AuthenticationReponse struct {
	Token  string            `json:"token,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}

// NewPaymentRepo
func NewPaymentRepo(conn *pgxpool.Pool,
	credentials *model.Credentials,
) payment.IPaymentRepository {
	return &PaymentRepo{
		DB:          conn,
		Credentials: credentials,
		Client:      &http.Client{},
		PaymentURL:  "https://api-et.hellocash.net",
	}
}

func (repo *PaymentRepo) Authenticate(ctx context.Context) error {
	payload := strings.NewReader(string(func() []byte {
		if bytes, er := json.Marshal(repo.Credentials); er != nil {
			println(er.Error())
			return []byte{}
		} else {
			return bytes
		}
	}()))
	var res *http.Response
	var err error
	req, err := http.NewRequest(http.MethodPost, repo.PaymentURL+"/authenticate", payload)
	if err != nil {
		return err
	}
	req.Header.Add("accept", "text/plain")
	req.Header.Add("Content-Type", "application/json")
	res, err = repo.Client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	response := &AuthenticationReponse{}
	decoder := json.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(response)
	repo.AuthToken = response.Token
	if repo.AuthToken == "" {
		return errors.New("can't authenticate")
	}
	return err
}

// ValidateInvoice
func (repo *PaymentRepo) ValidateInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error) {
	if invoice.Amount <= 0 {
		return nil, errors.New("missing amount")
	} else if invoice.Currency != "ETB" {
		return nil, errors.New("invalid Currency")
	} else if !(helper.MatchesPattern(invoice.From, helper.PhoneRX) && strings.HasPrefix(invoice.From, "+251")) {
		return nil, errors.New("invalid phone number information")
	} else if invoice.Description == "" {
		return nil, errors.New("missing description body")
	}
	var response model.HelloCashInvoice
	stcode, err := repo.SendAuthHTTPRequest(ctx, validateInvoiceRoute, http.MethodPost, invoice, &response, true)
	if err != nil {
		return nil, err
	} else if stcode == http.StatusInternalServerError {
		return nil, errors.New("internal server error message")
	} else if stcode == http.StatusBadRequest {
		return nil, errors.New("status bad request ")
	} else if stcode == http.StatusExpectationFailed {
		return nil, errors.New("expectation failed")
	}
	if response.Fromname == "undefined" {
		return nil, errors.New("undefined account")
	}
	return &response, nil
}

// ValidateInvoice
func (repo *PaymentRepo) SendAnInvoice(ctx context.Context, invoice *model.HellocashInvoiceRequest) (*model.HelloCashInvoice, error) {
	if invoice.Amount <= 0 {
		return nil, errors.New("missing amount information")
	} else if invoice.Currency != "ETB" {
		return nil, errors.New("invalid Currency")
	} else if !(helper.MatchesPattern(invoice.From, helper.PhoneRX) && strings.HasPrefix(invoice.From, "+251")) {
		return nil, errors.New("invalid phone number information")
	} else if invoice.Description == "" {
		return nil, errors.New("missing description body")
	}
	var response model.HelloCashInvoice
	stcode, err := repo.SendAuthHTTPRequest(ctx, invoices, http.MethodPost, invoice, &response, true)
	if err != nil {
		return nil, err
	} else if stcode == http.StatusInternalServerError {
		return nil, errors.New("internal server error message")
	} else if stcode == http.StatusBadRequest {
		return nil, errors.New("status bad request ")
	} else if stcode == http.StatusExpectationFailed {
		return nil, errors.New("expectation failed")
	}
	if response.Fromname == "undefined" {
		return nil, errors.New("undefined account")
	}
	return &response, nil
}

// DeleteAnInvoiceByID
func (repo *PaymentRepo) DeleteAnInvoiceByID(ctx context.Context, invoiceID string) error {
	val := false
	stcode, err := repo.SendAuthHTTPRequest(ctx, invoices+"/"+invoiceID, http.MethodDelete, nil, &val, true)
	if err != nil {
		return err
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return errors.New("expectation failed")
		} else {
			return errors.New("request was not succesful")
		}
	}
	return nil
}

// SendAuthHTTPRequest
func (repo *PaymentRepo) SendAuthHTTPRequest(ctx context.Context, path, method string, input interface{}, response interface{}, authenticated bool) (int, error) {
	var payload *strings.Reader
	if input != nil {
		payload = strings.NewReader(string(func() []byte {
			if bytes, er := json.Marshal(input); er != nil {
				return []byte{}
			} else {
				return bytes
			}
		}()))
	} else {
		payload = strings.NewReader(string([]byte{}))
	}
	var res *http.Response
	var err error
	req, err := http.NewRequest(method, repo.PaymentURL+path, payload)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	if authenticated {
		req.Header.Add("Authorization", "Bearer "+repo.AuthToken)
	}
	res, err = repo.Client.Do(req)
	if err != nil {
		return res.StatusCode, err
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, err
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(response)
	return res.StatusCode, err
}

// GetInvoiceByID
func (repo *PaymentRepo) GetInvoiceByID(ctx context.Context, invoiceID string) (*model.HelloCashInvoice, error) {
	var response model.HelloCashInvoice
	stcode, er := repo.SendAuthHTTPRequest(ctx, invoices+"/"+invoiceID, http.MethodGet, nil, &response, true)
	if er != nil {
		return nil, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return nil, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return nil, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return nil, errors.New("expectation failed")
		} else {
			return nil, errors.New("request was not succesful")
		}
	}
	return &response, nil
}

// GetInvoices
func (repo *PaymentRepo) GetInvoices(ctx context.Context, status model.InvoiceStatus) ([]*model.HelloCashInvoice, error) {
	var response []*model.HelloCashInvoice
	stcode, er := repo.SendAuthHTTPRequest(ctx, invoices, http.MethodGet, nil, &response, true)
	if er != nil {
		return response, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return response, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return response, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return response, errors.New("expectation failed")
		} else {
			return response, errors.New("request was not succesful")
		}
	}
	return response, nil
}

// ValidateMoneyTransffer
func (repo *PaymentRepo) ValidateMoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error) {
	if transfer.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}
	if !(helper.MatchesPattern(transfer.To, helper.PhoneRX)) {
		return nil, errors.New("receiver phone is invalid")
	}
	if !(transfer.Currency != "ETB") {
		return nil, errors.New("invalid currency")
	}
	var response model.HellocashTransffer
	stcode, er := repo.SendAuthHTTPRequest(ctx, transfers+"/"+"validate", http.MethodPost, transfer, &response, true)
	if er != nil {
		return nil, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return nil, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return nil, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return nil, errors.New("expectation failed")
		} else {
			return nil, errors.New("request was not succesful")
		}
	}
	return &response, nil
}

// MoneyTransffer
func (repo *PaymentRepo) MoneyTransffer(ctx context.Context, transfer *model.HellocashTransfferRequest) (*model.HellocashTransffer, error) {
	if transfer.Amount <= 0 {
		return nil, errors.New("invalid amount")
	}
	if !(helper.MatchesPattern(transfer.To, helper.PhoneRX)) {
		return nil, errors.New("receiver phone is invalid")
	}
	if !(transfer.Currency != "ETB") {
		return nil, errors.New("invalid currency")
	}
	var response model.HellocashTransffer
	stcode, er := repo.SendAuthHTTPRequest(ctx, transfers, http.MethodPost, transfer, &response, true)
	if er != nil {
		return nil, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return nil, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return nil, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return nil, errors.New("expectation failed")
		} else {
			return nil, errors.New("request was not succesful")
		}
	}
	return &response, nil
}

// GetTransferPaymentByID
func (repo *PaymentRepo) GetTransferPaymentByID(ctx context.Context, transferID string) (*model.HellocashTransffer, error) {
	if transferID == "" {
		return nil, errors.New("invalid transfer id")
	}
	var response model.HellocashTransffer
	stcode, er := repo.SendAuthHTTPRequest(ctx, transfers+"/"+transferID, http.MethodGet, nil, &response, true)
	if er != nil {
		return nil, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return nil, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return nil, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return nil, errors.New("expectation failed")
		} else {
			return nil, errors.New("request was not succesful")
		}
	}
	return &response, nil
}

// DeleteTransferPaymentByID
func (repo *PaymentRepo) DeleteTransferPaymentByID(ctx context.Context, transferID string) (bool, error) {
	if transferID == "" {
		return false, errors.New("invalid transfer id")
	}
	response := false
	stcode, er := repo.SendAuthHTTPRequest(ctx, transfers+"/"+transferID, http.MethodGet, nil, &response, true)
	if er != nil {
		return false, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return false, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return false, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return false, errors.New("expectation failed")
		} else if stcode == http.StatusNotFound {
			return false, errors.New("transfer not found")
		} else {
			return false, errors.New("request was not succesful")
		}
	}
	return response, nil
}

// GetTransferPayments
func (repo *PaymentRepo) GetTransferPayments(ctx context.Context) ([]*model.HellocashTransffer, error) {
	var response []*model.HellocashTransffer
	stcode, er := repo.SendAuthHTTPRequest(ctx, transfers, http.MethodGet, nil, &response, true)
	if er != nil {
		return nil, er
	} else if stcode != http.StatusOK {
		if stcode != http.StatusInternalServerError {
			return nil, errors.New("internal server error message")
		} else if stcode == http.StatusBadRequest {
			return nil, errors.New("status bad request ")
		} else if stcode == http.StatusExpectationFailed {
			return nil, errors.New("expectation failed")
		} else {
			return nil, errors.New("request was not succesful")
		}
	}
	return response, nil
}

// CreateTransactionPayment
func (repo *PaymentRepo) CreateTransactionPayment(ctx context.Context, tp *model.TransactionPayment) (int, error) {
	status := 0
	er := repo.DB.QueryRow(ctx, `select * from reigisterNewTransactionPayment( 
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	 )`, tp.TransactionID, tp.SellerID, tp.SellerInvoiceID,
		tp.BuyerID, tp.SellerInvoiceID, tp.KebdAmount,
		tp.GuaranteeAmount).Scan(
		&status,
	)
	if status > 0 {
		tp.ID = uint64(status)
	}
	return status, er
}

// GetTransactionPaymentByTransactionID
func (repo *PaymentRepo) GetTransactionPaymentByTransactionID(ctx context.Context, transactionID uint64) (*model.TransactionPayment, error) {
	t := &model.TransactionPayment{}
	er := repo.DB.QueryRow(ctx, `select transaction_payment_info_id
	,transaction_id
	,state
	,created_at
	,seller_id
	,seller_invoice_id
	,buyer_id
	,buyer_invoice_id
	,kebd_amount
	,guarantee_amount
	,kebd_completed
	,guarantee_completed from transaction_payment_info where transaction_id=$1`, transactionID).Scan(
		&(t.ID),
		&(t.TransactionID),
		&(t.State),
		&(t.CreatedAt),
		&(t.SellerID),
		&(t.SellerInvoiceID),
		&(t.BuyerID),
		&(t.BuyerInvoiceID), //
		&(t.KebdAmount),
		&(t.GuaranteeAmount),
		&(t.KebdCompleted),
		&(t.GuaranteeCompleted),
	)
	if er != nil {
		return nil, er
	}
	return t, nil
}

// UpdateTransactionPaymentStateByTransactionID
func (repo *PaymentRepo) UpdateTransactionPaymentStateByTransactionID(ctx context.Context, transactionID uint, state uint) error {
	status := 0
	er := repo.DB.QueryRow(ctx, "select * from updateTransactionPaymentState($1,$2)", state, transactionID).
		Scan(&state)
	if status == 0 {
		return errors.New("no transaction payment instance was found")
	}
	return er
}

// GetPendingPayment
func (repo *PaymentRepo) GetPendingPayment(ctx context.Context) ([]*model.TransactionPayment, error) {
	payments := []*model.TransactionPayment{}
	rows, er := repo.DB.Query(ctx, `select transaction_payment_info_id
	,transaction_id
	,state
	,created_at
	,seller_id
	,seller_invoice_id
	,buyer_id
	,buyer_invoice_id
	,kebd_amount
	,guarantee_amount
	,kebd_completed
	,guarantee_completed from transaction_payment_info`)
	if er != nil {
		return nil, er
	}
	for rows.Next() {
		p := &model.TransactionPayment{}
		er := rows.Scan(
			&(p.ID),
			&(p.TransactionID),
			&(p.State),
			&(p.CreatedAt),
			&(p.SellerID),
			&(p.SellerInvoiceID),
			&(p.BuyerID),
			&(p.BuyerInvoiceID),
			&(p.KebdAmount),
			&(p.GuaranteeAmount),
			&(p.KebdCompleted),
			&(p.GuaranteeCompleted),
		)
		if er != nil {
			continue
		}
		payments = append(payments, p)
	}
	return payments, nil
}

// DeletePaymentByTransactionID
func (repo *PaymentRepo) DeletePaymentByTransactionID(ctx context.Context, transactionID uint64) error {
	uc, er := repo.DB.Exec(ctx, "delete from transaction_payment_info where transaction_id=$1", transactionID)
	if uc.RowsAffected() == 0 {
		return errors.New("no row was affected")
	}
	return er
}

// UpdatePaymentState
func (repo *PaymentRepo) UpdatePaymentState(ctx context.Context, state uint8, transactionid uint64) error {
	er := repo.DB.QueryRow(ctx, "update transaction set state=$1 where transaction_id=$2 returning transaction_id", state, transactionid).Scan(&transactionid)
	return er
}
