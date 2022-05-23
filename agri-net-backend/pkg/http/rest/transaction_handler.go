package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
	"github.com/samuael/agri-net/agri-net-backend/pkg/store"
	"github.com/samuael/agri-net/agri-net-backend/pkg/transaction"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
}

// TransactionHandler transaction handler instance
type TransactionHandler struct {
	Service         transaction.ITransactionService
	ProductService  crop.ICropService
	MerchantService merchant.IMerchantService
	StoreService    store.IStoreService
}

func NewTransactionHandler(
	service transaction.ITransactionService,
	productService crop.ICropService,
	merchantService merchant.IMerchantService,
	storeService store.IStoreService,
) ITransactionHandler {
	return &TransactionHandler{
		Service:         service,
		ProductService:  productService,
		MerchantService: merchantService,
		StoreService:    storeService,
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
		resp.Msg = translation.Translate(session.Lang, "product instance with this id doesn't exist")
		resp.StatusCode = http.StatusNotFound
		c.JSON(resp.StatusCode, resp)
		return
	}
	if !product.Closed {
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
	// transaction  , er :=
	println(transaction)

}
