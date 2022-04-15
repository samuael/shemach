package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
)

type IInfoadminHandler interface {
	CreateInfoadmin(c *gin.Context)
}

// InfoAdminHandler
type InfoadminHandler struct {
	Service infoadmin.IInfoadminService
}

func NewInfoAdminHandler(ser infoadmin.IInfoadminService) IInfoadminHandler {
	return &InfoadminHandler{
		Service: ser,
	}
}

func (ihandler InfoadminHandler) CreateInfoadmin(c *gin.Context) {
	ctx := c.Request.Context()
	// ----
	println(ctx)
}
