package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
)

type IMessageHandler interface {
}

type MessageHandler struct {
	Service message.IMessageService
}

func NewMessageHandler(service message.IMessageService) IMessageHandler {
	return &MessageHandler{
		Service: service,
	}
}

func (mhan *MessageHandler) GetRecentMessages(c *gin.Context) {
	// offset := strconv.Atoi(c.Query("offset"))

	// ctx := c.Request.Context()
	// ctx = context.WithValue(ctx, "message", &model.Message{})
	// messages, status, er := mhan.Service.SaveMessage(ctx)

	// state.STATUS_OK

}
