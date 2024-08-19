package rest

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/agent"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/payment"
	"github.com/samuael/shemach/shemach-backend/pkg/user"
	"github.com/samuael/shemach/shemach-backend/platforms/form"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
	"github.com/samuael/shemach/shemach-backend/platforms/telda_sms"
	"github.com/samuael/shemach/shemach-backend/platforms/translation"
)

type IAgentHandler interface {
	RegisterAgent(c *gin.Context)
	// GetAgentByID(c *gin.Context)
	AgentsSearch(c *gin.Context)
	DeleteAgentByID(c *gin.Context)
}

type AgentHandler struct {
	Service        agent.IAgentService
	UserService    user.IUserService
	PaymentService payment.IPaymentService
}

func NewAgentHandler(
	service agent.IAgentService,
	userser user.IUserService,
) IAgentHandler {
	return &AgentHandler{
		Service:     service,
		UserService: userser,
	}
}

func (ahandler *AgentHandler) RegisterAgent(c *gin.Context) {
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
		Agent      *model.Agent      `json:"agent"`
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
			if user, _, _, err := ahandler.UserService.GetUserByPhone(ctx, input.Phone); user != nil && err == nil {
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
					if user, _, _, err := ahandler.UserService.GetUserByPhone(ctx, input.Phone); user != nil && err == nil {
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
		_, er := ahandler.PaymentService.ValidateInvoice(ctx, &model.HellocashInvoiceRequest{
			Amount:      50,
			Description: "Please complete this payment to continue to the transaction!",
			From:        input.Phone,
			Currency:    "ETB",
		})
		if er != nil && !strings.Contains(er.Error(), "undefined account") {
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.Translate(input.Lang, "internal problem, please try again later!")
			c.JSON(resp.StatusCode, resp)
			return
		} else if er != nil && strings.Contains(er.Error(), "undefined account") {
			resp.StatusCode = http.StatusExpectationFailed
			resp.Msg = translation.Translate(input.Lang, "there is no valid account with this phone number")
			c.JSON(resp.StatusCode, resp)
			return
		}
		ctx, _ = context.WithTimeout(ctx, time.Second*15)
		tempo := &model.TempoCXP{
			CreatedAt: uint64(time.Now().Unix()),
			Phone:     input.Phone,
			Role:      state.AGENT_ROLE_INT,
		}
		randomNumber := helper.GenerateRandomString(5, helper.NUMBERS)
		tempo.Confirmation = randomNumber
		otpCode := &model.TeldaOTP{
			Phone:      tempo.Phone,
			OTP:        randomNumber,
			SenderName: translation.Translate(session.Lang, os.Getenv("SYSTEM_NAME")),
			Remark:     translation.Translate(session.Lang, `This is your confirmation and temporary password code`),
		}
		// ahandler.OtpService <- otpResponse
		er = ahandler.UserService.RegisterTempoCXP(ctx, tempo)
		if er != nil {
			if strings.Contains(er.Error(), "duplicate key value violates unique constraint") {
				resp.Msg = translation.Translate(session.Lang, "Agent with this information already registered")
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
		agent := &model.Agent{}
		agent.Firstname = input.Firstname
		agent.Lastname = input.Lastname
		agent.Phone = input.Phone
		agent.CreatedAt = uint64(time.Now().Unix())
		agent.RegisteredBy = uint(session.ID)
		agent.Password = hashed
		agent.FieldAddress = input.Address
		agent.Lang = session.Lang
		if agent.FieldAddress == nil {
			agent.FieldAddress = &model.Address{}
		}
		status, er := ahandler.Service.RegisterAgent(ctx, agent)
		if er != nil {
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.TranslateIt(er.Error())
		}

		if status == -1 {
			resp.StatusCode = http.StatusUnauthorized
			resp.Msg = translation.TranslateIt("you are not authorized to create this agent instance")
		} else if status == -2 {
			// missing aaddress information
			resp.StatusCode = http.StatusBadRequest
			resp.Msg = translation.Translate(session.Lang, "missing important address information, please complete necessary datas.")
		} else if status == -3 {
			resp.StatusCode = http.StatusInternalServerError
			resp.Msg = translation.Translate(session.Lang, "internal problem, please try again later")
		} else if status == -4 {
			resp.StatusCode = http.StatusConflict
			resp.Msg = translation.Translate(session.Lang, "agent with this information had already registered")

		} else if status > 0 {
			resp.Msg = translation.Translate(session.Lang, `You will recieve an SMS a message containing the confirmation code\nplease confirm your phone number with in 30 minutes.\n The Confirmation numbe also serve as your password`)
			resp.OTP = otpCode
			resp.Agent = agent
			resp.StatusCode = http.StatusOK
		}
		c.JSON(resp.StatusCode, resp)
	} else {
		resp.Msg = translation.TranslateIt("bad request body for agent creation")
		resp.StatusCode = http.StatusBadRequest
		c.JSON(resp.StatusCode, resp)
	}
}

// MerchantsSearch
func (ahandler *AgentHandler) AgentsSearch(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	response := &struct {
		StatusCode int            `json:"status_code"`
		Msg        string         `json:"msg"`
		Agents     []*model.Agent `json:"agents"`
	}{
		Agents: []*model.Agent{},
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
	name = strings.Trim(name, " ")
	agents, er := ahandler.Service.SearchAgents(ctx, phone, name, uint64(createdBy), uint(offset), uint(limit))
	if er != nil {
		response.StatusCode = http.StatusNotFound
		response.Msg = translation.Translate(session.Lang, "agents not found ")
		c.JSON(response.StatusCode, response)
		return
	}
	response.StatusCode = http.StatusOK
	response.Msg = translation.Translate(session.Lang, "agents found")
	response.Agents = agents
	c.JSON(response.StatusCode, response)
}

func (ahandler *AgentHandler) DeleteAgentByID(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)
	response := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg"`
	}{}
	agentid, er := strconv.Atoi(c.Param("id"))
	if er != nil || agentid <= 0 {
		response.Msg = translation.Translate(session.Lang, "invalid agent id")
		response.StatusCode = http.StatusBadRequest
		c.JSON(response.StatusCode, response)
		return
	}
	er = ahandler.Service.DeleteAgentByID(ctx, uint64(agentid))
	if er != nil {
		response.Msg = translation.Translate(session.Lang, "can't found an agent instance with this id")
		response.StatusCode = http.StatusNotFound
		c.JSON(response.StatusCode, response)
		return
	}
	response.Msg = translation.Translate(session.Lang, "agent deleted succesfuly")
	response.StatusCode = http.StatusOK
	c.JSON(response.StatusCode, response)
}
