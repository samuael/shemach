package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/telda_sms"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ISubscriberHandler interface {
	RegisterSubscriber(c *gin.Context)
}

type SubscriberHandler struct {
	Service subscriber.ISubscriberService
}

func NewSubscriberHandler(service subscriber.ISubscriberService) ISubscriberHandler {
	return &SubscriberHandler{
		Service: service,
	}
}

func (shan *SubscriberHandler) RegisterSubscriber(c *gin.Context) {
	input := &struct {
		Fullname string `json:"fullname"`
		Phone    string `json:"phone"`
		Role     uint8  `json:"role"`
		Lang     string `json:"lang"`
	}{}

	res := &struct {
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
		OTP        *model.TeldaOTP   `json:"otp"`
	}{}
	jsonDec := json.NewDecoder(c.Request.Body)
	err := jsonDec.Decode(input)
	if err != nil {
		res.Msg = "bad request body"
		res.StatusCode = http.StatusBadRequest
		c.JSON(res.StatusCode, res)
		return
	}
	passed := true
	if len(input.Fullname) < 7 || len(strings.Split(input.Fullname, " ")) < 2 {
		passed = false
		res.Errors["Full name"] = "Invalid fullname value"
	} else {
		if len(strings.Split(input.Fullname, " ")) >= 2 && len(strings.Split(input.Fullname, " ")[0]) < 3 {
			passed = false
			res.Errors["First name"] = "Invalid firstname value"
		}
		if len(strings.Split(input.Fullname, " ")) >= 2 && len(strings.Split(input.Fullname, " ")[1]) < 3 {
			passed = false
			res.Errors["Last name"] = "Invalid lastname value"
		}
	}

	ctx := c.Request.Context()
	if input.Phone != "" {
		if len(input.Phone) <= 13 && len(input.Phone) >= 10 && form.MatchesPattern(input.Phone, form.PhoneRX) {
			if strings.HasPrefix(input.Phone, "0") {
				input.Phone = strings.Replace(input.Phone, "0", "+251", 1)
			}
			ctx = context.WithValue(ctx, "input_phone", input.Phone)
			status, er := shan.Service.CheckTheExistanceOfPhone(ctx)
			if status == -1 || er != nil {
				res.StatusCode = http.StatusInternalServerError
				res.Msg = "internal problem happened"
				c.JSON(http.StatusInternalServerError, res)
				return
			} else if status >= 1 {
				res.StatusCode = http.StatusConflict
				res.Msg = "account with this phone already exist"
				res.Errors["phone"] = "account with this phone already exist"
				c.JSON(http.StatusConflict, res)
				return
			}
		} else {
			passed = false
			res.Errors["phone"] = "invalid phone number value"
		}
	}

	if !passed || len(res.Errors) > 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = "internal data problem, plese try again correcting the issues"
		c.JSON(res.StatusCode, res)
		return
	}

	tempo := &model.TempoSubscriber{
		Unix:     time.Now().Unix(),
		Fullname: input.Fullname,
		Phone:    input.Phone,
		Lang:     input.Lang,
		Role:     input.Role,
	}
	randomNumber := helper.GenerateRandomString(5, helper.NUMBERS)
	otpCode := &model.TeldaOTP{
		Phone:      tempo.Phone,
		OTP:        randomNumber,
		SenderName: os.Getenv("SYSTEM_NAME"),
		Remark:     translation.Translate(tempo.Lang, "This is your confirmation code from Agri-Info systems"),
	}
	res.OTP = otpCode
	otpResponse, err := telda_sms.SendOtp(otpCode)
	if err != nil || otpResponse.MsgShortMessage != "Success" {
		res.Msg = "internal problemm please try again"
		res.StatusCode = http.StatusInternalServerError
		c.JSON(res.StatusCode, res)
		return
	}
	c.JSON(http.StatusOK, res)
}
