package websocket

import (
	// "fmt"
	"ecomplaint/config"
	"ecomplaint/model/schema"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

func (c *Client) WriteMessage(roomID string) {
	defer func() {
		c.Conn.Close()
	}()
	var messages []schema.Message
	config.DB.Where("room_id = ?", roomID).Find(&messages)
	c.Conn.WriteJSON(messages)
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()
	//extract Token sender
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		message := schema.Message{
			RoomID:  msg.RoomID,
			Message: msg.Content,
		}
		fmt.Println(msg)

		config.DB.Create(&message)

		hub.Broadcast <- msg
	}
}
