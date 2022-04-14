package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
)

type IProductHandler interface {
	CreateProductInstance(c *gin.Context)
}

type ProductHandler struct {
	Service product.IProductService
}

func NewProductService(service product.IProductService) IProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (pser *ProductHandler) CreateProductInstance(c *gin.Context) {
	ctx := c.Request.Context()
	// CREATINT THE PRODUCT INSTANCE.
	println(ctx)
}
