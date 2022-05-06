package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-passwd/validator"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/mail"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IUserHandler interface {
	ChangePassword(c *gin.Context)
	UpdateProfilePicture(c *gin.Context)
	UpdateProfile(c *gin.Context)
	DeleteProfilePicture(c *gin.Context)
}

type UserHandler struct {
	Service       user.IUserService
	Authenticator auth.Authenticator
}

func NewUserHandler(service user.IUserService,
	authenticator auth.Authenticator,
) IUserHandler {
	return &UserHandler{
		Service:       service,
		Authenticator: authenticator,
	}
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
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	image, header, erro := c.Request.FormFile("image")
	if erro != nil {
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
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		defer newImage.Close()
		ctx = context.WithValue(ctx, "user_id", uint64(session.ID))
		oldImage = uhandler.Service.GetImageUrl(ctx)
		_, er := io.Copy(newImage, image)
		if er != nil {
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
	println(input.Email, user.Email)
	if len(strings.Trim(input.Email, " ")) > 0 && helper.MatchesPattern(input.Email, helper.EmailRX) && (input.Email != user.Email) {
		println("Email Changed")
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
				res.StatusCode = http.StatusNotModified
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
		success = mail.ConfirmUpdateEmailAccount([]string{input.Email}, tokent, user.Firstname, "127.0.0.1")
		if !success {
			println("Sending an Email ")
			res.Msg = translation.TranslateIt("internal problem, please trya again later")
			res.StatusCode = http.StatusInternalServerError
			c.JSON(res.StatusCode, res)
			return
		}
		updated = true
	}
	status, er = uhandler.Service.UpdateUser(ctx, user)
	if (er != nil) || (status != state.STATUS_OK) {
		print("Updating User ", er.Error())
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

func (uhandler *UserHandler) ConfirmEmail(c *gin.Context) {
	//
	// ctx := c.Request.Context()
	//
}
