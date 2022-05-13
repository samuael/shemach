package rest

import (
	"encoding/json"
	"net/http"
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
