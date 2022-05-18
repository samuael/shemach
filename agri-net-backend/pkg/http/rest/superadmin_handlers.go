package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/superadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ISuperadminHandler interface {
	SuperadminRegistration(c *gin.Context)
	GetSuperadminByID(c *gin.Context)
	GetSystemSuperadmin(c *gin.Context)
}
type SuperadminHandler struct {
	Service       superadmin.ISuperadminService
	UserService   user.IUserService
	Authenticator auth.Authenticator
}

func NewSuperadminHandler(
	service superadmin.ISuperadminService,
	authenticator auth.Authenticator,
	uservice user.IUserService,
) ISuperadminHandler {
	return &SuperadminHandler{
		Service:       service,
		Authenticator: authenticator,
		UserService:   uservice,
	}
}

func (suhhandler *SuperadminHandler) SuperadminRegistration(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
	}{}
	res := &struct {
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
	}{
		Errors: map[string]string{},
	}

	jsonDecode := json.NewDecoder(c.Request.Body)
	er := jsonDecode.Decode(input)
	if er != nil ||
		input.Firstname == "" ||
		input.Lastname == "" ||
		input.Email == "" ||
		input.Phone == "" {
		res.Msg = translation.TranslateIt("bad input")
		res.StatusCode = http.StatusBadRequest
		if (input.Firstname == "") || (len(strings.Split(input.Firstname, " ")) > 1) {
			res.Errors["firstname"] = "bad first name value"
		}
		if (input.Lastname == "") || (len(strings.Split(input.Lastname, " ")) > 1) {
			res.Errors["lastname"] = "bad last name value"
		}
		if (input.Email == "") || (len(strings.Split(input.Email, " ")) > 1) {
			res.Errors["email"] = "bad email address"
		}
		if (input.Phone == "") || (len(strings.Split(input.Phone, " ")) > 1) {
			res.Errors["phone"] = "bad Phone number"
		}
		c.JSON(res.StatusCode, res)
		return
	}
	println(session.ID)
	// check the existance of the user
}

func (suhandler *SuperadminHandler) GetSuperadminByID(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	res := &struct {
		StatusCode int               `json:""status_code`
		Msg        string            `json:"msg"`
		Superadmin *model.Superadmin `json:"superadmin,omitempty"`
	}{}
	id, er := strconv.Atoi(c.Query("id"))
	if er != nil || id <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "bad query string")
		c.JSON(res.StatusCode, res)
		return
	}
	superadmin, er := suhandler.Service.GetSuperadminByID(ctx, id)
	if er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "superadmin with this id doesn't exist")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Superadmin = superadmin
	res.StatusCode = http.StatusOK
	res.Msg = translation.Translate(session.Lang, "superadmin instance found!")
	c.JSON(res.StatusCode, res)
}

// GetSystemSuperadmins
func (suhandler *SuperadminHandler) GetSystemSuperadmin(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		Superadmins []*model.Superadmin `json:"superadmins"`
		StatusCode  int                 `json:"status_code"`
	}{}
	superadmins, er := suhandler.Service.GetAllSuperadmins(ctx)
	if er != nil {
		res.StatusCode = http.StatusNotFound
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Superadmins = superadmins
	c.JSON(res.StatusCode, res)
}
