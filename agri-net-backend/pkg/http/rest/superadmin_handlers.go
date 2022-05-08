package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/superadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ISuperadminHandler interface {
	AdminsLogin(c *gin.Context)
	SuperadminRegistration(c *gin.Context)
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
func (suhandler *SuperadminHandler) AdminsLogin(c *gin.Context) {
	ctx := c.Request.Context()
	input := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	jsonDecoder := json.NewDecoder(c.Request.Body)
	res := &struct {
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
		User       *model.User       `json:"user,omitempty"`
		Role       string            `json:"role,omitempty"`
		Token      string            `json:"token,omitempty"`
	}{
		Errors: map[string]string{},
	}
	ers := jsonDecoder.Decode(input)
	if ers != nil || !(form.MatchesPattern(input.Email, form.EmailRX)) || len(input.Password) < 4 {
		res.Msg = translation.TranslateIt("bad request body")
		if !(form.MatchesPattern(input.Email, form.EmailRX)) {
			res.Errors["email"] = translation.TranslateIt("invalid phone value")
		}
		if len(input.Password) < 4 {
			res.Errors["password"] = translation.TranslateIt("unacceptable password characters length")
		}
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	// check the existance of the user using his  email only.
	ctx = context.WithValue(ctx, "user_email", input.Email)
	user, role, status, er := suhandler.UserService.GetUserByEmailOrID(ctx)
	var failed = false
	if status == state.STATUS_DBQUERY_ERROR || er != nil {
		failed = true
		res.StatusCode = http.StatusInternalServerError
		res.Msg = translation.TranslateIt("internal problem, please try again later!")
	} else if status == state.STATUS_RECORD_NOT_FOUND {
		failed = true
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("invalid email or password!")
	}
	if role == 0 {
		failed = true
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("user with this account doesn't exist")
	}

	if failed {
		c.JSON(res.StatusCode, res)
		return
	}
	if !(helper.CompareHash(user.Password, input.Password)) {
		res.Msg = translation.TranslateIt("invalid email or password")
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	session := &model.Session{
		ID:    user.ID,
		Email: user.Email,
		Lang:  user.Lang,
	}
	// var duser interface{}
	if role == 1 {
		session.Role = state.SUPERADMIN
		res.Role = state.SUPERADMIN
	} else if role == 2 {
		session.Role = state.INFO_ADMIN
		res.Role = state.INFO_ADMIN
	} else if role == 3 {
		session.Role = state.ADMIN
		res.Role = state.ADMIN
	} else if role == 4 {
		session.Role = state.MERCHANT
		res.Role = state.MERCHANT
	} else if role == 5 {
		session.Role = state.AGENT
		res.Role = state.AGENT
	}
	suhandler.Authenticator.SaveSession(c.Writer, session)
	res.Token = strings.Trim(strings.TrimPrefix(c.Writer.Header().Get("Authorization"), "Bearer "), " ")
	res.Msg = translation.TranslateIt("authenticated")
	res.StatusCode = http.StatusOK
	res.User = user
	c.JSON(res.StatusCode, res)
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
