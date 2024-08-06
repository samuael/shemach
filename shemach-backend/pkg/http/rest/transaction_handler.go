package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/crop"
	"github.com/samuael/shemach/shemach-backend/pkg/merchant"
	"github.com/samuael/shemach/shemach-backend/pkg/payment"
	"github.com/samuael/shemach/shemach-backend/pkg/store"
	"github.com/samuael/shemach/shemach-backend/pkg/user"

	"github.com/samuael/shemach/shemach-backend/pkg/transaction"
	"github.com/samuael/shemach/shemach-backend/platforms/translation"
)

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetMyActiveTransactions(c *gin.Context)
	DeclineTransaction(c *gin.Context)
	TransactionAmendmenRequest(c *gin.Context)
	AcceptAmendmentRequest(c *gin.Context)
	PerformAmend(c *gin.Context)
	RequestKebd(c *gin.Context)
	RequestKebdRequestAmendment(c *gin.Context)
	AmendKebdRequest(c *gin.Context)
	//
	RequestGuaranteePayment(c *gin.Context)
	SellerAcceptedTransaction(c *gin.Context)
	BuyerAcceptTransaction(c *gin.Context)
	ReactivateTransaction(c *gin.Context)
	GetMyTransactionNotifications(c *gin.Context)
}

// TransactionHandler transaction handler instance
type TransactionHandler struct {
	Service         transaction.ITransactionService
	UserService     user.IUserService
	ProductService  crop.ICropService
	MerchantService merchant.IMerchantService
	StoreService    store.IStoreService
	PaymentService  payment.IPaymentService
}

func NewTransactionHandler(
	service transaction.ITransactionService,
	userService user.IUserService,
	productService crop.ICropService,
	merchantService merchant.IMerchantService,
	storeService store.IStoreService,
	paymentService payment.IPaymentService,
) ITransactionHandler {
	return &TransactionHandler{
		Service:         service,
		UserService:     userService,
		ProductService:  productService,
		MerchantService: merchantService,
		StoreService:    storeService,
		PaymentService:  paymentService,
	}
}

