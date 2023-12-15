package controller

import (
	"ecomplaint/config"
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	ws "ecomplaint/utils/helper/websocket"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatController interface {
	CreateRoom(ctx echo.Context) error
	JoinRoom(c echo.Context) error
	GetRooms(c echo.Context) error
	GetClients(c echo.Context) error
	GetChats(c echo.Context) error
}

type ChatControllerImpl struct {
	UserService  service.UserService
	AdminService service.AdminService
	hub          *ws.Hub
}

func NewChatController(UserService service.UserService, AdminService service.AdminService, hub *ws.Hub) *ChatControllerImpl {
	return &ChatControllerImpl{
		UserService:  UserService,
		AdminService: AdminService,
		hub:          hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *ChatControllerImpl) CreateRoom(ctx echo.Context) error {
	var createRoomReq web.CreateRoomReq
	var user *domain.User
	if err := ctx.Bind(&createRoomReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error binding createRoomRequest",
		})
	}

	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	user, err := c.UserService.FindById(ctx, ID)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("User Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Id Error"))
	}

	room := schema.Rooms{
		ID:           user.ID,
		Name:         user.Name,
		PhotoProfile: user.ImageUrl,
	}
	config.DB.Create(&room)
	c.hub.Rooms[room.ID] = &ws.Room{
		ID:     room.ID,
		Name:   room.Name,
		Client: make(map[string]*ws.Client),
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Create Room", nil))
}

func (c *ChatControllerImpl) JoinRoom(ctx echo.Context) error {
	conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "error binding request",
		})

	}

	userToken := ctx.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)
	Role := claims["role"].(string)
	
	var clientID, username string
 	if Role == "user" {
		user, err := c.UserService.FindById(ctx, ID)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("User Not Found"))
			}
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Id Error"))
		}
		clientID = user.ID
		username = user.Name

	} else {
		admin, err := c.AdminService.FindById(ctx, ID)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("User Not Found"))
			}
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Id Error"))
		}
		clientID = admin.ID
		username = admin.Name
	}

	roomID := ctx.Param("roomId")

	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	c.hub.Register <- cl

	go cl.WriteMessage(roomID)
	cl.ReadMessage(c.hub)

	return nil
}

func (c *ChatControllerImpl) GetRooms(ctx echo.Context) error {
	rooms := make([]web.RoomRes, 0)

	for _, r := range c.hub.Rooms {
		rooms = append(rooms, web.RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return ctx.JSON(http.StatusOK, rooms)
}

func (c *ChatControllerImpl) GetClients(ctx echo.Context) error {
	var clients []web.ClientRes
	roomId := ctx.Param("roomId")

	if _, ok := c.hub.Rooms[roomId]; !ok {
		clients = make([]web.ClientRes, 0)
		ctx.JSON(http.StatusOK, clients)
	}

	for _, c := range c.hub.Rooms[roomId].Client {
		clients = append(clients, web.ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return ctx.JSON(http.StatusOK, clients)
}

func (c *ChatControllerImpl) GetChats(ctx echo.Context) error {
	roomID := ctx.Param("roomId")
	var messages []schema.Message
	config.DB.Where("room_id = ?", roomID).Find(&messages)

	return ctx.JSON(http.StatusOK, messages)
}
