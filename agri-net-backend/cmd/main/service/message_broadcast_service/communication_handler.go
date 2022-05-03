package message_broadcast_service

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
)

type IClientConnetionHandler interface {
	SubscriberHandleWebsocketConnection(c *gin.Context)
	AdminsHandleWebsocketConnection(c *gin.Context)
}

type ClientConnetionHandler struct {
	SubscriberService subscriber.ISubscriberService
	TheHub            *MainBroadcastHub
}

// NewClientConnectionHandler ...
func NewClientConnectionHandler(
	subscriberService subscriber.ISubscriberService,
	hub *MainBroadcastHub,
) IClientConnetionHandler {
	return &ClientConnetionHandler{
		TheHub:            hub,
		SubscriberService: subscriberService,
		// MessageService:    messageservice,
	}
}

var upgrader = websocket.Upgrader{
	WriteBufferSize: 2048,
	ReadBufferSize:  2048,
}

// SubscriberHandleWebsocketConnection(c *gin.Context) ...
func (cch *ClientConnetionHandler) SubscriberHandleWebsocketConnection(c *gin.Context) {
	response := c.Writer
	request := c.Request
	connection, erra := upgrader.Upgrade(response, request, nil)
	if connection == nil || erra != nil {
		return
	}
	ctx := request.Context()
	session := ctx.Value("session").(*model.SubscriberSession)
	// ---------------
	ctx = context.WithValue(ctx, "subscriber_phone", session.Phone)
	subscriber, status, er := cch.SubscriberService.GetSubscriberByPhone(ctx)
	if status != state.STATUS_OK || er != nil || subscriber == nil {
		return
	}
	isubscriptions := []int{}
	for _, sub := range subscriber.Subscriptions {
		isubscriptions = append(isubscriptions, int(sub))
	}
	client := &Client{
		ID:            session.ID,
		Conn:          connection,
		Role:          state.SUBSCRIBER,
		Firstname:     strings.Split(session.Fullname, " ")[0],
		Lastname:      strings.Split(session.Fullname, " ")[1],
		Lang:          session.Lang,
		Phone:         session.Phone,
		Subscriptions: isubscriptions,
		Message:       make(chan *BinaryMessage),
		TheHub:        cch.TheHub,
	}
	cch.TheHub.Register <- client
	go client.WriteMessage()
	message := map[string]interface{}{
		"type": 1,
		"body": &model.Message{
			ID:        0,
			Targets:   []int{-1},
			Lang:      "all",
			Data:      "Welcome to Agri-net systems",
			CreatedAt: uint64(time.Now().Unix()),
		}}
	data, _ := json.Marshal(message)
	client.Message <- &BinaryMessage{
		Data: data,
	}
}

func (cch *ClientConnetionHandler) AdminsHandleWebsocketConnection(c *gin.Context) {
	response := c.Writer
	request := c.Request
	connection, erra := upgrader.Upgrade(response, request, nil)
	if connection == nil || erra != nil {
		return
	}
	ctx := request.Context()
	session := ctx.Value("session").(*model.Session)
	client := &Client{
		ID:            session.ID,
		Conn:          connection,
		Role:          session.Role,
		Firstname:     session.Firstname,
		Lastname:      session.Lastname,
		Lang:          session.Lang,
		Email:         session.Email,
		Phone:         session.Email,
		Message:       make(chan *BinaryMessage),
		Subscriptions: []int{0},
		TheHub:        cch.TheHub,
	}
	cch.TheHub.Register <- client
	go client.ReadMessage()
	go client.WriteMessage()
	message := map[string]interface{}{
		"type": 1,
		"body": &model.Message{
			ID:        0,
			Targets:   []int{-1},
			Lang:      "all",
			Data:      "Welcome to Agri-net systems",
			CreatedAt: uint64(time.Now().Unix()),
		},
	}
	data, _ := json.Marshal(message)
	client.Message <- &BinaryMessage{
		Data: data,
	}
}