// CreateTransaction creates a transaction instance using on a specified product.
func (thandler *TransactionHandler) CreateTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &struct {
		RequestingPrice   float64 `json:"price"`
		Quantity          uint64  `json:"qty"`
		Description       string  `json:"description"`
		ProductID         uint64  `json:"product_id"`
		RequesterStoreRef uint64  `json:"requester_store_ref"`
	}{}
	resp := &struct {
		Msg         string             `json:"msg"`
		StatusCode  int                `json:"status_code"`
		Transaction *model.Transaction `json:"transaction,omitempty"`
		Errors      map[string]string  `json:"errors,omitempty"`
	}{
		Errors: map[string]string{},
	}
	jdec := json.NewDecoder(c.Request.Body)
	er := jdec.Decode(input)
	if er != nil ||
		input.RequestingPrice <= 0 ||
		input.Quantity <= 0 ||
		input.ProductID <= 0 ||
		input.RequesterStoreRef <= 0 {
		if input.RequestingPrice <= 0 {
			resp.Errors["price"] = translation.Translate(session.Lang, "invalid requesting price \"price\" has to be >= 0")
		}
		if input.Quantity <= 0 {
			resp.Errors["qty"] = translation.Translate(session.Lang, "invalid quantity of amount \"qty\"")
		}
		if input.ProductID <= 0 {
			resp.Errors["product_id"] = translation.Translate(session.Lang, "product id must be greater than 0")
		}
		if input.RequesterStoreRef <= 0 {
			resp.Errors["requester_store_id"] = translation.Translate(session.Lang, "valid store id has to be specified")
		}
		resp.Msg = translation.Translate(session.Lang, "bad transaction request body")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	product, er := thandler.ProductService.GetPostByID(ctx, input.ProductID)
	if er != nil {
		resp.Msg = translation.Translate(session.Lang, "product instance with this id doesn't exist"+er.Error())
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	if product.Closed {
		resp.Msg = translation.Translate(session.Lang, "product with this id is not active for sell")
		resp.StatusCode = http.StatusExpectationFailed
		c.JSON(resp.StatusCode, resp)
		return
	}
	if product.RemainingQuantity < input.Quantity {
		resp.Msg = translation.Translate(session.Lang, "there is no storage to satisfy such amount of need.")
		resp.StatusCode = http.StatusExpectationFailed
		c.JSON(resp.StatusCode, resp)
		return
	}
	var SellerID uint64
	if !(product.StoreOwned) {
		SellerID = product.AgentID
	} else {
		store, er := thandler.StoreService.GetStoreByID(ctx, product.StoreID)
		if er == nil {
			SellerID = store.OwnerID
		} else {
			println(er.Error())
			resp.Msg = translation.Translate(session.Lang, "Internal Problem Please try again")
			resp.StatusCode = http.StatusInternalServerError
			c.JSON(resp.StatusCode, resp)
			return
		}
	}
	transaction := &model.Transaction{
		RequestingPrice:   input.RequestingPrice,
		Quantity:          input.Quantity,
		State:             state.TS_CREATED,
		Description:       input.Description,
		ProductID:         input.ProductID,
		RequesterID:       session.ID,
		RequesterStoreRef: input.RequesterStoreRef,
		SellerID:          SellerID,
		SellerStoreRef:    product.StoreID,
		CreatedAt:         uint64(time.Now().Unix()), //
	}
	er = thandler.Service.CreateNewTransaction(ctx, transaction)
	if er != nil {
		println(er.Error())
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "transaction created succesfuly")
	resp.Transaction = transaction
	c.JSON(resp.StatusCode, resp)
}

func (thandler *TransactionHandler) GetMyActiveTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	session, ok := ctx.Value("session").(*model.Session)
	if !ok {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	res := &struct {
		StatusCode   int                  `json:"status_code"`
		Msg          string               `json:"msg"`
		Transactions []*model.Transaction `json:"transactions,omitempty"`
	}{}
	transactions, er := thandler.Service.GetMyActiveTransactions(ctx, session.ID)
	if er != nil {
		println(er.Error())
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "no transacion found")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Transactions = transactions
	res.StatusCode = http.StatusOK
	res.Msg = translation.Translate(session.Lang, "found transactions")
	c.JSON(res.StatusCode, res)
}

func (thandler *TransactionHandler) DeclineTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	transactionID, er := strconv.Atoi(c.Param("id"))
	res := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg"`
	}{}
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "invalid transaction id")
		c.JSON(res.StatusCode, res)
		return
	}
	state, er := thandler.Service.DeclineTransaction(ctx, uint64(transactionID), session.ID)
	if state < 0 || er != nil {
		if state == -1 {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.Translate(session.Lang, "transaction not found")
		} else if state == -2 {
			res.StatusCode = http.StatusUnauthorized
			res.Msg = translation.Translate(session.Lang, "transaction is already in contract state")
		} else if state == -3 {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.Translate(session.Lang, "no transaction deleted")
		} else {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		}
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.Translate(session.Lang, "Transaction Deleted Succesfuly")
	c.JSON(res.StatusCode, res)
}

func (thandler *TransactionHandler) TransactionAmendmenRequest(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.TransactionRequest{}
	jdecoder := json.NewDecoder(c.Request.Body)
	resp := &struct {
		Msg        string                    `json:"msg"`
		StatusCode int                       `json:"status_code"`
		Request    *model.TransactionRequest `json:"amendment_request,omitempty"`
		Errors     map[string]string         `json:"errors,omitempty"`
	}{}
	er := jdecoder.Decode(input)
	if er != nil || input.TransactionID <= 0 ||
		len(input.Description) < 3 ||
		len(input.Description) > 500 ||
		input.Quantity <= 0 || input.Price <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.TransactionID <= 0 {
			resp.Errors["transaction id"] = translation.Translate(session.Lang, "transaction with the specified id doesn't exist")
		}
		if len(input.Description) < 3 || len(input.Description) > 500 {
			resp.Errors["description"] = translation.Translate(session.Lang, "please write a description for the amendment request with a length of not more than 500 characters")
		}
		if input.Quantity <= 0 {
			resp.Errors["quantity"] = translation.Translate(session.Lang, "invalid quantity")
		}
		if input.Price <= 0 {
			resp.Errors["price"] = translation.Translate(session.Lang, "invalid unit price")
		}
		resp.Msg = translation.Translate(session.Lang, "bad message payload")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.State = state.TS_AMENDMENT_REQUESTED
	input.CreatedAt = uint64(time.Now().Unix())
	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er != nil {
		resp.Msg = translation.Translate(session.Lang, "transaction instance not found")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	} else if transaction.State > 3 {
		resp.Msg = translation.Translate(session.Lang, "can't amend basic transaction information at this stage")
		resp.StatusCode = http.StatusUnauthorized
		c.JSON(resp.StatusCode, resp)
		return
	} else if transaction.SellerID != session.ID {
		resp.Msg = translation.Translate(session.Lang, "you are not authorized to send transaction amendment request")
		resp.StatusCode = http.StatusUnauthorized
		c.JSON(resp.StatusCode, resp)
		return
	} else if input.Price == transaction.RequestingPrice && input.Quantity == float64(transaction.Quantity) {
		resp.StatusCode = http.StatusConflict
		resp.Msg = translation.Translate(session.Lang, "please specify your new recommended price or amount of product")
		c.JSON(resp.StatusCode, resp)
		return
	}
	era := thandler.Service.SaveTransactionAmendmentRequest(ctx, input)
	if era != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.StatusCode = http.StatusCreated
	resp.Msg = translation.Translate(session.Lang, "request sent")
	resp.Request = input
	c.JSON(resp.StatusCode, resp)
}

