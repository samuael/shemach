package message_broadcast_service

import (
	"context"
	"encoding/json"
	"time"

	tm "github.com/buger/goterm"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
)

type MainBroadcastHub struct {
	Register         chan *Client
	UnRegister       chan *Client
	BroadcastMessage chan *model.Message
	Clients          map[string]*Client
	// UserService      user.IUserService
	//  ----------------------------------
	MessageService message.IMessageService
	ProductUpdate  chan *model.ProductUpdate
}

func NewMainBroadcastHub(
	messageService message.IMessageService, // service user.IUserService,
) *MainBroadcastHub {
	return &MainBroadcastHub{
		Register:         make(chan *Client),
		UnRegister:       make(chan *Client),
		BroadcastMessage: make(chan *model.Message),
		Clients:          map[string]*Client{},
		ProductUpdate:    make(chan *model.ProductUpdate),
		MessageService:   messageService,
	}
}

func (mainbh *MainBroadcastHub) Run() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case client := <-mainbh.Register:
			{
				mainbh.Clients[client.Phone] = client
			}
		case client := <-mainbh.UnRegister:
			{
				delete(mainbh.Clients, client.Phone)
			}
		case message := <-mainbh.BroadcastMessage:
			{
				ctx := context.TODO()
				ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*2))
				ctx = context.WithValue(ctx, "message", message)
				// ---------------------------------------------------------
				message, status, er := mainbh.MessageService.SaveMessage(ctx)
				if (status != state.STATUS_OK) || (er != nil) {
				}
				mainbh.BroadcastNewMessage(message)
			}
		case update := <-mainbh.ProductUpdate:
			{
				mainbh.BroadcastProductUpdate(update)
			}
		case <-ticker.C:
			{
			}
		}
	}
}
func (mainbh *MainBroadcastHub) BroadcastProductUpdate(update *model.ProductUpdate) {
	data, er := json.Marshal(map[string]interface{}{"type": 2, "body": update})
	if er != nil || data == nil {
		return
	}
	message := &BinaryMessage{
		Targets: []int{},
		Data:    data,
	}

	for _, client := range mainbh.Clients {
		if client.Role != state.SUBSCRIBER {
			client.Message <- message
		}
		for _, csubsctiption := range client.Subscriptions {
			if csubsctiption == -1 {
				client.Message <- message
			}
			if int(update.ID) == csubsctiption {
				client.Message <- message
				break
			}
		}
	}
}

func (mainbh *MainBroadcastHub) BroadcastNewMessage(mes *model.Message) {
	data, er := json.Marshal(map[string]interface{}{"type": 1, "body": mes})
	tm.Println(tm.Bold(tm.Color(string(data), tm.RED)))
	tm.Clear()
	// tm.Println(tm.Bold(tm.Color(string("This Is Samuael adnew"), tm.RED)))
	tm.Flush()
	if er != nil || data == nil {
		return
	}
	message := &BinaryMessage{
		Targets: mes.Targets,
		Lang:    mes.Lang,
		Data:    data,
	}
	passed := true
	for _, client := range mainbh.Clients {
		if !(message.Lang == state.LANGUAGE_ALL || client.Lang == message.Lang) {
			continue
		}
		if (client.Role == state.SUPERADMIN || client.Role == state.INFO_ADMIN) && client.ID == mes.ID {
			client.Message <- message
			continue
		}
		if len(message.Targets) > 0 && message.Targets[0] == -1 {
			client.Message <- message
		}
		for _, subscription := range message.Targets {
			for _, csubsctiption := range client.Subscriptions {
				if subscription == csubsctiption {
					client.Message <- message
					passed = true
					break
				}
			}
			if passed {
				break
			}
		}
	}
}
