package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/mail"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IAdminHandler interface {
	RegisterAdmin(c *gin.Context)
	ListAdmins(c *gin.Context)
	DeleteAdminByID(c *gin.Context)
	GetAdminByID(c *gin.Context)
}

// InfoAdminHandler
type AdminHandler struct {
	Service admin.IAdminService
}

func NewAdminHandler(ser admin.IAdminService) IAdminHandler {
	return &AdminHandler{
		Service: ser,
	}
}

func (ahandler *AdminHandler) RegisterAdmin(c *gin.Context) {
	input := &struct {
		Firstname string         `json:"firstname"`
		Lastname  string         `json:"lastname"`
		Email     string         `json:"email"`
		Address   *model.Address `json:"address"`
	}{}
	resp := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Errors     map[string]string `json:"errors"`
		Admin      *model.Admin      `json:"admin"`
	}{
		Errors: map[string]string{},
	}
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)

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
			if admin, err := ahandler.Service.GetAdminByEmail(ctx, input.Email); admin != nil && err == nil {
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

			admin := &model.Admin{}
			admin.Firstname = input.Firstname
			admin.Lastname = input.Lastname
			admin.Email = input.Email //
			admin.CreatedAt = uint64(time.Now().Unix())
			admin.CreatedBy = int(session.ID)
			admin.Password = hashed
			admin.FieldAddress = input.Address
			admin.Lang = session.Lang
			if admin.FieldAddress == nil {
				admin.FieldAddress = &model.Address{}
			}
			// Send Email for the password if this doesn't work raise internal server error
			if success := mail.SendPasswordEmailSMTP([]string{admin.Email}, random, true, admin.Firstname+" "+admin.Lastname, c.Request.Host); success {
				ctx = c.Request.Context()
				adminID, addressID, er := ahandler.Service.CreateAdmin(ctx, admin)
				if adminID > 0 && er == nil {
					resp.Msg = "admin registered succesfully!"
					resp.StatusCode = http.StatusOK
					admin.ID = uint64(adminID)
					admin.FieldAddress.ID = uint(addressID)
					resp.Admin = admin
					c.JSON(resp.StatusCode, resp)
					return
				} else {
					if adminID == -1 {
						resp.StatusCode = http.StatusUnauthorized
						resp.Msg = translation.Translate(session.Lang, "unauthorized access")
					} else if adminID == -2 {
						resp.StatusCode = http.StatusInternalServerError
						resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later")
					} else {
						resp.StatusCode = http.StatusInternalServerError
						resp.Msg = translation.Translate(session.Lang, "Internal server error, please try again later")
					}
					c.JSON(resp.StatusCode, resp)
					return
				}
			} else {
				resp.Msg = "Internal server error!"
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
	}
	c.JSON(http.StatusBadRequest, resp)
}

func (ahandler *AdminHandler) ListAdmins(c *gin.Context) {
	ctx := c.Request.Context()
	error_res := &struct {
		Msg        string `json:"msg"`
		StatusCode int    `json:"status_code"`
	}{}
	offset, er := strconv.Atoi(c.Query("offset"))
	if er != nil {
		offset = 0
	}
	limit, er := strconv.Atoi(c.Query("limit"))
	if er != nil {
		limit = offset + 20
	}
	var infoadmins []*model.Admin
	infoadmins, er = ahandler.Service.GetAdmins(ctx, offset, limit)
	if er != nil {
		error_res.Msg = translation.TranslateIt("record  not found ")
		error_res.StatusCode = http.StatusNotFound
		c.JSON(error_res.StatusCode, error_res)
		return
	}
	c.JSON(http.StatusOK, infoadmins)
}

func (ahandler *AdminHandler) DeleteAdminByID(c *gin.Context) {
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
	// ctx = context.WithValue(ctx, "admin_id", uint64(id))
	status, er := ahandler.Service.DeleteAdminByID(ctx, uint64(id))
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
func (ahandler *AdminHandler) GetAdminByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, er := strconv.Atoi(c.Query("id"))

	res := &struct {
		StatusCode int          `json:"status_code"`
		Msg        string       `json:"msg"`
		Admin      *model.Admin `json:"admin"`
	}{}

	if er != nil || id <= 0 {
		res.StatusCode = http.StatusOK
		res.Msg = translation.TranslateIt("missing admin id \"id\" , positive integer")
		c.JSON(res.StatusCode, res)
		return
	}
	admin, er := ahandler.Service.GetAdminByID(ctx, uint64(id))
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
