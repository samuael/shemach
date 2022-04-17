package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/cmd/server/service"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type ISubscriberHandler interface {
	RegisterSubscriber(c *gin.Context)
	SubscriberLoginWithPhone(c *gin.Context)
	ConfirmRegistrationSubscription(c *gin.Context)
	ConfirmLoginSubscription(c *gin.Context)
}

type SubscriberHandler struct {
	Service       subscriber.ISubscriberService
	OtpService    *service.OtpService
	Authenticator auth.Authenticator
}

func NewSubscriberHandler(authenticator auth.Authenticator, service subscriber.ISubscriberService, otpService *service.OtpService) ISubscriberHandler {
	return &SubscriberHandler{
		Service:       service,
		OtpService:    otpService,
		Authenticator: authenticator,
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
		println(err.Error())
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
				res.Msg = translation.Translate(input.Lang, "internal server problem, please try again")
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
	status, er := shan.Service.RegisterTempoSubscriber(ctx)
	if er != nil {
		println(er.Error())
	}
	if status == state.STATUS_DBQUERY_ERROR || er != nil {
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
	input := &struct {
		Phone string `json:"phone"`
		Lang  string `json:"lang,omitempty"`
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
	var subscriber *model.Subscriber
	var era error
	passed := true
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
				ctx = context.WithValue(ctx, "subscriber_phone", input.Phone)
				subscriber, status, era = shan.Service.GetSubscriberByPhone(ctx)
				if status != state.STATUS_OK || era != nil {
					if era != nil {
						println(era.Error())
					}
					res.StatusCode = http.StatusInternalServerError
					res.Msg = translation.Translate("amh", "internal problem, please try again")
					c.JSON(res.StatusCode, res)
					return
				}
			} else if status == 2 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.Translate(input.Lang, "user with this phone is already in confirmation process")
				c.JSON(res.StatusCode, res)
				return
			} else {
				res.StatusCode = http.StatusNotFound
				res.Msg = translation.Translate("eng", "user account with this id doesn't exist")
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
	input.Lang = subscriber.Lang
	tempo := &model.TempoLoginSubscriber{
		Unix:  uint64(time.Now().Unix()),
		Phone: input.Phone,
	}
	randomNumber := helper.GenerateRandomString(5, helper.NUMBERS)
	tempo.Confirmation = randomNumber
	otpCode := &model.TeldaOTP{
		Phone:      tempo.Phone,
		OTP:        randomNumber,
		SenderName: translation.Translate(input.Lang, os.Getenv("SYSTEM_NAME")),
		Remark:     translation.Translate(input.Lang, "This is your login confirmation code from Agri-Info systems"),
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
	ctx = context.WithValue(ctx, "login_tempo_subscriber", tempo)
	status, er := shan.Service.RegisterTempoLoginSubcriber(ctx)
	if er != nil {
		println(er.Error())
	}
	if status == state.STATUS_DBQUERY_ERROR || er != nil {
		if err != nil {
			println(err.Error())
		}
		res.StatusCode = http.StatusInternalServerError
		res.Msg = translation.Translate(input.Lang, "internal problem, please try again")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Msg = translation.Translate(input.Lang, "You will recieve an SMS a message containing the login confirmation code\nplease confirm your account with in 30 minutes.")
	res.OTP = otpCode
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}

func (shan *SubscriberHandler) ConfirmRegistrationSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	input := &struct {
		Phone        string `json:"phone"`
		Confirmation string `json:"confirmation"`
	}{}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Subscriber *model.Subscriber `json:"subscriber"`
		Token      string            `json:"token,omitempty"`
		Msg        string            `json:"msg"`
	}{}
	jsonDecode := json.NewDecoder(c.Request.Body)
	era := jsonDecode.Decode(input)
	if era != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate("amh", "bad request values")
		c.JSON(int(res.StatusCode), res)
		return
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
				res.Msg = translation.Translate("amh", "internal server problem, plese try again")
				c.JSON(http.StatusInternalServerError, res)
				return
			} else if status == 1 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.TranslateIt("user already exist")
				c.JSON(res.StatusCode, res)
				return

			} else if status == 2 {
				// 	res.StatusCode = http.StatusConflict
				// 	res.Msg = translation.TranslateIt("user with this phone is already in confirmation process")
				// 	c.JSON(res.StatusCode, res)
				// 	return
				// }
			} else {
				res.StatusCode = http.StatusNotFound
				res.Msg = translation.TranslateIt("user account with this id doesn't exist")
				c.JSON(res.StatusCode, res)
				return
			}
		} else {
			res.Msg = "invalid phone number value"
			res.StatusCode = http.StatusNotFound
			c.JSON(res.StatusCode, res)
			return
		}
	}
	ctx = context.WithValue(ctx, "subscriber_phone", input.Phone)
	pendingRegistrationSubscription, status, er := shan.Service.GetPendingRegistrationSubscriptionByPhone(ctx)
	if status != state.STATUS_OK || er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("no pending subscription with this phone number")
		c.JSON(res.StatusCode, res)
		return
	}
	if input.Confirmation == pendingRegistrationSubscription.ConfirmationCode {

		subscriber := &model.Subscriber{
			Fullname:      pendingRegistrationSubscription.Fullname,
			Lang:          pendingRegistrationSubscription.Lang,
			Subscriptions: []uint8{},
			Role:          pendingRegistrationSubscription.Role,
			Phone:         pendingRegistrationSubscription.Phone,
		}
		ctx = context.WithValue(ctx, "subscriber", subscriber)
		sub, status, er := shan.Service.RegisterSubscriber(ctx)
		if er != nil || status != state.STATUS_OK {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = translation.
				Translate(pendingRegistrationSubscription.Lang, "internal problem, please try again later")
			c.JSON(res.StatusCode, res)
			return
		}
		ctx = context.WithValue(ctx, "subscription_id", uint64(pendingRegistrationSubscription.ID))
		shan.Service.DeletePendingRegistrationSubscriptionByID(ctx)
		res.Subscriber = sub
		res.StatusCode = http.StatusOK
		res.Msg = translation.Translate(sub.Lang, "You are registered succesfuly. welcome!")
		subscriberSession := &model.SubscriberSession{
			ID:       sub.ID,
			Phone:    sub.Phone,
			Fullname: sub.Fullname,
			Lang:     sub.Lang,
		}
		saved := shan.Authenticator.SaveSubscriberSession(c.Writer, subscriberSession)
		if !saved {
			println("Subscription was not succesful")
		}
		res.Token = strings.TrimPrefix(c.Writer.Header().Get("Authorization"), "Bearer ")
		c.JSON(res.StatusCode, res)
		return
	} else {
		res.StatusCode = http.StatusExpectationFailed
		res.Msg = translation.Translate(pendingRegistrationSubscription.Lang, "incorrect confirmation code")
		c.JSON(res.StatusCode, res)
		return
	}
}

