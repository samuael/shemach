package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-passwd/validator"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IUserHandler interface {
	ChangePassword(c *gin.Context)
}

type UserHandler struct {
	Service user.IUserService
}

func NewUserHandler(service user.IUserService) IUserHandler {
	return &UserHandler{Service: service}
}

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
	return
}
