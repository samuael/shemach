package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/resource"
)

type IResourceHandler interface {
	GetProductImage(c *gin.Context)
	GetBlurredImage(c *gin.Context)
}

type ResourceHandler struct {
	Service resource.IResourceService
}

func NewResourceHandler(service resource.IResourceService) IResourceHandler {
	return &ResourceHandler{
		Service: service,
	}
}

// GetProductImage
func (rhandler *ResourceHandler) GetProductImage(c *gin.Context) {
	id, er := strconv.Atoi(c.Query("id"))
	if id <= 0 || er != nil {
		// c.JSON(http.StatusNotFound)
	}
}

// GetBlurredImage
func (rhandler *ResourceHandler) GetBlurredImage(c *gin.Context) {
	//
}
