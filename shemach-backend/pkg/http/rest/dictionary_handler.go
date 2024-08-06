package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/dictionary"
	"github.com/samuael/shemach/shemach-backend/platforms/translation"
)

type IDictionaryHandler interface {
	CreateDictionary(c *gin.Context)
	UpdateDictionary(c *gin.Context)
	DeleteDictionary(c *gin.Context)
	Translate(c *gin.Context)
	GetRecentDictionaries(c *gin.Context)
}

type DictionaryHandler struct {
	Service dictionary.IDictionaryService
}

func NewDictionaryHandler(service dictionary.IDictionaryService) IDictionaryHandler {
	return &DictionaryHandler{
		Service: service,
	}
}

// CreateDictionary ...
func (dhandler *DictionaryHandler) CreateDictionary(c *gin.Context) {
	ctx := c.Request.Context()
	input := &model.Dictionary{}
	res := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Dictionary *model.Dictionary `json:"dictionary"`
	}{}
	jdec := json.NewDecoder(c.Request.Body)
	er := jdec.Decode(input)
	if er != nil || input.Lang == "" || input.Text == "" || input.Translation == "" {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("missing important input variable")
		c.JSON(res.StatusCode, res)
		return
	}
	er = dhandler.Service.NewDictionary(ctx, input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt(er.Error())
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = "Created succesfuly"
	res.Dictionary = input
	c.JSON(res.StatusCode, res)
}

func (dhandler *DictionaryHandler) UpdateDictionary(c *gin.Context) {
	ctx := c.Request.Context()
	input := &model.Dictionary{}
	res := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Dictionary *model.Dictionary `json:"dictionary"`
	}{}
	jdec := json.NewDecoder(c.Request.Body)
	er := jdec.Decode(input)
	if er != nil || input.Lang == "" || input.Text == "" || input.Translation == "" || input.ID <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("missing important input variable")
		c.JSON(res.StatusCode, res)
		return
	}
	er = dhandler.Service.UpdateTranslation(ctx, input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt(er.Error())
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("update")
	res.Dictionary = input
	c.JSON(res.StatusCode, res)
}

// DeleteDictionary ...
func (dhandler *DictionaryHandler) DeleteDictionary(c *gin.Context) {
	ctx := c.Request.Context()
	input := &model.Dictionary{}

	id, era := strconv.Atoi(c.Query("id"))
	res := &struct {
		Msg        string `json:"msg"`
		StatusCode int    `json:"status_code"`
	}{}
	if era != nil {
		jdec := json.NewDecoder(c.Request.Body)
		er := jdec.Decode(input)
		if er != nil || input.ID <= 0 {
			res.StatusCode = http.StatusBadRequest
			res.Msg = translation.TranslateIt("missing important input variable")
			c.JSON(res.StatusCode, res)
			return
		}
	} else {
		input.ID = uint64(id)
	}
	count, er := dhandler.Service.DeleteTranslation(ctx, input)
	if er != nil || count > 0 {
		if er != nil {
			println(er.Error())
		}
		res.StatusCode = http.StatusOK
		res.Msg = translation.TranslateIt("record deleted succesfuly")
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("deleted succesfuly")
	c.JSON(res.StatusCode, res)
}

// Translate
func (dhandler *DictionaryHandler) Translate(c *gin.Context) {
	ctx := c.Request.Context()
	input := &model.Dictionary{}
	res := &struct {
		Msg        string              `json:"msg,omitempty"`
		StatusCode int                 `json:"status_code"`
		Dictionary []*model.Dictionary `json:"dictionary"`
	}{}
	jdec := json.NewDecoder(c.Request.Body)
	er := jdec.Decode(input)
	if er != nil || input.Lang == "" || input.Text == "" {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("missing important input variable")
		c.JSON(res.StatusCode, res)
		return
	}
	er = dhandler.Service.Translate(ctx, input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("can't find a translation for the input text")
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Dictionary = []*model.Dictionary{input}
	c.JSON(res.StatusCode, res)
}

// GetRecentDictionaries ...
func (dhandler *DictionaryHandler) GetRecentDictionaries(c *gin.Context) {
	ctx := c.Request.Context()
	offset, er := strconv.Atoi(c.Query("offset"))
	if er != nil {
		offset = 0
	}
	limit, er := strconv.Atoi(c.Query("limit"))
	if er != nil {
		limit = 0
	}
	res := &struct {
		Msg          string              `json:"msg"`
		StatusCode   int                 `json:"status_code"`
		Dictionaries []*model.Dictionary `json:"dictionaries,omitempty"`
	}{}
	dictionaries, er := dhandler.Service.GetDictionaries(ctx, uint(offset), uint(limit))
	if er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("no dictionaries found ")
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Dictionaries = dictionaries
	c.JSON(res.StatusCode, res)
}