// AcceptAmendmentRequest
func (thandler *TransactionHandler) AcceptAmendmentRequest(c *gin.Context) {
	requestid, er := strconv.Atoi(c.Param("id"))
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	resp := &struct {
		StatusCode  int                `json:"status_code"`
		Msg         string             `json:"msg"`
		Transaction *model.Transaction `json:"transaction,omitempty"`
	}{}
	if er != nil || requestid <= 0 {
		resp.StatusCode = http.StatusBadRequest
		resp.Msg = translation.Translate(session.Lang, "bad request")
		c.JSON(resp.StatusCode, resp)
		return
	}
	transactionreq, er := thandler.Service.GetTransactionRequestByID(ctx, uint64(requestid))
	if er != nil || transactionreq == nil {
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "transaction request not found")
		c.JSON(resp.StatusCode, resp)
		return
	}
	state, er := thandler.Service.AcceptTransactionAmendmentRequest(ctx, session.ID, transactionreq.ID)
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if state < 0 {
		switch state {
		case -1:
			resp.Msg = translation.Translate(session.Lang, "unauthorized user")
			resp.StatusCode = http.StatusUnauthorized
		case -2, -3:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this operation")
			resp.StatusCode = http.StatusUnauthorized
		case -4:
			resp.Msg = translation.Translate(session.Lang, "request not found")
			resp.StatusCode = http.StatusNotFound
		case -5:
			resp.Msg = translation.Translate(session.Lang, "can't update the change")
			resp.StatusCode = http.StatusInternalServerError
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, transactionreq.TransactionID)
	if er == nil {
		resp.Transaction = transaction
	}
	resp.Msg = translation.Translate(session.Lang, "request accepted")
	resp.StatusCode = http.StatusOK
	c.JSON(resp.StatusCode, resp)
}

