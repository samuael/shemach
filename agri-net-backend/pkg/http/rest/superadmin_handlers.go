package rest

import "github.com/gin-gonic/gin"

type ISuperadminHandler interface{}
type SuperadminHandler struct {
	//
}

func NewSuperadminHandler() ISuperadminHandler {
	return &SuperadminHandler{}
}

func (suhhandler *SuperadminHandler) SuperadminLogin(c *gin.Context) {
	ctx := c.Request.Context()
	println(ctx)
}
