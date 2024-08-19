package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/store"
	"github.com/samuael/shemach/shemach-backend/platforms/translation"
)

type IStoreHandler interface {
	CreateStore(c *gin.Context)
	GetMerchantStores(c *gin.Context)
	GetMerchantByID(c *gin.Context)
}

type StoreHandler struct {
	Service store.IStoreService
}

// NewStoreHandler  creating a new store handler instance.
func NewStoreHandler(service store.IStoreService) IStoreHandler {
	return &StoreHandler{
		Service: service,
	}
}

func (shandler *StoreHandler) CreateStore(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &struct {
		Name    string         `json:"name"`
		OwnerID uint           `json:"owner_id"`
		Address *model.Address `json:"address"`
	}{}

	resp := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		Store      *model.Store      `json:"store,omitempty"`
	}{
		Errors: map[string]string{},
	}

	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(input)
	if (err != nil) || (input.Address == nil || input.Address.Latitude == 0 || input.Address.Longitude == 0) || (input.OwnerID <= 0) || (len(strings.Trim(input.Name, " ")) <= 3) {
		resp.Msg = translation.TranslateIt("bad request data")
		resp.StatusCode = http.StatusBadRequest
		if input.Address == nil || input.Address.Latitude == 0 || input.Address.Longitude == 0 {
			resp.Errors["address"] = translation.TranslateIt("missing important coordinate information")
		}
		if input.OwnerID <= 0 {
			resp.Errors["owner_id"] = translation.TranslateIt("owner ID is missing")
		}
		if len(input.Name) <= 3 {
			resp.Errors["name"] = translation.TranslateIt("store name has to be greater that 3 character long")
		}
		c.JSON(resp.StatusCode, resp)
		return
	}

	store := &model.Store{
		StoreName: input.Name,
		OwnerID:   uint64(input.OwnerID),
		Address:   input.Address,
		CreatedBy: session.ID,
	}

	status, err := shandler.Service.CreateStore(ctx, store)
	if status <= 0 {
		if status == -1 {
			resp.StatusCode = http.StatusUnauthorized
			resp.Msg = translation.Translate(session.Lang, err.Error())
		} else if status == -2 {
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = translation.Translate(session.Lang, err.Error())
		} else if status == -3 {
			resp.StatusCode = http.StatusNotFound
			resp.Msg = translation.Translate(session.Lang, err.Error())
		} else if status == -4 {
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.Translate(session.Lang, err.Error())
		} else if status == -5 {
			println(err.Error())
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		} else if status == -6 {
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = translation.Translate(session.Lang, "unacceptable characters length input")
		}
		c.JSON(resp.StatusCode, resp)
		return
	}

	// Success end
	resp.Store = store
	resp.StatusCode = http.StatusCreated
	resp.Msg = translation.Translate(session.Lang, "created succesfuly")
	c.JSON(resp.StatusCode, resp)
}

func (shandler *StoreHandler) GetMerchantStores(c *gin.Context) {
	ctx := c.Request.Context()
	merchantid, er := strconv.Atoi(c.Query("id"))
	session := ctx.Value("session").(*model.Session)
	res := &struct {
		StatusCode int            `json:"status_code"`
		Msg        string         `json:"msg"`
		Stores     []*model.Store `json:"stores"`
	}{}
	if er != nil || merchantid <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "missing merchant id \"id\" of type integer >0 ")
		c.JSON(res.StatusCode, res)
		return
	}
	stores, er := shandler.Service.GetMerchantStores(ctx, uint64(merchantid))
	if er != nil || len(stores) == 0 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "no store instance found")
	} else {
		res.Stores = stores
		res.Msg = translation.Translate(session.Lang, "found")
		res.StatusCode = http.StatusOK
	}
	c.JSON(res.StatusCode, res)
}

func (shandler *StoreHandler) GetMerchantByID(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	res := &struct {
		StatusCode int          `json:"status_code"`
		Msg        string       `json:"msg"`
		Store      *model.Store `json:"store"`
	}{}
	storeid, er := strconv.Atoi(c.Query("id"))
	if er != nil || storeid <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "missing s itore id \"id\" of type integer >0 ")
		c.JSON(res.StatusCode, res)
		return
	}
	store, er := shandler.Service.GetStoreByID(ctx, uint64(storeid))
	if er != nil || store == nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "store instance not found")
	} else {
		res.Store = store
		res.Msg = translation.Translate(session.Lang, "found")
		res.StatusCode = http.StatusOK
	}
	c.JSON(res.StatusCode, res)
}
