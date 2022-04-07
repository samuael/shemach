package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/payment"
	"github.com/samuael/agri-net/agri-net-backend/platforms"
)

type IPaymentHandler interface {
	GetDailyPaymentReport(c *gin.Context)
	GetMonthlyPaymentReport(c *gin.Context)
	Payin(c *gin.Context)
	PayoutRequest(c *gin.Context)
	ApprovePayIn(c *gin.Context)
	ApprovePayout(c *gin.Context)
	GetUnseenDailyPayments(c *gin.Context)
	GetSeenDailyPayments(c *gin.Context)
	GetApprovedDailyPayments(c *gin.Context)

	GetStudentsPayments(c *gin.Context)
	GetStudentsPayout(c *gin.Context)
	DeleteStudentPayment(c *gin.Context)
	DeletePayment(c *gin.Context)
}
type PaymentHandler struct {
	Service payment.IPaymentService
}

func NewPaymentHandler(ser payment.IPaymentService) IPaymentHandler {
	return &PaymentHandler{Service: ser}
}

func (phan *PaymentHandler) GetDailyPaymentReport(c *gin.Context)   {}
func (phan *PaymentHandler) GetMonthlyPaymentReport(c *gin.Context) {}

// Payin
func (phan *PaymentHandler) Payin(c *gin.Context) {
	ctx := c.Request.Context()
	println(ctx)
	in := &model.PayinInput{}
	res := &struct {
		Errors     map[string]string `json:"errors"`
		Msg        string            `json:"msg"`
		StatusCode int               `json:"status_code"`
		Payment    *model.PayIn      `json:"payment"`
	}{
		Errors: map[string]string{},
	}
	jdecode := json.NewDecoder(c.Request.Body)
	er := jdecode.Decode(in)
	if er != nil {
		res.Msg = "invalid bad request body"
		res.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}
	in.CreatedAt = platforms.UnixToEthiopianDate(int(in.UnixTime))
	// check the validity of the data types I have received above.
	var failes = false
	if in.Amount <= 0 {
		failes = true
		res.Errors["Amount"] = "Unacceptable payment amount"
	}
	if in.PayedBy <= 0 {
		failes = true
		res.Errors["Student"] = "student with this id doesn't exist"
	}
	if in.RoundID <= 0 {
		res.Errors["Round"] = "round with this id doesn't exist"
	}
	if in.CreatedAt.Years <= 0 || in.CreatedAt.Years < (platforms.GetCurrentEthiopianTime().Years-1) || in.CreatedAt.Years > uint(platforms.GetCurrentEthiopianTime().Years) {
		failes = true
		res.Errors["TimeStamp"] = "invalid timestamp information"
	} else if in.CreatedAt.Months >= (platforms.GetCurrentEthiopianTime().Months + 1) {
		failes = true
		res.Errors["TimeStamp"] = "invalid time stamp information"
	} else if in.CreatedAt.Days <= 0 || in.CreatedAt.Hours <= 0 || in.CreatedAt.Minutes <= 0 || in.CreatedAt.Seconds <= 0 {
		failes = true
		res.Errors["TimeStamp"] = "invalid timestamp information"
	}
	if in.UChars == "" || len(in.UChars) != 2 {
		failes = true
		res.Errors["unique character"] = "invalid unique characters\n It must have at least a length of 2"
	}
	if failes {
		res.StatusCode = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, res)
		return
	}
	in.ReceivedBy = int(ctx.Value("session").(*model.Session).ID)
	ctx = context.WithValue(ctx, "payin_input", in)
	status := phan.Service.CheckTheExistanceOfPaymentInstance(ctx)
	if status == -3 {
		res.Msg = "Internal issue occured, Please try again"
		res.StatusCode = http.StatusInternalServerError
		c.JSON(res.StatusCode, res)
	} else if status == -2 {
		res.Msg = "Round with this ID doesn't exist"
		res.StatusCode = http.StatusNotFound
		res.Errors["Round"] = "round with this information doesn't exist"
		c.JSON(res.StatusCode, res)
	} else if status == -1 {
		res.Msg = "Student with this information doesn't exist"
		res.Errors["Student"] = "student with this information doesn't exist"
		res.StatusCode = http.StatusNotFound
		c.JSON(res.StatusCode, res)
	} else if status == 1 {
		res.Msg = "Transaction with this information already exist"
		ctx = context.WithValue(ctx, "payin_input", in)
		payin, status, er := phan.Service.GetPaymentUsingPayinInput(ctx)
		if status != state.DT_STATUS_OK || er != nil {
			log.Println("Internal Server Problem while fetching the status")
		}
		res.Payment = payin
		res.StatusCode = http.StatusConflict
		c.JSON(res.StatusCode, res)
	}
	ctx = context.WithValue(ctx, "student_id", uint64(in.PayedBy))
	remainingAmount, status, er := phan.Service.GetRemainingPaymentOfStudentForRound(ctx)
	if er != nil || remainingAmount == -1 || remainingAmount == -2 {
		if remainingAmount == -1 {
			res.Msg = "Student with this information doesn't exist"
			res.Errors["Student"] = "student with this information doesn't exist"
			res.StatusCode = http.StatusNotFound
			c.JSON(res.StatusCode, res)
			return
		}
		if remainingAmount == -2 {
			res.Msg = "Round with this ID doesn't exist"
			res.StatusCode = http.StatusNotFound
			res.Errors["Round"] = "round with this information doesn't exist"
			c.JSON(res.StatusCode, res)
			return
		}
		if er != nil {
			res.Msg = "internal problem"
			res.StatusCode = http.StatusInternalServerError
			c.JSON(res.StatusCode, res)
			return
		}
		if remainingAmount == 0 {
			res.Msg = "student had completed the payment for the round"
			res.StatusCode = http.StatusNotImplemented
			c.JSON(res.StatusCode, res)
			return
		}
	}
	if remainingAmount < in.Amount {
		res.Msg = "payment amount is more than needed, remainint payment needed is " + strconv.Itoa(int(remainingAmount))
		res.StatusCode = http.StatusUnauthorized
		c.JSON(res.StatusCode, res)
		return
	}
	ctx = context.WithValue(ctx, "payin_input", in)
	payment, status, _ := phan.Service.CreatePayinInstance(ctx)
	if status == state.DT_STATUS_DBQUERY_ERROR {
		res.StatusCode = http.StatusInternalServerError
		res.Msg = "internal problem please try again"
		c.JSON(res.StatusCode, res)
		return
	} else if status == state.DT_STATUS_INCOMPLETE_DATA {
		res.StatusCode = http.StatusFailedDependency
		res.Msg = "transaction creation failed"
		c.JSON(res.StatusCode, res)
		return
	}
	res.Payment = payment
	res.StatusCode = http.StatusCreated
	c.JSON(http.StatusCreated, res)
}
func (phan *PaymentHandler) PayoutRequest(c *gin.Context)            {}
func (phan *PaymentHandler) ApprovePayIn(c *gin.Context)             {}
func (phan *PaymentHandler) ApprovePayout(c *gin.Context)            {}
func (phan *PaymentHandler) GetUnseenDailyPayments(c *gin.Context)   {}
func (phan *PaymentHandler) GetSeenDailyPayments(c *gin.Context)     {}
func (phan *PaymentHandler) GetApprovedDailyPayments(c *gin.Context) {}

