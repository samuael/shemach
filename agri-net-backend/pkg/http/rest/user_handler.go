package rest

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-passwd/validator"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/agent"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
	"github.com/samuael/agri-net/agri-net-backend/pkg/superadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/mail"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IUserHandler interface {
	Login(c *gin.Context)
	ChangePassword(c *gin.Context)
	UpdateProfilePicture(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteProfilePicture(c *gin.Context)
	ConfirmTempoCXP(c *gin.Context)
	ConfirmEmail(c *gin.Context)
}

type UserHandler struct {
	Service           user.IUserService
	AgentService      agent.IAgentService
	MerchantService   merchant.IMerchantService
	SuperadminService superadmin.ISuperadminService
	AdminService      admin.IAdminService
	InfoadminService  infoadmin.IInfoadminService
	Authenticator     auth.Authenticator
	Templates         *template.Template
}

func NewUserHandler(
	templates *template.Template,
	service user.IUserService,
	authenticator auth.Authenticator,
	adminservice admin.IAdminService,
	superadminService superadmin.ISuperadminService,
	agentService agent.IAgentService,
	merchantService merchant.IMerchantService,
	infoadminService infoadmin.IInfoadminService,
) IUserHandler {
	return &UserHandler{
		Service:           service,
		Authenticator:     authenticator,
		Templates:         templates,
		MerchantService:   merchantService,
		AgentService:      agentService,
		AdminService:      adminservice,
		SuperadminService: superadminService,
		InfoadminService:  infoadminService,
	}
}

func (suhandler *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	input := &struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}
	jsonDecoder := json.NewDecoder(c.Request.Body)
	res := &struct {
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
		User       interface{}       `json:"user,omitempty"`
		Role       string            `json:"role,omitempty"`
		Token      string            `json:"token,omitempty"`
	}{
		Errors: map[string]string{},
	}
	ers := jsonDecoder.Decode(input)
	if ers == nil {
		input.Email = strings.Trim(input.Email, " ")
	}
	if ers != nil || !(form.MatchesPattern(input.Email, form.EmailRX) || form.MatchesPattern(input.Phone, form.PhoneRX)) || len(input.Password) < 4 {
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

	if input.Phone != "" && strings.HasPrefix(input.Phone, "09") {
		input.Phone = strings.Replace(input.Phone, "09", "+251", 0)
	}
	// check the existance of the user using his  email only.
	var user *model.User
	var role int
	var status int
	var er error
	ctx = context.WithValue(ctx, "user_email", input.Email)
	if form.MatchesPattern(input.Phone, form.PhoneRX) {
		log.Println("------- Phone----------------------------")
		user, role, status, er = suhandler.Service.GetUserByPhone(ctx, input.Phone)
		if er != nil {
			log.Println(er.Error())
		}
	}
	if form.MatchesPattern(input.Email, form.EmailRX) && user == nil {
		user, role, status, er = suhandler.Service.GetUserByEmailOrID(ctx)
	}
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
		res.Msg = translation.TranslateIt("invalid login{Phone or Email} or password")
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
		superadmin, _ := suhandler.SuperadminService.GetSuperadminByID(ctx, int(user.ID))
		if superadmin == nil {
			res.User = user
		} else {
			res.User = superadmin
		}
	} else if role == 2 {
		session.Role = state.INFO_ADMIN
		res.Role = state.INFO_ADMIN
		infoadmin, _ := suhandler.InfoadminService.GetInfoadminByID(ctx, user.ID)
		if infoadmin == nil {
			res.User = user
		} else {
			res.User = infoadmin
		}
	} else if role == 3 {
		session.Role = state.ADMIN
		res.Role = state.ADMIN
		admin, _ := suhandler.AdminService.GetAdminByID(ctx, user.ID)
		if admin == nil {
			res.User = user
		} else {
			res.User = admin
		}
	} else if role == 4 {
		session.Role = state.MERCHANT
		res.Role = state.MERCHANT
		merchant, _ := suhandler.MerchantService.GetMerchantByID(ctx, int(user.ID))
		if merchant == nil {
			res.User = user
		} else {
			res.User = merchant
		}
	} else if role == 5 {
		session.Role = state.AGENT
		res.Role = state.AGENT
		agent, _ := suhandler.AgentService.GetAgentByID(ctx, int(user.ID))
		if agent == nil {
			res.User = user
		} else {
			res.User = agent
		}
	}
	suhandler.Authenticator.SaveSession(c.Writer, session)
	res.Token = strings.Trim(strings.TrimPrefix(c.Writer.Header().Get("Authorization"), "Bearer "), " ")
	res.Msg = translation.TranslateIt("authenticated")
	res.StatusCode = http.StatusOK
	// res.User = user
	c.JSON(res.StatusCode, res)
}

