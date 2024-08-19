package message_broadcast_service

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/message"
)

type Client struct {
	ID            uint64
	Role          string
	Firstname     string
	Lastname      string
	Lang          string
	Phone         string
	Email         string
	Subscriptions []int
	Conn          *websocket.Conn
	TheHub        *MainBroadcastHub
	Message       chan *BinaryMessage
	// -----------------------------------
	MessageService message.IMessageService
}

const (
	writeWait       = 10 * time.Second
	pongWait        = 60 * time.Second
	pmessagegPeriod = (pongWait * 9) / 10
	maxMessageSize  = 99999999999
)

func (client *Client) WriteMessage() {
	ticker := time.NewTicker(pongWait)
	defer func() {
		defer func() {
			recover()
			client.TheHub.UnRegister <- client
		}()
		ticker.Stop()
		client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.Message:
			{
				if !ok {
					client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}
				client.Conn.WriteMessage(websocket.TextMessage, append(message.Data, '\n'))
				time.Sleep(time.Millisecond * 100)
			}
		case <-ticker.C:
			{
				if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}
}

func (client *Client) ReadMessage() {
	defer func() {
		defer func() {
			recover()
			client.TheHub.UnRegister <- client
		}()
		client.Conn.Close()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	client.Conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})
	client.Conn.SetPongHandler(func(string) error {
		/*client.Conns[key].Conn.SetReadDeadline(time.Now().Add(pongWait)); */
		return nil
	})
	for {
		message := &model.Message{}
		err := client.Conn.ReadJSON(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				return
			}
			break
		}
		if message == nil {
			continue
		}
		// Brodcast the message
		//
		message.CreatedBy = uint(client.ID)
		client.TheHub.BroadcastMessage <- message
	}
}
