package message_broadcast_service

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
)

type IClientConnetionHandler interface {
	SubscriberHandleWebsocketConnection(c *gin.Context)
	AdminsHandleWebsocketConnection(c *gin.Context)
}

type ClientConnetionHandler struct {
	SubscriberService subscriber.ISubscriberService
	UserService       user.IUserService
	TheHub            *MainBroadcastHub
}

// NewClientConnectionHandler ...
func NewClientConnectionHandler(
	subscriberService subscriber.ISubscriberService,
	userservice user.IUserService,
	hub *MainBroadcastHub,
) IClientConnetionHandler {
	return &ClientConnetionHandler{
		TheHub:            hub,
		SubscriberService: subscriberService,
		UserService:       userservice,
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
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil || id <= 0 {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "subscriber_id", uint64(id))
	subscriber, status, er := cch.SubscriberService.GetSubscriberByID(ctx)
	connection, erra := upgrader.Upgrade(response, request, nil)
	if connection == nil || erra != nil {
		return
	}

	session := &model.SubscriberSession{
		ID:       subscriber.ID,
		Fullname: subscriber.Fullname,
		Phone:    subscriber.Phone,
		Lang:     subscriber.Lang,
	}
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
	id, er := strconv.Atoi(c.Param("id"))
	if er != nil || id <= 0 {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "user_id", uint64(id))
	user, role, _, er := cch.UserService.GetUserByEmailOrID(ctx)
	if user == nil || role <= 0 || role >= 5 || er != nil {
		if er != nil {
			println(er.Error())
		}
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}
	session := &model.Session{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Lang:      user.Lang,
	}
	// var duser interface{}
	if role == 1 {
		session.Role = state.SUPERADMIN
	} else if role == 2 {
		session.Role = state.INFO_ADMIN
	} else if role == 3 {
		session.Role = state.ADMIN
	} else if role == 4 {
		session.Role = state.MERCHANT
	} else if role == 5 {
		session.Role = state.AGENT
	}
	connection, erra := upgrader.Upgrade(response, request, nil)
	if connection == nil || erra != nil {
		return
	}
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
