package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/category"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/round"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
)

type IRoundHandler interface {
	// CreateRound ... create a round instance.
	CreateRound(c *gin.Context)
	UpdateRound(c *gin.Context)
	DeactivateRound(c *gin.Context)
	ActivateRound(c *gin.Context)
	DeleteRound(c *gin.Context)
	GetRoundsOfCategory(c *gin.Context)
	GetAllRounds(c *gin.Context)
	PopulateRoundStudents(c *gin.Context)
	SearchRoundByNumber(*gin.Context)
}

type RoundHandler struct {
	CategorySer category.ICategoryService
	RoundSer    round.IRoundService
}

// NewRoundHandler ... return a round handler Instance.
func NewRoundHandler(roundser round.IRoundService) IRoundHandler {
	return &RoundHandler{
		RoundSer: roundser,
	}
}

// CreateRound .. creates a new instance of round.
func (roundh *RoundHandler) CreateRound(c *gin.Context) {
	input := &struct {
		CategoryID   uint    `json:"category_id"`
		TrainingHour uint    `json:"training_hour,omitempty"`
		RoundNo      uint    `json:"round_no"` // category
		StartDate    string  `json:"start_date,omitempty"`
		EndDate      string  `json:"end_date,omitempty"`
		Lang         string  `json:"lang,omitempty"` // active_amount
		Fee          float64 `json:"fee,omitempty"`
		Active       bool    `json:"active"` // active
	}{}
	res := &struct {
		Msg   string       `json:"msg"`
		Round *model.Round `json:"round"`
	}{"bad request body", nil}
	// if active it must have training hour  , start date and end date set
	//
	if ere := c.BindJSON(input); ere != nil {
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.Active {
		if input.TrainingHour <= 0 || input.StartDate == "" || input.EndDate == "" {
			res.Msg = "invalid data\n Training Hour must be greater than 0 , start date and end date must be set for active trainings."
			c.JSON(http.StatusBadRequest, res)
			return
		}
	}
	if input.CategoryID == 0 || input.RoundNo == 0 {
		if input.CategoryID == 0 && input.RoundNo == 0 {
			res.Msg = " invalid category id and round id"
		} else if input.CategoryID == 0 {
			res.Msg = "invalid category id"
		} else {
			res.Msg = "invalid round id"
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", uint64(input.CategoryID))
	category, err, code := roundh.CategorySer.GetCategoryByID(ctx)
	if err != nil || category == nil || code != state.DT_STATUS_OK {
		res.Msg = "category with this id doesn't exist"
		c.JSON(http.StatusNotFound, res)
		return
	}
	// check the existance of a round having same round number in the same category.
	// if so , return an error message.
	ctx = context.WithValue(ctx, "round_number", uint64(input.RoundNo))
	round, statusCode, er := roundh.RoundSer.GetRoundByRoundNumberAndCategoryID(ctx)
	if (er == nil || statusCode != state.DT_STATUS_OK) && round != nil {
		res.Msg =
			"round with category ID : " +
				strconv.Itoa(int(input.CategoryID)) +
				" round number " +
				strconv.Itoa(int(input.RoundNo)) +
				" already exist"
		// ------------
		c.JSON(http.StatusConflict, res)
		return
	}
	// create the round instance from the input.
	round = &model.Round{
		CategoryID:   input.CategoryID,
		TrainingHour: input.TrainingHour,
		RoundNo:      input.RoundNo,
		Students:     0,
		ActiveAmount: 0,
		StartDate:    helper.ToValidDate(input.StartDate),
		EndDate:      helper.ToValidDate(input.EndDate),
		Lang:         input.Lang,
		Active:       input.Active,
		Fee: func() float64 {
			if input.Fee == 0.0 {
				return category.Fee
			}
			return input.Fee
		}(),
	}
	ctx = context.WithValue(ctx, "round", round)
	round, statusCode, er = roundh.RoundSer.CreateRound(ctx)
	if er != nil || round == nil {
		if er != nil {
			println(er.Error())
		}
		res.Msg = "internal server error"
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// Update the category this information is added on.
	category.RoundsCount++

	res.Round = round
	ctx = context.WithValue(ctx, "category", category)
	if ers, code := roundh.CategorySer.UpdateCategory(ctx); code != state.DT_STATUS_OK {
		ctx = context.WithValue(ctx, "round_id", round.ID)
		er = roundh.RoundSer.DeleteRoundByID(ctx)
		res.Msg = "internal server error"
		println(ers.Error())
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Msg = " round created succesfuly "
	c.JSON(http.StatusOK, res)
}

// UpdateRound  update a round instance.
func (roundh *RoundHandler) UpdateRound(c *gin.Context) {
	input := &model.RoundInput{}
	res := &struct {
		Msg   string       `json:"msg"`
		Round *model.Round `json:"round"`
	}{}
	if er := c.BindJSON(input); er != nil || input.ID <= 0 {
		res.Msg = "bad request body"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "round_id", uint64(input.ID))
	round, er := roundh.RoundSer.GetRoundByID(ctx)
	if round == nil || er != nil {
		res.Msg = "round with this id doesn't exist!"
		c.JSON(http.StatusNotFound, res)
		return
	}
	changed := false
	if input.TrainingHour != 0 && input.TrainingHour != round.TrainingHour {
		// Updating training hour
		round.TrainingHour = input.TrainingHour
		changed = true
	}
	if input.RoundNo > 0 && input.RoundNo != round.RoundNo {
		// round number change
		// check the existance of the round with same round number in same category with this round.
		ctx = context.WithValue(ctx, "round_number", uint64(input.RoundNo))
		ctx = context.WithValue(ctx, "category_id", uint64(round.CategoryID))
		if rd, _, er := roundh.RoundSer.GetRoundByRoundNumberAndCategoryID(ctx); er == nil && rd != nil {
			// there is alsready a reound with this category and round number
			res.Msg = " round number in this category already exist"
			c.JSON(http.StatusUnauthorized, res)
			return
		}
		round.RoundNo = input.RoundNo
		changed = true
	}
	input.StartDate = helper.ToValidDate(input.StartDate)
	input.EndDate = helper.ToValidDate(input.EndDate)

	if input.StartDate != "" &&
		round.StartDate != input.StartDate &&
		len(input.StartDate) >= 8 ||
		(func() bool {
			_, vv := helper.IsValidDate(input.StartDate)
			return vv
		}()) {
		round.StartDate = helper.ToValidDate(input.StartDate)
		changed = true
	} else if len(input.StartDate) < 8 &&
		input.StartDate != "" {
		res.Msg = "invalid start start date"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.EndDate != "" &&
		round.EndDate != input.EndDate &&
		len(input.EndDate) >= 8 ||
		(func() bool {
			_, vv := helper.IsValidDate(input.EndDate)
			return vv
		}()) {
		// ----------------------------------------------
		round.EndDate = helper.ToValidDate(input.EndDate)
		changed = true
		// ----------------------------------------------
	} else if len(input.EndDate) < 8 &&
		input.EndDate != "" {
		res.Msg = "invalid round end date"
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.Lang != "" && len(input.Lang) <= 4 && input.Lang != round.Lang {
		round.Lang = input.Lang
		changed = true
	}

	if input.Fee > 0 && input.Fee != round.Fee {
		changed = true
		round.Fee = input.Fee
	}
	if changed {
		ctx = context.WithValue(ctx, "round", round)
		if code, er := roundh.RoundSer.UpdateRound(ctx); er != nil {
			if code == state.DT_STATUS_RECORD_NOT_FOUND {
				res.Msg = " record not found "
				c.JSON(http.StatusNotFound, res)
			} else {
				res.Msg = "internal problem please try again "
				c.JSON(http.StatusInternalServerError, res)
			}
			return
		}
		res.Round = round
		res.Msg = "round updated succesfuly"
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotModified, nil)
	}
}

// This method is sused to deactivate a round ....
// deactivation is a simple process to make the activeroudn to false.
func (roundh *RoundHandler) DeactivateRound(c *gin.Context) {
	eres := &struct {
		Error string `json:"err"`
	}{
		"bad request",
	}
	ctx := c.Request.Context()
	sres := &struct {
		Msg string `json:"msg"`
	}{}
	roundID, er := strconv.Atoi(c.Query("id"))
	if er != nil || roundID <= 0 {
		eres.Error = "bad request"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx = context.WithValue(ctx, "round_id", uint64(roundID))
	round, er := roundh.RoundSer.GetRoundByID(ctx)
	if round == nil || er != nil {
		eres.Error = "round not found"
		c.JSON(http.StatusNotFound, eres)
		return
	}
	if !(round.Active) {
		sres.Msg = "round is already deactivated"
		c.JSON(http.StatusOK, sres)
		return
	}
	round.Active = false
	ctx = context.WithValue(ctx, "round", round)
	if scode, er := roundh.RoundSer.UpdateRound(ctx); er != nil {
		if scode == state.DT_STATUS_NO_RECORD_UPDATED {
			eres.Error = " deactivation was not succesful"
			c.JSON(http.StatusNotModified, eres)
		} else {
			eres.Error = " internal problem "
			c.JSON(http.StatusInternalServerError, eres)
		}
		return
	}
	sres.Msg = " round deactivated succesfuly"
	c.JSON(http.StatusOK, sres)
}

// ActivateRound to make already previously
// active round but bit deactive round to active.
func (roundh *RoundHandler) ActivateRound(c *gin.Context) {
	eres := &struct {
		Error string `json:"err"`
	}{
		"bad request",
	}
	ctx := c.Request.Context()
	sres := &struct {
		Msg string `json:"msg"`
	}{}
	roundID, er := strconv.Atoi(c.Query("id"))
	if er != nil || roundID <= 0 {
		eres.Error = "bad request"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx = context.WithValue(ctx, "round_id", uint64(roundID))
	round, er := roundh.RoundSer.GetRoundByID(ctx)
	if er != nil || round == nil {
		eres.Error = "round not found"
		c.JSON(http.StatusNotFound, eres)
		return
	}
	if round.Active {
		sres.Msg = "round is already activated"
		c.JSON(http.StatusOK, sres)
		return
	}
	round.Active = true
	ctx = context.WithValue(ctx, "round", round)
	if scode, er := roundh.RoundSer.UpdateRound(ctx); er != nil {
		if scode == state.DT_STATUS_NO_RECORD_UPDATED {
			eres.Error = " activation was not succesful"
			c.JSON(http.StatusNotModified, eres)
		} else {
			eres.Error = " internal problem "
			c.JSON(http.StatusInternalServerError, eres)
		}
		return
	}
	sres.Msg = " round activated succesfuly"
	c.JSON(http.StatusOK, sres)
}

// DeleteRound delete a round include deleting all the students referencing that round so, i have to create the student instances before.
func (roundh *RoundHandler) DeleteRound(c *gin.Context) {

}

// GetRoundsOfCategory -- this method gets all the rounds of a single category instance.
//
func (roundh *RoundHandler) GetRoundsOfCategory(c *gin.Context) {
	eres := &model.ErMsg{}
	categoryID, er := strconv.Atoi(c.Request.FormValue("category_id"))
	if er != nil {
		eres.Error = "bad category id"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "category_id", uint(categoryID))
	rounds, scode, err := roundh.RoundSer.GetRoundsOfCategory(ctx)
	if err != nil {
		if scode == state.DT_STATUS_NO_RECORD_FOUND {
			c.JSON(http.StatusNotFound, rounds)
		} else {
			c.JSON(http.StatusInternalServerError, rounds)
		}
		return
	}
	c.JSON(http.StatusOK, rounds)
}

func (roundh *RoundHandler) GetAllRounds(c *gin.Context) {

}

func (roundh *RoundHandler) PopulateRoundStudents(c *gin.Context) {

}

func (roundh *RoundHandler) SearchRoundByNumber(c *gin.Context) {

}

func (roundh *RoundHandler) StudentsOfARound(c *gin.Context) {}
