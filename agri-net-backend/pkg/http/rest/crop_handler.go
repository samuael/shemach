package rest

import (
	"github.com/samuael/agri-net/agri-net-backend/pkg/crop"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
	"github.com/samuael/agri-net/agri-net-backend/pkg/store"
)

type ICropHandler interface {
}
type CropHandler struct {
	Service        crop.ICropService
	ProductService product.IProductService
	StoreService   store.IStoreService
}

func NewCropHandler(
	service crop.ICropService,
	productservice product.IProductService,
	StoreService store.IStoreService,
) ICropHandler {
	return &CropHandler{
		Service:        service,
		ProductService: productservice,
		StoreService:   StoreService,
	}
}

// func (chandler *CropHandler) CreateProduct(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	input := &struct {
// 		TypeID       uint           `json:"type_id"`
// 		SellingPrice float64        `json:"selling_price"`
// 		Address      *model.Address `json:"address"`
// 		StoreID      uint64         `json:"store_id"`
// 	}{}

// 	res := &struct {
// 		StatusCode int               `json:"status_code"`
// 		Msg        string            `json:"msg"`
// 		Crop       *model.Crop       `json:"crop"`
// 		Errors     map[string]string `json:"errors"`
// 	}{}

// 	session := ctx.Value("session").(*model.Session)
// 	jdecode := json.NewDecoder(c.Request.Body)
// 	er := jdecode.Decode(input)
// 	if er != nil {
// 		res.StatusCode = http.StatusBadRequest
// 		res.Msg = translation.Translate(session.Lang, "bad request input")
// 		c.JSON(res.StatusCode, res)
// 		return
// 	}
// 	failed := false
// 	ack, desc := chandler.ProductService.GetProductInfoByID(int(input.TypeID))
// 	if ack == "" || desc == "" {
// 		res.Errors["type"] = translation.Translate(session.Lang, "product type with this id does not exist")
// 		failed = true

// 		res.Msg = translation.Translate(session.Lang, "bad requuest")
// 		res.StatusCode = http.StatusBadRequest
// 		c.JSON(res.StatusCode, res)
// 		return
// 	}
// 	if input.SellingPrice <= 0 {
// 		res.Errors["selling_price"] = translation.Translate(session.Lang, "invalid selling price")
// 		failed = true
// 	}
// 	if input.Address != nil && input.Address.Latitude == 0 || input.Address.Longitude == 0 {
// 		res.Errors["address"] = translation.Translate(session.Lang, "invalid product address information")
// 	}
// 	if session.Role == state.MERCHANT && input.StoreID <= 0 {
// 		res.Errors["store_id"] = translation.Translate(session.Lang, "the store id is missing ")

// 		store, er := chandler.StoreService.GetStoreByID(ctx, input.StoreID)
// 		if store.OwnerID != session.ID {
// 			chandler.
// 		}
// 	}

// }