func (thandler *TransactionHandler) PerformAmend(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.TransactionRequest{}
	resp := &struct {
		StatusCode  int                `json:"status_code"`
		Msg         string             `json:"msg"`
		Errors      map[string]string  `json:"errors,omitempty"`
		Transaction *model.Transaction `json:"transaction,omitempty"`
	}{}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil || input.ID <= 0 || input.Price <= 0 || input.Quantity <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.ID <= 0 {
			resp.Errors["request id"] = translation.Translate(session.Lang, "request id has to be specified")
		}
		if input.Price <= 0 {
			resp.Errors["price"] = translation.Translate(session.Lang, "the new price has to be specified")
		}
		if input.Quantity <= 0 {
			resp.Errors["quantity"] = translation.Translate(session.Lang, "The new quantity of product has to be specified")
		}
		resp.Msg = translation.Translate(session.Lang, "bad request")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	status, er := thandler.Service.SetTransactionAmendment(ctx, session.ID, input)
	if er != nil {
		println(er.Error())
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if status < 0 {
		switch status {

		case -1:
			resp.Msg = translation.Translate(session.Lang, "amendment request not found")
			resp.StatusCode = http.StatusNotFound
		case -2:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this operation; only the buyer can.")
			resp.StatusCode = http.StatusUnauthorized
		case -3:
			resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
			resp.StatusCode = http.StatusInternalServerError
		case -4, -5, -6:
			resp.Msg = translation.Translate(session.Lang, "all amendment informations has to be specified")
			resp.StatusCode = http.StatusBadRequest
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er == nil {
		resp.Transaction = transaction
	}
	resp.Msg = translation.Translate(session.Lang, "transaction amended succefuly")
	resp.StatusCode = http.StatusOK
	c.JSON(resp.StatusCode, resp)
}

func (thandler *TransactionHandler) RequestKebd(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.KebdAmountRequest{}
	resp := &struct {
		StatusCode int                      `json:"status_code"`
		Msg        string                   `json:"msg"`
		Errors     map[string]string        `json:"errors,omitempty"`
		Kebd       *model.KebdAmountRequest `json:"kebd_request,omitempty"`
	}{}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil || input.KebdAmount <= 0 || input.Deadline <= 0 || input.TransactionID <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.KebdAmount <= 0 {
			resp.Errors["kebd_amount"] = translation.Translate(session.Lang, "invalid kedb amount")
		}
		if input.Deadline <= 0 {
			resp.Errors["deadline"] = translation.Translate(session.Lang, "invalid deadline timstamp")
		}
		if input.TransactionID <= 0 {
			resp.Errors["transaction_id"] = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		}
		resp.Msg = translation.Translate(session.Lang, "bad request")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	deadline := time.Unix(int64(input.Deadline), 0)
	if !(deadline.After(time.Now().Add(time.Hour * 6))) {
		resp.Msg = translation.Translate(session.Lang, "deadline has to be at least 6 hours after contract time")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er != nil || transaction == nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.State = state.TS_KEBD_REQUESTED
	status, er := thandler.Service.CreateKebdRequest(ctx, session.ID, input)
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if status < 0 {
		switch status {
		case -1:
			resp.Msg = translation.Translate(session.Lang, "transaction not found")
			resp.StatusCode = http.StatusNotFound
		case -2:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform kebd request at this stage")
			resp.StatusCode = http.StatusUnauthorized
		case -3:
			resp.Msg = translation.Translate(session.Lang, "kebd with similar information already exist")
			resp.StatusCode = http.StatusConflict
		case -4, -6:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this action")
			resp.StatusCode = http.StatusUnauthorized
		case -5:
			resp.Msg = translation.Translate(session.Lang, "internal server error")
			resp.StatusCode = http.StatusInternalServerError
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.ID = uint64(status)
	input.CreatedAt = uint64(time.Now().Unix())
	resp.Kebd = input
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "kebd created succesfuly")
	c.JSON(resp.StatusCode, resp)
}

// TODO: TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT
func (thandler *TransactionHandler) RequestKebdRequestAmendment(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.KebdAmountRequest{}
	resp := &struct {
		StatusCode int                      `json:"status_code"`
		Msg        string                   `json:"msg"`
		Errors     map[string]string        `json:"errors,omitempty"`
		Kebd       *model.KebdAmountRequest `json:"kebd_request,omitempty"`
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil || input.KebdAmount <= 0 || input.Deadline <= 0 || input.TransactionID <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.KebdAmount <= 0 {
			resp.Errors["kebd_amount"] = translation.Translate(session.Lang, "invalid kedb amount")
		}
		if input.Deadline <= 0 {
			resp.Errors["deadline"] = translation.Translate(session.Lang, "invalid deadline timstamp")
		}
		if input.TransactionID <= 0 {
			resp.Errors["transaction_id"] = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		}
		resp.Msg = translation.Translate(session.Lang, "bad request")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	deadline := time.Unix(int64(input.Deadline), 0)
	if !(deadline.After(time.Now().Add(time.Hour * 6))) {
		resp.Msg = translation.Translate(session.Lang, "deadline has to be at least 6 hours after contract time")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er != nil || transaction == nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	} else if session.ID != transaction.RequesterID {
		resp.Msg = translation.Translate(session.Lang, "You are not authorized to perform this action")
		resp.StatusCode = http.StatusUnauthorized
		resp.Errors = map[string]string{}
		c.JSON(resp.StatusCode, resp)
		return
	} else if !(transaction.State == state.TS_KEBD_REQUESTED ||
		transaction.State == state.TS_KEBD_AMENDED ||
		transaction.State == state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT) {
		resp.Msg = translation.Translate(session.Lang, "operation is not allowed")
		resp.Errors["state"] = state.TransactionStateMaps[transaction.State]
		resp.StatusCode = http.StatusUnauthorized
		resp.Errors = map[string]string{}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.State = state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT
	status, er := thandler.Service.CreateKebdAmendmentRequest(ctx, session.ID, input)
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if status < 0 {
		switch status {
		case -1:
			resp.Msg = translation.Translate(session.Lang, "transaction not found")
			resp.StatusCode = http.StatusNotFound
		case -2:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this operation at this stage")
			resp.StatusCode = http.StatusUnauthorized
		case -3:
			resp.Msg = translation.Translate(session.Lang, "kebd with similar information already exist")
			resp.StatusCode = http.StatusConflict
		case -4, -6:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this action")
			resp.StatusCode = http.StatusUnauthorized
		case -5:
			resp.Msg = translation.Translate(session.Lang, "internal server error")
			resp.StatusCode = http.StatusInternalServerError
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.ID = uint64(status)
	input.CreatedAt = uint64(time.Now().Unix())
	resp.Kebd = input
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "kebd amendment request is sent succesfully")
	c.JSON(resp.StatusCode, resp)
}

// TODO: TS_KEBD_AMENDED
func (thandler *TransactionHandler) AmendKebdRequest(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.KebdAmountRequest{}
	resp := &struct {
		StatusCode int                      `json:"status_code"`
		Msg        string                   `json:"msg"`
		Errors     map[string]string        `json:"errors,omitempty"`
		Kebd       *model.KebdAmountRequest `json:"kebd_request,omitempty"`
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil || input.KebdAmount <= 0 || input.Deadline <= 0 || input.TransactionID <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.KebdAmount <= 0 {
			resp.Errors["kebd_amount"] = translation.Translate(session.Lang, "invalid kedb amount")
		}
		if input.Deadline <= 0 {
			resp.Errors["deadline"] = translation.Translate(session.Lang, "invalid deadline timstamp")
		}
		if input.TransactionID <= 0 {
			resp.Errors["transaction_id"] = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		}
		resp.Msg = translation.Translate(session.Lang, "bad request")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	deadline := time.Unix(int64(input.Deadline), 0)
	if !(deadline.After(time.Now().Add(time.Hour * 6))) {
		resp.Msg = translation.Translate(session.Lang, "deadline has to be at least 6 hours after contract time")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er != nil || transaction == nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	} else if session.ID != transaction.SellerID {
		resp.Msg = translation.Translate(session.Lang, "you are not a seller!\nYou are not authorized to perform this action")
		resp.StatusCode = http.StatusUnauthorized
		resp.Errors = map[string]string{}
		c.JSON(resp.StatusCode, resp)
		return
	} else if !(transaction.State == state.TS_KEBD_REQUESTED ||
		transaction.State == state.TS_KEBD_AMENDED ||
		transaction.State == state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT) {
		resp.Msg = translation.Translate(session.Lang, "operation is not allowed")
		resp.Errors["state"] = state.TransactionStateMaps[transaction.State]
		resp.StatusCode = http.StatusUnauthorized
		resp.Errors = map[string]string{}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.State = state.TS_KEBD_AMENDED
	status, er := thandler.Service.AmendKebdRequest(ctx, session.ID, input)
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if status < 0 {
		switch status {
		case -1:
			resp.Msg = translation.Translate(session.Lang, "transaction not found")
			resp.StatusCode = http.StatusNotFound
		case -2:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this operation at this stage")
			resp.StatusCode = http.StatusUnauthorized
		case -3:
			resp.Msg = translation.Translate(session.Lang, "kebd with similar information already exist")
			resp.StatusCode = http.StatusConflict
		case -4, -6:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this action")
			resp.StatusCode = http.StatusUnauthorized
		case -5:
			resp.Msg = translation.Translate(session.Lang, "internal server error")
			resp.StatusCode = http.StatusInternalServerError
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.ID = uint64(status)
	input.CreatedAt = uint64(time.Now().Unix())
	resp.Kebd = input
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "kebd request amended succesfully")
	c.JSON(resp.StatusCode, resp)
}

func (thandler *TransactionHandler) RequestGuaranteePayment(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &model.GuaranteeAmountRequest{}
	resp := &struct {
		StatusCode int                           `json:"status_code"`
		Msg        string                        `json:"msg"`
		Errors     map[string]string             `json:"errors,omitempty"`
		Guarantee  *model.GuaranteeAmountRequest `json:"guarantee,omitempty"`
	}{}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil || input.Amount <= 0 || input.TransactionID <= 0 {
		if er == nil {
			resp.Errors = map[string]string{}
		}
		if input.Amount <= 0 {
			resp.Errors["amount"] = translation.Translate(session.Lang, "invalid amount")
		}
		if input.TransactionID <= 0 {
			resp.Errors["transaction_id"] = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		}
		resp.Msg = translation.Translate(session.Lang, "bad request")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}

	transaction, er := thandler.Service.GetTransactionByID(ctx, input.TransactionID)
	if er != nil || transaction == nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id doesn't exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.State = state.TS_GUARANTEE_AMOUNT_REQUEST_SENT
	status, er := thandler.Service.CreateGuaranteeRequest(ctx, session.ID, input)
	if er != nil {
		println(er.Error())
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if status < 0 {
		switch status {
		case -1:
			resp.Msg = translation.Translate(session.Lang, "transaction not found")
			resp.StatusCode = http.StatusNotFound
		case -2:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform Guarantee request at this stage")
			resp.StatusCode = http.StatusUnauthorized
		case -3:
			resp.Msg = translation.Translate(session.Lang, "Guarantee with similar information already exist")
			resp.StatusCode = http.StatusConflict
		case -4, -6:
			resp.Msg = translation.Translate(session.Lang, "you are not allowed to perform this action")
			resp.StatusCode = http.StatusUnauthorized
		case -5:
			resp.Msg = translation.Translate(session.Lang, "internal server error")
			resp.StatusCode = http.StatusInternalServerError
		}
		c.JSON(resp.StatusCode, resp)
		return
	}
	input.ID = uint64(status)
	input.CreatedAt = uint64(time.Now().Unix())
	resp.Guarantee = input
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "Guarantee Payment Request created succesfuly")
	c.JSON(resp.StatusCode, resp)
}

// TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT
// TS_GUARANTEE_AMOUNT_AMENDED

func (thandler *TransactionHandler) SellerAcceptedTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	transactionID, er := strconv.Atoi(c.Param("id"))
	resp := &struct {
		StatusCode  int                `json:"status_code"`
		Msg         string             `json:"msg"`
		Errors      map[string]string  `json:"errors"`
		Transaction *model.Transaction `json:"transaction,omitempty"`
	}{}
	if er != nil || transactionID <= 0 {
		resp.Msg = translation.Translate(session.Lang, "invalid transaction id")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, uint64(transactionID))
	if er != nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id does not exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	if !(transaction.State != state.TS_GUARANTEE_AMOUNT_REQUEST_SENT && transaction.State != state.TS_ACCEPTED && transaction.State != state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT && transaction.State != state.TS_DECLINED) {
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "you are not authorized to perform this action in this stage")
		resp.Errors = map[string]string{}
		if transaction.State == state.TS_GUARANTEE_AMOUNT_REQUEST_SENT {
			resp.Errors["TS_GUARANTEE_AMOUNT_REQUEST_SENT"] = translation.Translate(session.Lang, "please complete the guarantee ayment process prior to accepting the transaction")
		} else if transaction.State == state.TS_ACCEPTED {
			resp.Errors["TS_ACCEPTED"] = translation.Translate(session.Lang, "transaction is already accepted")
		} else if transaction.State == state.TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT {
			resp.Errors["TS_KEBD_REQUEST_AMENDMENT_REQUEST_SENT"] = translation.Translate(session.Lang, "Please complete the kebd process first")
		}
		c.JSON(resp.StatusCode, resp)
		return
	} else if transaction.SellerID != session.ID {
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "unauthorized access")
		c.JSON(resp.StatusCode, resp)
		return
	}
	er = thandler.Service.SellerAcceptTransaction(ctx, session.ID, uint64(transactionID))
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.Transaction = transaction
	resp.Msg = translation.Translate(session.Lang, "transaction accepted")
	resp.StatusCode = http.StatusOK
	c.JSON(resp.StatusCode, resp)
}

func (thandler *TransactionHandler) BuyerAcceptTransaction(c *gin.Context) {
	session := c.Request.Context().Value("session").(*model.Session)
	ctx := context.Background()
	transactionID, er := strconv.Atoi(c.Param("id"))
	resp := &struct {
		StatusCode  int                `json:"status_code"`
		Msg         string             `json:"msg"`
		Errors      map[string]string  `json:"errors"`
		Transaction *model.Transaction `json:"transaction,omitempty"`
	}{}
	if er != nil || transactionID <= 0 {
		resp.Msg = translation.Translate(session.Lang, "invalid transaction id")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, uint64(transactionID))
	if er != nil {
		resp.Msg = translation.Translate(session.Lang, "transaction with this id does not exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	if transaction.State != state.TS_SELLER_ACCEPTED {
		resp.Errors = map[string]string{}
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "you are not authorized to perform this action in this stage")
		if transaction.State == state.TS_KEBD_REQUESTED {
			resp.Errors["TS_KEBD_REQUESTED"] = translation.Translate(session.Lang, "please complete the Kebd process first")
		} else if transaction.State == state.TS_ACCEPTED {
			resp.Errors["TS_ACCEPTED"] = translation.Translate(session.Lang, "the transaction is already accepted")
		} else if transaction.State == state.TS_CREATED {
			resp.Errors["TS_CREATED"] = translation.Translate(session.Lang, "The transaction has to be accepted by the seller first")
		} else if transaction.State == state.TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT {
			resp.Errors["TS_GUARANTEE_AMOUNT_AMEND_REQUEST_SENT"] = translation.Translate(session.Lang, "Please complete the guarantee payment amendment process before accepting")
		}
		c.JSON(resp.StatusCode, resp)
		return
	} else if transaction.RequesterID != session.ID {
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "unauthorized access")
		c.JSON(resp.StatusCode, resp)
		return
	}
	er = thandler.Service.BuyerAcceptTransaction(ctx, session.ID, uint64(transactionID))
	if er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(resp.StatusCode, resp)
		return
	}
	go thandler.InstantiatePaymentTransaction(ctx, session, transaction.ID)
	resp.Transaction = transaction
	resp.Msg = translation.Translate(session.Lang, "transaction accepted")
	resp.StatusCode = http.StatusOK
	c.JSON(resp.StatusCode, resp)

	// here i will instantiate the transaction payment handler instances
}

// InstantiatePaymentTransaction
func (thandler *TransactionHandler) InstantiatePaymentTransaction(ctx context.Context, session *model.Session, transactionID uint) uint8 {
	ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(time.Minute*2))
	transaction, er := thandler.Service.GetTransactionByID(ctx, uint64(transactionID))
	if er != nil || transaction == nil {
		println(er.Error())
		return 0
	}
	ctx = context.WithValue(ctx, "user_id", transaction.SellerID)
	seller, _, status, er := thandler.UserService.GetUserByEmailOrID(ctx)
	println(status)
	if seller == nil || er != nil || status < 0 {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		return 0
	}
	ctx = context.WithValue(ctx, "user_id", transaction.RequesterID)
	buyer, _, status, er := thandler.UserService.GetUserByEmailOrID(ctx)
	if buyer == nil || er != nil || status < 0 {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		return 0
	}
	transactionPayment := &model.TransactionPayment{
		SellerID: seller.ID,
		BuyerID:  buyer.ID,
		TransactionState: model.TransactionState{
			State:         state.TS_PAYMENT_INSTANTIATED,
			CreatedAt:     uint64(time.Now().Unix()),
			TransactionID: uint64(transaction.ID),
		},
		KebdAmount:      transaction.KebdAmount,
		GuaranteeAmount: transaction.GuaranteeAmount,
		KebdCompleted:   false,
		ServiceFee:      state.ServiceFee,
	}
	sellerInvoiceRequest := &model.HellocashInvoiceRequest{
		Amount:      int(transaction.GuaranteeAmount + transactionPayment.ServiceFee),
		From:        seller.Phone,
		Description: translation.Translate(session.Lang, "Please complete the kebd payment to proceed to the Contract"),
		Currency:    "ETB",
	}
	sellerInvoiceValidation, er := thandler.PaymentService.ValidateInvoice(ctx, sellerInvoiceRequest)
	if er != nil || sellerInvoiceValidation == nil {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		return 0
	}
	sellerInvoice, er := thandler.PaymentService.SendAnInvoice(ctx, sellerInvoiceRequest)
	if er != nil || sellerInvoice == nil {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		return 0
	}
	buyerInvoiceRequest := &model.HellocashInvoiceRequest{
		Amount:      int(transaction.KebdAmount + transactionPayment.ServiceFee),
		From:        buyer.Phone,
		Description: translation.Translate(session.Lang, "Please complete the kebd payment to proceed to the Contract"),
		Currency:    "ETB",
	}
	buyerInvoiceValidation, er := thandler.PaymentService.ValidateInvoice(ctx, buyerInvoiceRequest)
	if er != nil || buyerInvoiceValidation == nil {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		thandler.PaymentService.DeleteAnInvoiceByID(ctx, sellerInvoice.ID)
		return 0
	}
	buyerInvoice, er := thandler.PaymentService.SendAnInvoice(ctx, buyerInvoiceRequest)
	if er != nil || buyerInvoice == nil {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		thandler.PaymentService.DeleteAnInvoiceByID(ctx, sellerInvoice.ID)
		return 0
	}
	transactionPayment.BuyerInvoiceID = buyerInvoice.ID
	transactionPayment.SellerInvoiceID = sellerInvoice.ID
	//
	thandler.PaymentService.UpdatePaymentState(context.Background(), 10, uint64(transaction.ID))
	//
	stCode, er := thandler.PaymentService.CreateTransactionPayment(ctx, transactionPayment)
	if er != nil || stCode < 0 {
		if er != nil {
			println(er.Error())
		}
		thandler.PaymentService.UpdateTransactionPaymentStateByTransactionID(ctx, transaction.ID, uint(state.TS_ERROR))
		thandler.PaymentService.DeleteAnInvoiceByID(ctx, sellerInvoice.ID)
		thandler.PaymentService.DeleteAnInvoiceByID(ctx, buyerInvoice.ID)
		return 0
	}
	return 1
}

func (thandler *TransactionHandler) ReactivateTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	resp := &struct {
		Msg                 string                    `json:"msg"`
		StatusCode          int                       `json:"code"`
		TransactionID       uint64                    `json:"transaction_id"`
		Transaction         *model.Transaction        `json:"transaction,omitempty"`
		PaymentNotification *model.TransactionPayment `json:"payment_notification,omitempty"`
	}{}
	transactionID, er := strconv.Atoi(c.Param("id"))
	if er != nil || transactionID <= 0 {
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "no transaction found")
		c.JSON(resp.StatusCode, resp)
		return
	}
	transaction, er := thandler.Service.GetTransactionByID(ctx, uint64(transactionID))
	if er != nil || transaction == nil {
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "can't find transaction with this id")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if !(session.ID == transaction.RequesterID || session.ID == transaction.SellerID) {
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "you are not authorized to perform this operation")
		c.JSON(resp.StatusCode, resp)
		return
	}
	if !(transaction.State == state.TS_ERROR) {
		resp.StatusCode = http.StatusUnauthorized
		resp.Msg = translation.Translate(session.Lang, "can't reactivate a transaction which is not in the error stage")
	}
	stcode := thandler.InstantiatePaymentTransaction(ctx, session, transaction.ID)
	if stcode == 0 {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "failed to perform the refresh")
	} else {
		resp.StatusCode = http.StatusOK
		resp.Msg = translation.Translate(session.Lang, "payment Instance created succesfuly")
	}
	c.JSON(resp.StatusCode, resp)
}

// GetMyTransactionNotifications
func (thandler *TransactionHandler) GetMyTransactionNotifications(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	resp := &model.TransactionNotificationResponse{}
	results, _, er := thandler.Service.GetTransactionNotifications(ctx, session.Role, session.ID)
	if er != nil {
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "no transaction found")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, fmt.Sprintf(" found %d notifications ", len(results)))
	resp.Notifications = results
	c.JSON(resp.StatusCode, resp)
}