func (shan *SubscriberHandler) ConfirmLoginSubscription(c *gin.Context) {
	ctx := c.Request.Context()
	input := &struct {
		Phone        string `json:"phone"`
		Confirmation string `json:"confirmation"`
	}{}
	res := &struct {
		StatusCode int               `json:"status_code"`
		Subscriber *model.Subscriber `json:"subscriber"`
		Token      string            `json:"token,omitempty"`
		Msg        string            `json:"msg"`
	}{}
	jsonDecode := json.NewDecoder(c.Request.Body)
	era := jsonDecode.Decode(input)
	if era != nil {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate("amh", "bad request values")
		c.JSON(int(res.StatusCode), res)
		return
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
				res.Msg = translation.Translate("amh", "internal server problem, plese try again")
				c.JSON(http.StatusInternalServerError, res)
				return
			} else if status == 1 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.TranslateIt("user already exist")
				c.JSON(res.StatusCode, res)
				return
			} else if status == 2 {
				res.StatusCode = http.StatusConflict
				res.Msg = translation.TranslateIt("please complete your registration!")
				c.JSON(res.StatusCode, res)
				return
			} else if status == 4 {
			} else {
				res.StatusCode = http.StatusNotFound
				res.Msg = translation.TranslateIt("user account with this id doesn't exist")
				c.JSON(res.StatusCode, res)
				return
			}
		} else {
			res.Msg = "invalid phone number value"
			res.StatusCode = http.StatusNotFound
			c.JSON(res.StatusCode, res)
			return
		}
	}
	ctx = context.WithValue(ctx, "subscriber_phone", input.Phone)
	pendingRegistrationSubscription, status, er := shan.Service.GetPendingLoginSubscriptionByPhone(ctx)
	if status != state.STATUS_OK || er != nil {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("no pending subscription with this phone number")
		c.JSON(res.StatusCode, res)
		return
	}
	if input.Confirmation == pendingRegistrationSubscription.Confirmation {
		subscriber, status, er := shan.Service.GetSubscriberByPhone(ctx)
		if status != state.STATUS_OK || er != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Msg = translation.TranslateIt("Internal problem, please try again later!")
			c.JSON(res.StatusCode, res)
			return
		}
		ctx = context.WithValue(ctx, "subscription_id", uint64(pendingRegistrationSubscription.ID))
		_, ers := shan.Service.DeletePendingLoginSubscriptionByID(ctx)
		if ers != nil {
			println(ers.Error())
		}
		res.Subscriber = subscriber
		res.StatusCode = http.StatusOK
		res.Msg = translation.Translate(subscriber.Lang, "You are registered succesfuly. welcome!")
		subscriberSession := &model.SubscriberSession{
			ID:             subscriber.ID,
			Phone:          subscriber.Phone,
			Fullname:       subscriber.Fullname,
			Lang:           subscriber.Lang,
			StandardClaims: jwt.StandardClaims{},
		}
		subscribed := shan.Authenticator.SaveSubscriberSession(c.Writer, subscriberSession)
		if !subscribed {
			log.Println("User Authentication Failed Bro")
		}
		res.Token = strings.TrimPrefix(c.Writer.Header().Get("Authorization"), "Bearer ")
		c.JSON(res.StatusCode, res)
		return
	} else {
		res.StatusCode = http.StatusExpectationFailed
		res.Msg = translation.TranslateIt("incorrect confirmation code")
		c.JSON(res.StatusCode, res)
		return
	}
}
