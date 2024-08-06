package rest

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/merchant"
	"github.com/samuael/shemach/shemach-backend/pkg/payment"
	"github.com/samuael/shemach/shemach-backend/pkg/user"
	"github.com/samuael/shemach/shemach-backend/platforms/form"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
	"github.com/samuael/shemach/shemach-backend/platforms/telda_sms"
	"github.com/samuael/shemach/shemach-backend/platforms/translation"
)

type IMerchantHandler interface {
	RegisterMerchant(c *gin.Context)
	SubscribeForProduct(c *gin.Context)
	UnsubscriberForProduct(c *gin.Context)
	MerchantsSearch(c *gin.Context)
	DeleteMerchantByID(c *gin.Context)
}

type MerchantHandler struct {
	Service        merchant.IMerchantService
	UserService    user.IUserService
	PaymentService payment.IPaymentService
}

func NewMerchantHandler(
	service merchant.IMerchantService,
	userser user.IUserService,
	paymentservice payment.IPaymentService,
) IMerchantHandler {
	return &MerchantHandler{
		Service:        service,
		UserService:    userser,
		PaymentService: paymentservice,
	}
}

func (mhandler *MerchantHandler) RegisterMerchant(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	input := &struct {
		Firstname string         `json:"firstname"`
		Lastname  string         `json:"lastname"`
		Phone     string         `json:"phone"`
		Address   *model.Address `json:"address"`
		Lang      string         `json:"lang"`
	}{}
	resp := &struct {
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Errors     map[string]string `json:"errors"`
		Merchant   *model.Merchant   `json:"merchant"`
		OTP        *model.TeldaOTP   `json:"otp"`
	}{
		Errors: map[string]string{},
	}
	if er := c.BindJSON(input); er == nil {
		if input.Lang == "" {
			input.Lang = "amh"
		}
		fail := false
		if !form.MatchesPattern(input.Phone, form.PhoneRX) {
			resp.Errors["phone"] = "invalid phone address"
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
			if user, _, _, err := mhandler.UserService.GetUserByPhone(ctx, input.Phone); user != nil && err == nil {
				resp.Msg = "account with this phone already exist."
				resp.StatusCode = http.StatusUnauthorized
				c.JSON(resp.StatusCode, resp)
				return
			}
			if input.Phone != "" {
				if len(input.Phone) <= 13 && len(input.Phone) >= 10 && form.MatchesPattern(input.Phone, form.PhoneRX) {
					if strings.HasPrefix(input.Phone, "0") {
						input.Phone = strings.Replace(input.Phone, "0", "+251", 1)
					}
					if user, _, _, err := mhandler.UserService.GetUserByPhone(ctx, input.Phone); user != nil && err == nil {
						resp.Msg = "account with this phone already exist."
						resp.StatusCode = http.StatusUnauthorized
						c.JSON(resp.StatusCode, resp)
						return
					}
				} else {
					fail = true
					resp.Errors["phone"] = "invalid phone number value"
				}
			}
		}
		if fail || len(resp.Errors) > 0 {
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = translation.Translate(input.Lang, "bad request information")
			c.JSON(resp.StatusCode, resp)
			return
		}
		// chech payment for the
		// _, er := mhandler.PaymentService.ValidateInvoice(ctx, &model.HellocashInvoiceRequest{
		// 	Amount:      50,
		// 	Description: "Please complete this payment to continue to the transaction!",
		// 	From:        input.Phone,
		// 	Currency:    "ETB",
		// })
		// if er != nil && !strings.Contains(er.Error(), "undefined account") {
		// 	resp.StatusCode = http.StatusInternalServerError
		// 	resp.Msg = translation.Translate(input.Lang, "internal problem, please try again later!")
		// 	c.JSON(resp.StatusCode, resp)
		// 	return
		// } else if er != nil && strings.Contains(er.Error(), "undefined account") {
		// 	resp.StatusCode = http.StatusExpectationFailed
		// 	resp.Msg = translation.Translate(input.Lang, "there is no valid payment account with this phone number")
		// 	c.JSON(resp.StatusCode, resp)
		// 	return
		// }
		ctx, _ = context.WithTimeout(ctx, time.Second*15)
		tempo := &model.TempoCXP{
			CreatedAt: uint64(time.Now().Unix()),
			Phone:     input.Phone,
			Role:      state.MERCHANT_ROLE_INT,
		}
		randomNumber := helper.GenerateRandomString(5, helper.NUMBERS)
		tempo.Confirmation = randomNumber
		otpCode := &model.TeldaOTP{
			Phone:      tempo.Phone,
			OTP:        randomNumber,
			SenderName: translation.Translate(session.Lang, os.Getenv("SYSTEM_NAME")),
			Remark:     translation.Translate(session.Lang, `This is your confirmation and temporary password code`),
		}
		// mhandler.OtpService <- otpResponse
		er = mhandler.UserService.RegisterTempoCXP(ctx, tempo)
		if er != nil {
			if strings.Contains(er.Error(), "duplicate key value violates unique constraint") {
				println(er.Error())
				resp.Msg = translation.Translate(session.Lang, "merchant account with this information already in the process of confirmation")
				resp.StatusCode = http.StatusConflict
			} else {
				resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later")
				resp.StatusCode = http.StatusInternalServerError
			}
			c.JSON(resp.StatusCode, resp)
			return
		}

		otpResponse, err := telda_sms.SendOtp(otpCode)
		if err != nil || !strings.EqualFold(otpResponse.MsgShortMessage, "Success") {
			resp.Msg = translation.Translate(session.Lang, "internal problem, please try again")
			resp.StatusCode = http.StatusInternalServerError
			resp.OTP = nil
			c.JSON(resp.StatusCode, resp)
			return
		}

		hashed, _ := helper.HashPassword(randomNumber)

		merchant := &model.Merchant{}
		merchant.Firstname = input.Firstname
		merchant.Lastname = input.Lastname
		merchant.Phone = input.Phone
		merchant.CreatedAt = uint64(time.Now().Unix())
		merchant.RegisteredBy = uint64(session.ID)
		merchant.Password = hashed
		merchant.Address = input.Address
		merchant.Lang = session.Lang
		if merchant.Address == nil {
			merchant.Address = &model.Address{}
		}
		status, _ := mhandler.Service.RegisterMerchant(ctx, merchant)
		if er != nil {
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.TranslateIt(er.Error())
		}
		if status == -1 {
			resp.StatusCode = http.StatusUnauthorized
			resp.Msg = translation.TranslateIt("you are not authorized to create this merchant instance")
		} else if status == -2 {
			// missing aaddress information
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = translation.Translate(session.Lang, "missing important address information, please complete necessary datas.")
		} else if status == -3 {
			// internal problem    // "email" : "samuaeladnew@gmail.com",

			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later")
		} else if status == -4 {
			resp.StatusCode = http.StatusConflict
			resp.Msg = translation.Translate(session.Lang, "merchant with this information had already registered")

		} else if status > 0 {
			println(status)
			resp.Msg = translation.Translate(session.Lang, `You will recieve an SMS a message containing the confirmation code\nplease confirm your phone number with in 30 minutes.\n The Confirmation numbe also serve as your password`)
			resp.OTP = otpCode
			resp.Merchant = merchant
			resp.StatusCode = http.StatusOK
		}
		c.JSON(resp.StatusCode, resp)
	} else {
		println(er.Error())
		resp.Msg = translation.TranslateIt("bad request body for merchant creation")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
		return
	}
}

