package rest

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/resource"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
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
	id, er := strconv.Atoi(c.Param("id"))
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	if id <= 0 || er != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	diabledImageServer, _ := strconv.ParseBool(os.Getenv("DISALED_IMAGE_SERVER"))
	imgres, er := rhandler.Service.GetImageByID(ctx, uint64(id))
	if er != nil {
		println(er.Error())
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	var role int8
	var f *os.File
	var err error
	if !diabledImageServer {
		if imgres.Authorized {
			role = helper.RoleIntFromStringRole(session.Role)
			if !(role == state.SUPERADMIN_ROLE_INT && (imgres.Authorizations&state.SUPERADMIN_ROLE_INT) == state.SUPERADMIN_ROLE_INT) ||
				!(role == state.ADMIN_ROLE_INT && (imgres.Authorizations&state.ADMIN_ROLE_INT) == state.ADMIN_ROLE_INT) ||
				!(role == state.INFOADMIN_ROLE_INT && (imgres.Authorizations&state.INFOADMIN_ROLE_INT) == state.INFOADMIN_ROLE_INT) ||
				!(role == state.MERCHANT_ROLE_INT && (imgres.Authorizations&state.MERCHANT_ROLE_INT) == state.MERCHANT_ROLE_INT) ||
				!(role == state.AGENT_ROLE_INT && (imgres.Authorizations&state.AGENT_ROLE_INT) == state.AGENT_ROLE_INT) {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		if er != nil || imgres == nil {
			println(er.Error())
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		f, err = os.Open(os.Getenv("ASSETS_DIRECTORY") + imgres.Resource)
	} else {
		f, err = os.Open("../../templates/dummy_img/eudhWR.jpg")
	}
	if err != nil {
		println(err.Error())
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		println(err.Error())
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=YOURNAME")
	http.ServeContent(c.Writer, c.Request, helper.GetExtension(helper.GenerateRandomString(5, helper.CHARACTERS)), fi.ModTime(), f)
}

// GetBlurredImage
func (rhandler *ResourceHandler) GetBlurredImage(c *gin.Context) {
	println("Get Blurred Image get called.")
	id, er := strconv.Atoi(c.Param("id"))
	ctx := c.Request.Context()
	if id <= 0 || er != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	imgres, er := rhandler.Service.GetImageByID(ctx, uint64(id))
	if er != nil || imgres == nil {
		if er != nil {
			println(er.Error())
		}
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	diabledImageServer, _ := strconv.ParseBool(os.Getenv("DISALED_IMAGE_SERVER"))

	var f *os.File
	var err error
	if !diabledImageServer {
		println("The Image Reference " + imgres.BlurredRe)
		f, err = os.Open(os.Getenv("ASSETS_DIRECTORY") + imgres.BlurredRe)
	} else {
		f, err = os.Open("../../templates/dummy_img/C2Dtn6_shJ.jpg")
	}
	if err != nil {
		println(err.Error())
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		println(err.Error())
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=YOURNAME")
	http.ServeContent(c.Writer, c.Request, helper.GetExtension(helper.GenerateRandomString(5, helper.CHARACTERS)), fi.ModTime(), f)
}