// ChangePassword ...
func (uhandler *UserHandler) ChangePassword(c *gin.Context) {
	input := &struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	res := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"password"`
		Errors     map[string]string `json:"errors,omitempty"`
	}{
		Errors: map[string]string{},
	}
	ctx := c.Request.Context()
	jdecoder := json.NewDecoder(c.Request.Body)
	er := jdecoder.Decode(input)
	println(input.ConfirmPassword, input.NewPassword)
	failed := false
	if er != nil {
		failed = true
	}
	passwordValidator := validator.New(validator.MinLength(4, errors.New("password characters length must be >=4")), validator.MaxLength(10, errors.New("password characters length must be <=10")))
	err := passwordValidator.Validate(input.NewPassword)
	if err != nil {
		res.Errors["newPassword"] = translation.TranslateIt(err.Error())
		failed = true
	}
	if input.NewPassword != input.ConfirmPassword {
		res.Errors["confirmPassword"] = translation.TranslateIt("confirm password different from new password")
		failed = true
	}
	if failed {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("bad request body")
		c.JSON(res.StatusCode, res)
		return
	}
	session := ctx.Value("session").(*model.Session)
	ctx = context.WithValue(ctx, "user_id", uint64(session.ID))
	hash, er := helper.HashPassword(input.NewPassword)
	if hash == "" || er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("bad password characters please use chanracters, numbers, and spacial characters as $#!@%^&*()  with length of 4 to 8 characters only")
		c.JSON(res.StatusCode, res)
		return
	}
	ctx = context.WithValue(ctx, "new_password", hash)
	er = uhandler.Service.UpdatePassword(ctx)
	if er != nil {
		res.Msg = translation.TranslateIt("can't update the password, please try again")
		res.StatusCode = http.StatusInternalServerError
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.TranslateIt("password updated succesfuly")
	c.JSON(res.StatusCode, res)
}

// UpdateProfilePicture ...
func (uhandler *UserHandler) UpdateProfilePicture(c *gin.Context) {
	var header *multipart.FileHeader
	var erro error
	var oldImage string
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	erro = c.Request.ParseMultipartForm(99999999999)
	if erro != nil {
		println(erro.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	image, header, erro := c.Request.FormFile("image")
	if erro != nil {
		println(erro.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	defer image.Close()
	if helper.IsImage(header.Filename) {
		newName := "images/profile/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
		var newImage *os.File
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		if erro != nil {
		println(erro.Error())

			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		defer newImage.Close()
		ctx = context.WithValue(ctx, "user_id", uint64(session.ID))
		oldImage = uhandler.Service.GetImageUrl(ctx)
		_, er := io.Copy(newImage, image)
		if er != nil {
		println(er.Error())

			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, "image_url", newName)
		success := uhandler.Service.ChangeImageUrl(ctx)
		if success {
			if oldImage != "" {
				if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + oldImage)
				} else {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + oldImage)
				}
			}
			c.JSON(http.StatusOK, &model.ShortSuccess{Msg: newName})
			return
		}
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		if er != nil{
			println(er.Error())
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.Writer.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// DeleteProfilePicture ...
func (uhandler *UserHandler) DeleteProfilePicture(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	ctx = context.WithValue(ctx, "user_id", uint64(session.ID))
	ctx = context.WithValue(ctx, "image_url", "")
	imageUrl := uhandler.Service.GetImageUrl(ctx)
	success := uhandler.Service.ChangeImageUrl(ctx)
	if success {
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			os.Remove(os.Getenv("ASSETS_DIRECTORY") + imageUrl)
		} else {
			os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + imageUrl)
		}
		c.JSON(http.StatusOK,
			&model.ShortSuccess{
				Msg: "Succesfully Deleted",
			})
		return
	} else {
		c.JSON(http.StatusInternalServerError, nil)
	}
}

func (uhandler *UserHandler) GenerateEmailConfirmationInformation(response http.ResponseWriter,
	confirmation model.EmailConfirmation) (string, bool) {
	emailConfirmationSession := &model.EmailConfirmationSession{
		EmailConfirmation: &confirmation,
		StandardClaims:    jwt.StandardClaims{},
	}
	//
	// uhandler.
	token, success := uhandler.Authenticator.SaveEmailConfirmationSession(emailConfirmationSession)
	if !success {
		return token, false
	}
	response.Header().Set("Authorization", "Bearer "+token)
	return token, true
}

// UpdateProfile ...
func (uhandler *UserHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Lang      string `json:"lang"`
	}{}
	res := &struct {
		Msg        string
		StatusCode int
		Errors     map[string]string
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(input)
	if er != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.TranslateIt("bad request")
		c.JSON(res.StatusCode, res)
		return
	}
	ctx = context.WithValue(ctx, "user_id", uint64(session.ID))
	user, _, status, er := uhandler.Service.GetUserByEmailOrID(ctx)
	updated := false
	if er != nil || status != state.STATUS_OK {
		if er != nil {
			println()
		}
		res.Msg = translation.TranslateIt("record not found")
		res.StatusCode = http.StatusNotFound
		c.JSON(res.StatusCode, res)
		return
	}
	if len(strings.Trim(input.Firstname, " ")) > 3 {
		user.Firstname = input.Firstname
		updated = true
	}
	if len(strings.Trim(input.Lastname, " ")) > 3 {
		user.Lastname = input.Lastname
		updated = true
	}

	if helper.MatchesPattern(input.Phone, helper.PhoneRX) && input.Phone != user.Phone {
		updated = true
		user.Phone = input.Phone
	}
	if helper.MatchesPattern(input.Email, helper.EmailRX) && input.Email != user.Email {
		// user.Email = input.Email
		updated = true
	}
	if !updated {
		res.StatusCode = http.StatusNotModified
		res.Msg = translation.TranslateIt("user not modified ")
		c.JSON(res.StatusCode, res)
		return
	}
	if len(strings.Trim(input.Email, " ")) > 0 && helper.MatchesPattern(input.Email, helper.EmailRX) && (input.Email != user.Email) {
		// check the existance of an email
		// user , _ , _ ,  er := uhandler.Service.
		emailInConfirmation := &model.EmailConfirmation{
			UserID:       user.ID,
			Email:        input.Email,
			OldEmail:     user.Email,
			CreatedAt:    uint64(time.Now().Unix()),
			IsNewAccount: false,
		}
		status, err := uhandler.Service.SaveEmailConfirmation(ctx, emailInConfirmation)
		if err != nil {
			message := translation.TranslateIt("update was not succesful, please trya again later")
			if status < 0 {
				if status == -1 {
					res.StatusCode = http.StatusNotFound
				} else if status == -2 {
					res.StatusCode = http.StatusConflict
				}
				message = translation.Translate(session.Lang, err.Error())
			} else {
				res.StatusCode = http.StatusOK
			}
			res.Msg = message
			res.Errors["email"] = translation.Translate(session.Lang, err.Error())
			c.JSON(res.StatusCode, res)
			return
		}
		emailInConfirmationSession := &model.EmailConfirmationSession{
			EmailConfirmation: emailInConfirmation,
			StandardClaims:    jwt.StandardClaims{},
		}
		tokent, success := uhandler.Authenticator.SaveEmailConfirmationSession(emailInConfirmationSession)
		if !success && tokent == "" {
			message := translation.TranslateIt("internal problem, please trya again later")
			println("Save Email Confirmation Session")
			res.Msg = message
			res.StatusCode = http.StatusInternalServerError
			c.JSON(res.StatusCode, res)
			return
		}
		println(" The Token :", tokent)
		// Now I am gonna send the Confirmation Link and the token to e clicked by the email owner through his email.
		success = mail.ConfirmUpdateEmailAccount(c.Writer, []string{input.Email}, tokent, user.Firstname, "127.0.0.1")
		if !success {
			res.Msg = translation.TranslateIt("internal problem, please trya again later")
			res.StatusCode = http.StatusInternalServerError
			c.JSON(res.StatusCode, res)
			return
		}
		updated = true
	}
	status, er = uhandler.Service.UpdateUser(ctx, user)
	if (er != nil) || (status != state.STATUS_OK) {
		if er != nil {
			println(er.Error())
		}
		res.StatusCode = http.StatusInternalServerError
		res.Msg = translation.Translate(session.Lang, "internal problem, please try again later")
		c.JSON(res.StatusCode, res)
		return
	}
	res.StatusCode = http.StatusOK
	res.Msg = translation.Translate(session.Lang, "user updated succesfuly")
	c.JSON(res.StatusCode, res)
}

// ConfirmEmail  ...
func (uhandler *UserHandler) ConfirmEmail(c *gin.Context) {
	ctx := c.Request.Context()
	// token := strings.Trim(strings.TrimPrefix(c.Request.Header["Authorization"][0], "Bearer "), " ")
	// if token == "" {
	token := c.Query("token")
	// }
	session, er := uhandler.Authenticator.GetEmailSession(token)
	if er != nil || session == nil {
		uhandler.Templates.ExecuteTemplate(c.Writer,
			"message.html",
			&struct {
				Message string
			}{
				Message: "Unauthorized Access!",
			})
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println(session.Email, session.OldEmail, session.ID, session.IsNewAccount)
	er = uhandler.Service.ConfirmUserEmailUpdate(ctx, session.ID, session.Email, session.OldEmail)
	if er != nil {

		if strings.Contains(er.Error(), "duplicate key value violates unique constraint \"users_email_ke\"") {
			uhandler.Templates.ExecuteTemplate(c.Writer,
				"message.html",
				&struct {
					Message string
				}{
					Message: "You can't have several role in our System!",
				})
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		uhandler.Templates.ExecuteTemplate(c.Writer,
			"message.html",
			&struct {
				Message string
			}{
				Message: "Confirmation was not succesful!",
			})
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	uhandler.Templates.ExecuteTemplate(c.Writer,
		"message.html",
		&struct {
			Message string
		}{
			Message: "Succesfuly confirmed",
		})
	c.Writer.WriteHeader(http.StatusOK)
}

// ConfirmTempoCXP
func (uhandler *UserHandler) ConfirmTempoCXP(c *gin.Context) {
	ctx := c.Request.Context()
	input := &struct {
		Phone string `json:"phone"`
		Code  string `json:"otp_code"`
	}{}
	res := &struct {
		StatusCode int         `json:"status_code"`
		Msg        string      `json:"msg"`
		Role       string      `json:"role"`
		User       *model.User `json:"user,omitempty"`
		Token      string      `json:"token,omitempty"`
	}{}

	jsonDec := json.NewDecoder(c.Request.Body)
	er := jsonDec.Decode(input)
	if input.Phone == "" || er != nil || len(input.Code) > 5 {
		res.Msg = translation.TranslateIt("missing important parameter")
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	input.Phone = strings.Trim(input.Phone, " ")
	if !form.MatchesPattern(input.Phone, form.PhoneRX) {
		res.Msg = "invalid \"phone\" address"
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	var tcxp model.TempoCXP
	err := uhandler.Service.GetTempoCXP(ctx, input.Phone, tcxp)
	if err != nil || &(tcxp) == nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("target user not found ")
		c.JSON(res.StatusCode, res)
		return
	}

	user, role, status, er := uhandler.Service.GetUserByPhone(ctx, input.Phone)
	if status != state.STATUS_OK || er != nil || user == nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("can't find any user with the specified phone. it may be because your confirmation expired")
		c.JSON(res.StatusCode, res)
		return
	}

	err = uhandler.Service.RemoveTempoCXP(ctx, input.Phone)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Msg = translation.TranslateIt("internal problem, please try again later")
		c.JSON(res.StatusCode, res)
		return
	}
	session := &model.Session{
		ID:    user.ID,
		Email: user.Email,
		Lang:  user.Lang,
	}
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
	uhandler.Authenticator.SaveSession(c.Writer, session)
	res.Token = strings.Trim(strings.TrimPrefix(c.Writer.Header().Get("Authorization"), "Bearer "), " ")
	res.Msg = translation.TranslateIt("authenticated")
	res.StatusCode = http.StatusOK
	res.User = user
	c.JSON(res.StatusCode, res)
}
