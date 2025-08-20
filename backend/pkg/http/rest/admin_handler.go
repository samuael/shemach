package rest

import (
	"github.com/gin-gonic/gin"
)

type IAdminHandler interface {
	RegisterAdmin(c *gin.Context)
	ListAdmins(c *gin.Context)
	DeleteAdminByID(c *gin.Context)
	GetAdminByID(c *gin.Context)
}