// SubscribeForProduct ...
func (mhandler *MerchantHandler) SubscribeForProduct(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg,omitempty"`
	}{}
	session := ctx.Value("session").(*model.Session)
	productID, er := strconv.Atoi(c.Param("id"))
	if er != nil || productID <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "invalid product id")
		c.JSON(res.StatusCode, res)
		return
	}
	status := mhandler.Service.CreateSubscriptions(ctx, uint8(productID), session.ID)
	if status == -1 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "merchant with this id doesn't exist")
	} else if status == -2 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "product with this id doesn't exist")
	} else if status == -3 {
		res.StatusCode = http.StatusConflict
		res.Msg = translation.Translate(session.Lang, "you have already subscribed for a product")
	}
	if status < 0 {
		c.JSON(res.StatusCode, res)
		return
	}
	res.Msg = translation.Translate(session.Lang, "subscription created succesfuly")
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}

// UnsubscriberForProduct
func (mhandler *MerchantHandler) UnsubscriberForProduct(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg,omitempty"`
	}{}
	session := ctx.Value("session").(*model.Session)
	productID, er := strconv.Atoi(c.Param("id"))
	if er != nil || productID <= 0 {
		res.StatusCode = http.StatusBadRequest
		res.Msg = translation.Translate(session.Lang, "invalid product id")
		c.JSON(res.StatusCode, res)
		return
	}
	status := mhandler.Service.UnsubscribeProduct(ctx, uint8(productID), session.ID)
	if status == -1 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "subscriber with this id doesn't exist")
	} else if status == -2 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "product with this id doesn't exist")
	} else if status == -3 {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.Translate(session.Lang, "you have not subscribed for this product yet")
	}
	if status < 0 {
		c.JSON(res.StatusCode, res)
		return
	}
	res.Msg = translation.Translate(session.Lang, "un-subscribed succesfuly")
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}

// MerchantsSearch
func (mhandler *MerchantHandler) MerchantsSearch(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	response := &struct {
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
		Merchants  []*model.Merchant `json:"merchants"`
	}{
		Merchants: []*model.Merchant{},
	}
	phone := c.Query("phone")
	name := c.Query("name")
	createdBy, er := strconv.Atoi(c.Query("created_by"))
	if er != nil {
		createdBy = 0
	}
	offset, er := strconv.Atoi(c.Query("offset"))
	if er != nil {
		offset = 0
	}
	limit, er := strconv.Atoi(c.Query("limit"))
	if er != nil {
		limit = offset + 10
	}
	merchants, er := mhandler.Service.SearchMerchants(ctx, phone, name, uint64(createdBy), uint(offset), uint(limit))
	if er != nil && strings.Contains(er.Error(), "no record was deleted") {
		response.StatusCode = http.StatusNotFound
		response.Msg = translation.Translate(session.Lang, "mercahts not found ")
	} else if er != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = translation.Translate(session.Lang, "can't delete the merchant instance")
	} else {
		response.StatusCode = http.StatusOK
		response.Msg = translation.Translate(session.Lang, "merchants found")
		response.Merchants = merchants
	}
	c.JSON(response.StatusCode, response)
}

func (mhandler *MerchantHandler) DeleteMerchantByID(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	response := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg"`
	}{}
	merchantid, er := strconv.Atoi(c.Param("id"))
	if er != nil || merchantid <= 0 {
		response.Msg = translation.Translate(session.Lang, "invalid merchant id")
		response.StatusCode = http.StatusBadRequest
		c.JSON(response.StatusCode, response)
		return
	}
	er = mhandler.Service.DeleteMerchantByID(ctx, uint64(merchantid))
	if er != nil {
		response.Msg = translation.Translate(session.Lang, "can't found a merchant instance with this id")
		response.StatusCode = http.StatusNotFound
		c.JSON(response.StatusCode, response)
		return
	}
	response.Msg = translation.Translate(session.Lang, "merchant deleted succesfuly")
	response.StatusCode = http.StatusOK
	c.JSON(response.StatusCode, response)
	return
}
