package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IInfoadminHandler interface {
	Registerinfoadmin(c *gin.Context)
	ListInfoadmins(c *gin.Context)
	DeleteInfoadminByID(c *gin.Context)
	GetInfoadminByID(c *gin.Context)
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

func (ihandler InfoadminHandler) Registerinfoadmin(c *gin.Context) {
	input := &struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
	}{}
	resp := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Errors     map[string]string `json:"errors"`
		Infoadmin  *model.Infoadmin  `json:"info_admin"`
	}{
		Errors: map[string]string{},
	}
	if er := c.BindJSON(input); er == nil {
		fail := false
		if !form.MatchesPattern(input.Email, form.EmailRX) {
			resp.Errors["email"] = "invalid email address"
			resp.StatusCode = http.StatusBadRequest
			fail = true
		}
		if len(input.Firstname) <= 3 || len(input.Lastname) <= 3 {
			if len(input.Firstname) <= 3 {
				resp.Errors["lastname"] = "unacceptable firstname value"
			}
			if len(input.Lastname) <= 3 {
				resp.Errors[""] = "unacceptable last name value"
			}
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = "request value errors"
			fail = true
		}
		if !fail {
			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, "infoadmin_email", input.Email)
			if admin, err := ihandler.Service.GetInfoadminByEmail(ctx); admin != nil && err == nil {
				println(admin.ID, admin.Email, admin.Phone, admin.Password)
				resp.Msg = "account with this email already exist."
				c.JSON(http.StatusUnauthorized, resp)
				return
			}
			if er != nil {
				resp.Msg = " Internal Server error "
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			random := "admin" //helper.GenerateRandomString(5, helper.CHARACTERS)
			hashed, _ := helper.HashPassword(random)

			admin := &model.Infoadmin{}
			admin.Firstname = input.Firstname
			admin.Lastname = input.Lastname
			admin.Email = input.Email //
			admin.CreatedAt = uint64(time.Now().Unix())
			admin.Password = hashed
			// Send Email for the password if this doesn't work raise internal server error
			// if success := mail.SendPasswordEmailSMTP([]string{admin.Email}, random, true, admin.Firstname+" "+admin.Lastname, c.Request.Host); success {
			ctx = c.Request.Context()
			ctx = context.WithValue(ctx, "info_admin", admin)
			if admin, er = ihandler.Service.CreateInfoadmin(ctx); admin != nil && er == nil {
				resp.Msg = " Info admin  created succesfully!"
				resp.Infoadmin = admin
				c.JSON(http.StatusOK, resp)
				return
			} else {
				if admin != nil && er != nil {
					resp.Msg = er.Error()
				} else {
					resp.Msg = "Internal server error!"
				}
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			// } else {
			// 	resp.Msg = "Internal server error!"
			// 	c.JSON(http.StatusInternalServerError, resp)
			// 	return
			// }
		}
	}
	c.JSON(http.StatusBadRequest, resp)
}

func (ihandler *InfoadminHandler) ListInfoadmins(c *gin.Context) {
	ctx := c.Request.Context()
	error_res := &struct {
		Msg        string `json:"msg"`
		StatusCode int    `json:"status_code"`
	}{}
	var infoadmins []*model.Infoadmin
	infoadmins, er := ihandler.Service.GetInfoadmins(ctx)
	if er != nil {
		error_res.Msg = translation.TranslateIt("record  not found ")
		error_res.StatusCode = http.StatusNotFound
		c.JSON(error_res.StatusCode, error_res)
		return
	}
	c.JSON(http.StatusOK, infoadmins)
}

func (ihandler *InfoadminHandler) DeleteInfoadminByID(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Errors     map[string]string `json:"errors,omitempty"`
	}{
		Errors: map[string]string{},
	}
	id, er := strconv.Atoi(c.Query("id"))
	if er != nil || id < 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("missing important parameter")
		res.Errors["id"] = translation.TranslateIt("parameter infoadmin \"id\" of type {{integer}} >0 must be provided")
		c.JSON(res.StatusCode, res)
		return
	}
	ctx = context.WithValue(ctx, "infoadmin_id", uint64(id))
	status, er := ihandler.Service.DeleteInfoadminByID(ctx)
	if er != nil || status != 0 {
		if status == -1 {
			res.StatusCode = http.StatusNotFound
			res.Msg = translation.TranslateIt("Info admin with the specified ID does not exist")
		} else {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = translation.TranslateIt("internal problem, please try again later")
		}
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("info admin deleted succesfuly")
	c.JSON(res.StatusCode, res)
}

// GetInfoadminByID(c *gin.Context)
func (ihandler *InfoadminHandler) GetInfoadminByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, er := strconv.Atoi(c.Query("id"))

	res := &struct {
		StatusCode int              `json:"status_code"`
		Msg        string           `json:"msg"`
		Admin      *model.Infoadmin `json:"infoadmin"`
	}{}

	if er != nil || id <= 0 {
		res.StatusCode = http.StatusOK
		res.Msg = translation.TranslateIt("missing admin id \"id\" , positive integer")
		c.JSON(res.StatusCode, res)
		return
	}
	admin, er := ihandler.Service.GetInfoadminByID(ctx, uint64(id))
	if admin == nil || er != nil {
		res.StatusCode = http.StatusNotFound
		res.Admin = admin
		res.Msg = translation.TranslateIt("admin no found ")
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Admin = admin
	res.Msg = translation.TranslateIt("found")
	c.JSON(res.StatusCode, res)
}
