package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/cmd/server/service"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ISubscriberHandler interface {
	RegisterSubscriber(c *gin.Context)
}

type SubscriberHandler struct {
	Service    subscriber.ISubscriberService
	OtpService *service.OtpService
}

func NewSubscriberHandler(service subscriber.ISubscriberService, otpService *service.OtpService) ISubscriberHandler {
	return &SubscriberHandler{
		Service:    service,
		OtpService: otpService,
	}
}

func (shan *SubscriberHandler) RegisterSubscriber(c *gin.Context) {
	input := &struct {
		Fullname string `json:"fullname"`
		Phone    string `json:"phone"`
		Role     uint8  `json:"role"`
		Lang     string `json:"lang"`
	}{}
	ctx := c.Request.Context()
	res := &struct {
		Msg        string            `json:"msg"`
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
		OTP        *model.TeldaOTP   `json:"otp,omitempty"`
	}{
		Errors: map[string]string{},
	}
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
	if input.Phone != "" {
		if len(input.Phone) <= 13 && len(input.Phone) >= 10 && form.MatchesPattern(input.Phone, form.PhoneRX) {
			if strings.HasPrefix(input.Phone, "0") {
				input.Phone = strings.Replace(input.Phone, "0", "+251", 1)
			}
			ctx = context.WithValue(ctx, "phone", input.Phone)
			status, er := shan.Service.CheckTheExistanceOfPhone(ctx)
			if status == 0 || er != nil {
				if er != nil {
					println(er.Error())
				}
				res.StatusCode = http.StatusInternalServerError
				res.Msg = translation.Translate(input.Lang, "internal server problem, plese try again")
				c.JSON(http.StatusInternalServerError, res)
				return
			} else if status == 1 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.Translate(input.Lang, "account with this phone already exist")
				res.Errors["phone"] = translation.Translate(input.Lang, "account with this phone already exist")
				c.JSON(http.StatusConflict, res)
				return
			} else if status == 2 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.Translate(input.Lang, "user with this phone is already in confirmation process")
				c.JSON(res.StatusCode, res)
				return
			}
		} else {
			passed = false
			res.Errors["phone"] = "invalid phone number value"
		}
	}
	if !passed || len(res.Errors) > 0 {
		if err != nil {
			println(err.Error())
		}
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(input.Lang, "internal server problem, plese try again")
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
	tempo.ConfirmationCode = randomNumber
	otpCode := &model.TeldaOTP{
		Phone:      tempo.Phone,
		OTP:        randomNumber,
		SenderName: translation.Translate(tempo.Lang, os.Getenv("SYSTEM_NAME")),
		Remark:     translation.Translate(tempo.Lang, "This is your confirmation code from Agri-Info systems"),
	}
	// otpResponse, err := telda_sms.SendOtp(otpCode)
	// if err != nil || otpResponse.MsgShortMessage != "Success" {
	// 	if err != nil {
	// 		println(err.Error())
	// 	}
	// 	res.Msg = translation.Translate(tempo.Lang, "internal problem, please try again")
	// 	res.StatusCode = http.StatusInternalServerError
	// 	res.OTP = nil
	// 	c.JSON(res.StatusCode, res)
	// 	return
	// }
	// println(otpResponse)
	// shan.OtpService <- otpResponse
	ctx = context.WithValue(ctx, "tempo_subscriber", tempo)
	status, er := shan.Service.RegisterSubscriber(ctx)
	if er != nil {
		println(er.Error())
	}
	if status == state.DT_STATUS_DBQUERY_ERROR || er != nil {
		if err != nil {
			println(err.Error())
		}
		res.StatusCode = http.StatusInternalServerError
		res.Msg = translation.Translate(tempo.Lang, "internal problem, please try again")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Msg = translation.Translate(tempo.Lang, "You will recieve an SMS a message containing the confirmation code\nplease confirm your account with in 30 minutes.")
	res.OTP = otpCode
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}

func (shan *SubscriberHandler) SubscriberLoginWithPhone(c *gin.Context) {
	// input := &struct {
	// 	Phone string `json:"phone"`
	// }{}
	// ctx := c.Request.Context()
	// res := &struct {
	// 	Msg        string            `json:"msg"`
	// 	Errors     map[string]string `json:"errors"`
	// 	StatusCode int               `json:"status_code"`
	// 	OTP        *model.TeldaOTP   `json:"otp,omitempty"`
	// }{
	// 	Errors: map[string]string{},
	// }
	// jsonDec := json.NewDecoder(c.Request.Body)
	// err := jsonDec.Decode(input)
	// if err != nil {
	// 	res.Msg = "bad request body"
	// 	res.StatusCode = http.StatusBadRequest
	// 	c.JSON(res.StatusCode, res)
	// 	return
	// }
	// passed := true
	// if input.Phone != "" {
	// 	if len(input.Phone) <= 13 && len(input.Phone) >= 10 && form.MatchesPattern(input.Phone, form.PhoneRX) {
	// 		if strings.HasPrefix(input.Phone, "0") {
	// 			input.Phone = strings.Replace(input.Phone, "0", "+251", 1)
	// 		}
	// 		ctx = context.WithValue(ctx, "phone", input.Phone)
	// 		status, er := shan.Service.CheckTheExistanceOfPhone(ctx)
	// 		if status == 0 || er != nil {
	// 			if er != nil {
	// 				println(er.Error())
	// 			}
	// 			res.StatusCode = http.StatusInternalServerError
	// 			res.Msg = translation.Translate(input.Lang, "internal server problem, plese try again")
	// 			c.JSON(http.StatusInternalServerError, res)
	// 			return
	// 		} else if status == 1 {

	// 			//
	// 		} else if status == 2 {
	// 			res.StatusCode = http.StatusConflict
	// 			res.Msg = translation.Translate(input.Lang, "user with this phone is already in confirmation process")
	// 			c.JSON(res.StatusCode, res)
	// 			return
	// 		} else {
	// 			res.StatusCode = http.StatusNotFound
	// 			res.Msg = translation.Translate("eng", "user account with this id doesn't exist")
	// 			c.JSON(res.StatusCode, res)
	// 			return
	// 		}
	// 	} else {
	// 		passed = false
	// 		res.Errors["phone"] = "invalid phone number value"
	// 	}
	// }
	// if !passed || len(res.Errors) > 0 {
	// 	if err != nil {
	// 		println(err.Error())
	// 	}
	// 	res.StatusCode = http.StatusBadRequest
	// 	res.Msg = translation.Translate(input.Lang, "internal server problem, plese try again")
	// 	c.JSON(res.StatusCode, res)
	// 	return
	// }
	// tempo := &model.TempoSubscriberLogin{
	// 	Unix:  uint64(time.Now().Unix()),
	// 	Phone: input.Phone,
	// }
	// randomNumber := helper.GenerateRandomString(5, helper.NUMBERS)
	// tempo.ConfirmationCode = randomNumber
	// otpCode := &model.TeldaOTP{
	// 	Phone:      tempo.Phone,
	// 	OTP:        randomNumber,
	// 	SenderName: translation.Translate(tempo.Lang, os.Getenv("SYSTEM_NAME")),
	// 	Remark:     translation.Translate(tempo.Lang, "This is your confirmation code from Agri-Info systems"),
	// }
	// // otpResponse, err := telda_sms.SendOtp(otpCode)
	// // if err != nil || otpResponse.MsgShortMessage != "Success" {
	// // 	if err != nil {
	// // 		println(err.Error())
	// // 	}
	// // 	res.Msg = translation.Translate(tempo.Lang, "internal problem, please try again")
	// // 	res.StatusCode = http.StatusInternalServerError
	// // 	res.OTP = nil
	// // 	c.JSON(res.StatusCode, res)
	// // 	return
	// // }
	// // println(otpResponse)
	// ctx = context.WithValue(ctx, "tempo_subscriber", tempo)
	// status, er := shan.Service.RegisterSubscriber(ctx)
	// if er != nil {
	// 	println(er.Error())
	// }
	// if status == state.DT_STATUS_DBQUERY_ERROR || er != nil {
	// 	if err != nil {
	// 		println(err.Error())
	// 	}
	// 	res.StatusCode = http.StatusInternalServerError
	// 	res.Msg = translation.Translate(tempo.Lang, "internal problem, please try again")
	// 	c.JSON(res.StatusCode, res)
	// 	return
	// }
	// res.Msg = translation.Translate(tempo.Lang, "You will recieve an SMS a message containing the confirmation code\nplease confirm your account with in 30 minutes.")
	// res.OTP = otpCode
	// res.StatusCode = http.StatusOK
	// c.JSON(res.StatusCode, res)
}

func (shan *SubscriberHandler) ConfirmRegistrationSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	println(ctx)
}

func (shan *SubscriberHandler) ConfirmLoginSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	println(ctx)
}