func (phan *PaymentHandler) GetStudentsPayments(c *gin.Context) {
	ctx := c.Request.Context()
	eres := &struct {
		Errors     map[string]string `json:"errors"`
		StatusCode int               `json:"status_code"`
		Msg        string            `json:"msg"`
	}{}
	studentID, er := strconv.Atoi(c.Query("student_id"))
	if er != nil || studentID <= 0 {
		eres.Errors["Student ID"] = "missing a valid student ID"
		eres.Msg = "missing a valid student ID"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx = context.WithValue(ctx, "student_id", uint64(studentID))
	payments, status, er := phan.Service.GetPaymentsOfAStudent(ctx)
	if status == state.DT_STATUS_NO_RECORD_FOUND {
		eres.Msg = "no payments instances were created by this student."
		eres.StatusCode = http.StatusNotFound
		c.JSON(http.StatusNotFound, eres)
		return
	} else if !(status == state.DT_STATUS_OK) {
		eres.Msg = "internal problem please try again"
		eres.StatusCode = http.StatusInternalServerError
		c.JSON(eres.StatusCode, eres)
		return
	}
	c.JSON(http.StatusOK, payments)

}
func (phan *PaymentHandler) GetStudentsPayout(c *gin.Context)    {}
func (phan *PaymentHandler) DeleteStudentPayment(c *gin.Context) {}
func (phan *PaymentHandler) DeletePayment(c *gin.Context)        {}
