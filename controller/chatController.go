package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcasts   = make(map[string]*Broadcast)
	broadcastsMu sync.Mutex
)

type Broadcast struct {
	ID       string
	Clients  map[*websocket.Conn]bool
	Username string
}

type ChatController interface {
	HandleWebsocket(ctx echo.Context) error
}

type ChatControllerImpl struct {
	UserService service.UserService
}

func NewChatController(UserService service.UserService) *ChatControllerImpl {
	return &ChatControllerImpl{UserService: UserService}
}

func (c *ChatControllerImpl) HandleWebsocket(ctx echo.Context) error {
	id := ctx.Param("id")

	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	result, err := c.UserService.FindById(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return conn.WriteJSON(helper.ErrorResponse("Users Not Found"))
		}
		return conn.WriteJSON(helper.ErrorResponse("Get User Data By Id Error"))
	}

	broadcastID := result.ID
	joinBroadcast(conn, broadcastID, result.Username)

	go onConnect(conn, broadcastID, result.Username)
	go onMessage(conn, broadcastID, result.Username)
	go onClose(conn, broadcastID)
	go receiveMessages(conn, broadcastID)

	return nil
}

func createBroadcast(id, username string) {
	broadcast := &Broadcast{
		ID:       id,
		Clients:  make(map[*websocket.Conn]bool),
		Username: username,
	}
	broadcasts[id] = broadcast
}

func joinBroadcast(conn *websocket.Conn, id, username string) {
	broadcastsMu.Lock()
	defer broadcastsMu.Unlock()

	broadcast, ok := broadcasts[id]
	if !ok {
		createBroadcast(id, username)
		broadcast = broadcasts[id]
	}

	broadcast.Clients[conn] = true
}

func onConnect(conn *websocket.Conn, id, username string) {
	conn.WriteJSON(map[string]string{
		"message": fmt.Sprintf("Welcome, %s, to %s broadcast", username, id),
	})
}

func onMessage(conn *websocket.Conn, id, username string) {
	for {
		var msg web.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		sendMessage(conn, id, username, msg)
	}
}

func onClose(conn *websocket.Conn, id string) {
	broadcastsMu.Lock()
	defer broadcastsMu.Unlock()

	leaveBroadcast(conn, id)
	conn.Close()
}

func sendMessage(conn *websocket.Conn, id, username string, msg web.Message) {
	broadcastsMu.Lock()
	defer broadcastsMu.Unlock()

	broadcast, ok := broadcasts[id]
	if !ok {
		return
	}

	for client := range broadcast.Clients {
		if client != conn {
			msg.Username = username
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func receiveMessages(conn *websocket.Conn, id string) {
	for {
		var msg web.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		handleReceivedMessage(id, msg)
	}
}

func handleReceivedMessage(broadcastID string, receivedMsg web.Message) {
	broadcastsMu.Lock()
	defer broadcastsMu.Unlock()

	broadcast, ok := broadcasts[broadcastID]
	if !ok {
		return
	}

	for client := range broadcast.Clients {
		err := client.WriteJSON(receivedMsg)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	fmt.Printf("Received message from %s\n", receivedMsg.Content)
}

func leaveBroadcast(conn *websocket.Conn, id string) {
	broadcastsMu.Lock()
	defer broadcastsMu.Unlock()

	broadcast, ok := broadcasts[id]
	if !ok {
		return
	}

	delete(broadcast.Clients, conn)
}
