package rest

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/telda_sms"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IMerchantHandler interface {
	RegisterMerchant(c *gin.Context)
}

type MerchantHandler struct {
	Service     merchant.IMerchantService
	UserService user.IUserService
}

func NewMerchantHandler(
	service merchant.IMerchantService,
	userser user.IUserService) IMerchantHandler {
	return &MerchantHandler{
		Service:     service,
		UserService: userser,
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
		er := mhandler.UserService.RegisterTempoCXP(ctx, tempo)
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
		if status == -1 {
			// unauthorized
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
