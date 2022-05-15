package rest

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/agent"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
	"github.com/samuael/agri-net/agri-net-backend/pkg/store"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ICropHandler interface {
	CreateProduct(c *gin.Context)
	UploadProductImages(c *gin.Context)
}
type CropHandler struct {
	Service         crop.ICropService
	ProductService  product.IProductService
	StoreService    store.IStoreService
	MerchantService merchant.IMerchantService
	AgentService    agent.IAgentService
}

func NewCropHandler(
	service crop.ICropService,
	productservice product.IProductService,
	storeService store.IStoreService,
	merchantService merchant.IMerchantService,
	agentService agent.IAgentService,
) ICropHandler {
	return &CropHandler{
		Service:         service,
		ProductService:  productservice,
		StoreService:    storeService,
		MerchantService: merchantService,
		AgentService:    agentService,
	}
}

func (chandler *CropHandler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	if !(session.Role == state.MERCHANT || session.Role == state.AGENT) {
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}
	input := &struct {
		TypeID            uint    `json:"type_id"`
		SellingPrice      float64 `json:"selling_price"`
		StoreID           uint64  `json:"store_id"`
		Description       string  `json:"description"`
		RemainingQuantity uint64  `json:"quantity"`
		Negotiable        bool    `json:"negotiable_price"`
	}{}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Crop       *model.Crop       `json:"crop"`
		Errors     map[string]string `json:"errors,omitempty"`
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad request input")
		c.JSON(res.StatusCode, res)
		return
	}
	failed := false
	ack, desc := chandler.ProductService.GetProductInfoByID(int(input.TypeID))
	if ack == "" || desc == "" {
		res.Errors["type"] = translation.Translate(session.Lang, "product type with this id does not exist")
		failed = true
	}
	if input.RemainingQuantity <= 0 {
		res.Errors["remaining_quantity"] = translation.Translate(session.Lang, "unacceptable remaining quantity")
		failed = true
	}
	if input.SellingPrice <= 0 {
		res.Errors["selling_price"] = translation.Translate(session.Lang, "invalid selling price")
		failed = true
	}
	var address *model.Address
	isstoreOwned := true
	if session.Role == state.MERCHANT && input.StoreID <= 0 {
		res.Errors["store_id"] = translation.Translate(session.Lang, "the store id is missing ")
		failed = true
	} else if session.Role == state.MERCHANT {
		store, er := chandler.StoreService.GetStoreByID(ctx, input.StoreID)
		if store.OwnerID != session.ID || er != nil {
			res.StatusCode = http.StatusBadRequest
			res.Errors["store"] = translation.Translate(session.Lang, "you are not authorized to access this store")
			failed = true
		}
		address = store.Address
	} else if session.Role == state.AGENT {
		isstoreOwned = false
		ad, er := chandler.AgentService.GetAgentsAddress(ctx, int(session.ID))
		if er != nil {
			res.StatusCode = http.StatusNotImplemented
			res.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
			c.JSON(res.StatusCode, res)
			return
		} else {
			address = ad
		}
	}
	if address == nil {
		res.StatusCode = http.StatusNotImplemented
		res.Msg = translation.Translate(session.Lang, "internal problem, please try again later!")
		c.JSON(res.StatusCode, res)
		return
	}
	if failed {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad request payload")
		c.JSON(res.StatusCode, res)
		return
	}
	crop := &model.Crop{
		TypeID:            input.TypeID,
		Description:       input.Description,
		Address:           address,
		RemainingQuantity: input.RemainingQuantity,
		Negotiable:        input.Negotiable,
		StoreID: func() uint64 {
			if isstoreOwned {
				return input.StoreID
			}
			return 0
		}(),
		StoreOwned:   isstoreOwned,
		SellingPrice: input.SellingPrice,
		AddressRef:   uint64(address.ID),
		AgentID: func() uint64 {
			if session.Role == state.AGENT {
				return session.ID
			}
			return 0
		}(),
	}
	status, err := chandler.Service.CreateCrop(ctx, crop)
	if err != nil || status < 0 {
		if status == -1 {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.Translate(session.Lang, "missing address information")
		} else if status == -2 {
			res.StatusCode = http.StatusUnauthorized
			res.Msg = translation.Translate(session.Lang, "you are not authorized")
		} else {
			res.Msg = translation.Translate(session.Lang, "internal server error")
			res.StatusCode = http.StatusInternalServerError
		}
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusCreated
	res.Msg = translation.Translate(session.Lang, "created")
	res.Crop = crop
	res.Errors = nil
	c.JSON(res.StatusCode, res)
}

type MultipartData struct {
	File   multipart.File
	Header *multipart.FileHeader
	Error  error
}

// UploadProductImages ...
func (chandler CropHandler) UploadProductImages(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	err := c.Request.ParseMultipartForm(999999999999999999)
	res := &struct {
		Msg               string   `json:"msg"`
		StatusCode        int      `json:"status_code"`
		BluredImageRoutes []string `json:"blured_image_routes"`
		ImageRoutes       []string `json:"image_routes"`
		Error             string   `json:"error,omitempty"`
	}{}
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad multipart for file file Size Exceeds the limit")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var subArticleImageFiles = map[int]*os.File{}
	subArticleImages := map[int]*MultipartData{}
	for a := 0; a < 5; a++ {
		sf := &MultipartData{}
		sf.File, sf.Header, sf.Error = c.Request.FormFile("image" + strconv.Itoa(a))
		if sf.File == nil || sf.Header == nil || sf.Error != nil {
			continue
		}
		defer sf.File.Close()
		filename :=
			state.POST_IMAGES_RELATIVE_PATH +
				helper.GenerateRandomString(6, helper.CHARACTERS) +
				"." + helper.GetExtension(sf.Header.Filename)
		file, ee := os.Create(os.Getenv("ASSETS_DIRECTORY") + filename)
		if ee != nil {
			println(ee.Error())
			res.StatusCode = http.StatusInternalServerError
			res.Error = " internal server error "
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		subArticleImageFiles[a] = file
		defer subArticleImageFiles[a].Close()
		subArticleImages[a] = sf
	}
	println("length : ", len(subArticleImages))
}
