package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/cmd/main/service/message_broadcast_service"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IMessageHandler interface {
	GetRecentMessages(c *gin.Context)
	SendMessage(c *gin.Context)
	DeleteMessageByID(c *gin.Context)
	GetAllMessages(c *gin.Context)
}

type MessageHandler struct {
	Service           message.IMessageService
	SubscriberService subscriber.ISubscriberService
	BroadcastHub      *message_broadcast_service.MainBroadcastHub
}

func NewMessageHandler(service message.IMessageService, ser subscriber.ISubscriberService,
	broadcastHub *message_broadcast_service.MainBroadcastHub,
) IMessageHandler {
	return &MessageHandler{
		Service:           service,
		SubscriberService: ser,
		BroadcastHub:      broadcastHub,
	}
}

func (mhan *MessageHandler) GetRecentMessages(c *gin.Context) {
	ctx := c.Request.Context()
	res := &struct {
		Messages   []*model.Message `json:"messages"`
		Msg        string           `json:"msg,omitempty"`
		StatusCode int              `json:"status_code"`
	}{}
	subscriberSession := ctx.Value("session").(*model.SubscriberSession)
	// println(subscriberSession.ID, subscriberSession.Lang, subscriberSession.Phone, subscriberSession.Fullname)
	unix, era := strconv.Atoi(c.Query("unix"))
	if era != nil {
		unix = 0
	}
	offset, erb := strconv.Atoi(c.Query("offset"))
	if erb != nil {
		offset = 0
	}
	limit, erc := strconv.Atoi(c.Query("limit"))
	if erc != nil {
		limit = offset + 5
	}
	ctx = context.WithValue(ctx, "subscriber_id", subscriberSession.ID)
	subscriber, status, era := mhan.SubscriberService.GetSubscriberByID(ctx)

	if status != state.STATUS_OK || era != nil || subscriber == nil {
		if era != nil {
			println(era.Error())
		}
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("no record found")
		c.JSON(res.StatusCode, res)
		return
	}
	psubscriptions := []int{}
	for _, sub := range subscriber.Subscriptions {
		psubscriptions = append(psubscriptions, int(sub))
	}
	ctx = context.WithValue(ctx, "offset", uint(offset))
	ctx = context.WithValue(ctx, "limit", uint(limit))
	ctx = context.WithValue(ctx, "unix_time", uint64(unix))
	ctx = context.WithValue(ctx, "subscriptions", psubscriptions)
	ctx = context.WithValue(ctx, "lang", subscriber.Lang)
	messages, status, er := mhan.Service.GetRecentMessages(ctx)
	if er != nil || status != state.STATUS_OK {
		res.StatusCode = http.StatusNotFound
		res.Msg = translation.TranslateIt("record not found")
		c.JSON(res.StatusCode, res)
		return
	}
	res.Messages = messages
	println(helper.MarshalThis(messages))
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}

// SendMessage ...
func (mhan *MessageHandler) SendMessage(c *gin.Context) {
	ctx := c.Request.Context()
	resp := &struct {
		StatusCode int            `json:"status_code"`
		Msg        string         `json:"msg"`
		Message    *model.Message `json:"message"`
	}{}
	session := ctx.Value("session").(*model.Session)

	input := &model.Message{}
	jsonDecoder := json.NewDecoder(c.Request.Body)
	er := jsonDecoder.Decode(input)
	if er != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Msg = translation.Translate(session.Lang, "bad message request")
		c.JSON(resp.StatusCode, resp)
		return
	}
	ctx = context.WithValue(ctx, "message", input)
	input, _, er = mhan.Service.SaveMessage(ctx)
	if input == nil || er != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Msg = translation.Translate(session.Lang, "internal server error ")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.StatusCode = http.StatusOK
	input.CreatedBy = uint(session.ID)
	input.CreatedAt = uint64(time.Now().Unix())
	resp.Message = input
	resp.Msg = translation.Translate(session.Lang, "created succesfuly")
	mhan.BroadcastHub.BroadcastMessage <- input
	c.JSON(resp.StatusCode, resp)
}

func (mhan *MessageHandler) GetAllMessages(c *gin.Context) {
	ctx := c.Request.Context()
	resp := &struct {
		StatusCode int              `json:"status_code"`
		Msg        string           `json:"msg"`
		Messages   []*model.Message `json:"message"`
	}{}
	session := ctx.Value("session").(*model.Session)
	offset, er := strconv.ParseInt(c.Query("offset"), 10, 64)
	if offset <= 0 || er != nil {
		offset = 0
	}
	limit, ers := strconv.Atoi(c.Query("limit"))
	if ers != nil || (limit <= int(offset)) {
		limit = (int(offset) + 10)
	}
	messages, er := mhan.Service.GetMessages(ctx, int(offset), int(limit))
	if er != nil {
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "no message found")
		resp.Messages = []*model.Message{}
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.Messages = messages
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "messages found")
	c.JSON(resp.StatusCode, resp)
}

func (mhan MessageHandler) DeleteMessageByID(c *gin.Context) {
	ctx := c.Request.Context()
	resp := &struct {
		StatusCode int    `json:"status_code"`
		Msg        string `json:"msg"`
	}{}
	session := ctx.Value("session").(*model.Session)
	messageid, ers := strconv.Atoi(c.Param("id"))
	if ers != nil || messageid <= 0 {
		resp.StatusCode = http.StatusBadRequest
		resp.Msg = translation.Translate(session.Lang, "bad message id")
		c.JSON(resp.StatusCode, resp)
		return
	}
	er := mhan.Service.DeleteMessageBYID(ctx, uint(messageid))
	if er != nil {
		println(er.Error())
		resp.StatusCode = http.StatusNotFound
		resp.Msg = translation.Translate(session.Lang, "message instance not found")
		c.JSON(resp.StatusCode, resp)
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Msg = translation.Translate(session.Lang, "message instance deleted")
	c.JSON(resp.StatusCode, resp)
}
