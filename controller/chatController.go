package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatController interface {
	SendChat(ctx echo.Context) error
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendChat(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		broadcast(msg)
	}

	return nil
}

var clients = make(map[*websocket.Conn]bool)
var broadcastChannel = make(chan []byte)

func broadcast(message []byte) {
	broadcastChannel <- message
}

func handleMessages() {
	for {
		message := <-broadcastChannel

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
