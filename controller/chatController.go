package controller

import (
	"bytes"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	room string
}

type Hub struct {
	rooms      map[string]*Room
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

type Room struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	client *Client
	data   []byte
}

func newHub() *Hub {
	return &Hub{
		rooms:      make(map[string]*Room),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			room, ok := h.rooms[client.room]
			if !ok {
				room = newRoom()
				h.rooms[client.room] = room
				go room.run()
			}
			h.mu.Unlock()

			room.register <- client

		case client := <-h.unregister:
			h.mu.Lock()
			if room, ok := h.rooms[client.room]; ok {
				room.unregister <- client
				if len(room.clients) == 0 {
					delete(h.rooms, client.room)
				}
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			if room, ok := h.rooms[message.client.room]; ok {
				room.broadcast <- message
			}
			h.mu.Unlock()
		}
	}
}

func newRoom() *Room {
	return &Room{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- message.data:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- &Message{client: c, data: message}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveWs(hub *Hub, c echo.Context, room, userName string) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return err
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), room: room}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	message := []byte(userName + " has joined the room.")
	hub.broadcast <- &Message{client: client, data: message}

	log.Println(userName, "connected to room:", room)

	return nil
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

	user, err := c.UserService.FindById(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("User Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Id Error"))
	}

	chatHub := newHub()
	go chatHub.run()

	room := "room_" + id

	return serveWs(chatHub, ctx, room, user.Name)
}
