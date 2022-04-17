package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/platforms/translation"
)

type IMessageHandler interface {
	GetRecentMessages(c *gin.Context)
}

type MessageHandler struct {
	Service           message.IMessageService
	SubscriberService subscriber.ISubscriberService
}

func NewMessageHandler(service message.IMessageService, ser subscriber.ISubscriberService) IMessageHandler {
	return &MessageHandler{
		Service:           service,
		SubscriberService: ser,
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
	println(subscriberSession.ID, subscriberSession.Lang, subscriberSession.Phone, subscriberSession.Fullname)
	unix, era := strconv.Atoi(c.Query("unix"))
	if era != nil {
		unix = 0
	}
	offset, erb := strconv.Atoi(c.Query("unix"))
	if erb != nil {
		offset = 0
	}
	limit, erc := strconv.Atoi(c.Query("unix"))
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
	res.StatusCode = http.StatusOK
	c.JSON(res.StatusCode, res)
}
