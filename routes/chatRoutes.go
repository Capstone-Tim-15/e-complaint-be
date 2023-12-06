package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"ecomplaint/utils/helper/websocket"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChatRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate, hub *websocket.Hub) {
	userRepository := repository.NewUserRepository(db)
	adminRepository := repository.NewAdminRepository(db)
	userService := service.NewUserService(userRepository, validate)
	adminService := service.NewAdminService(adminRepository, validate)
	chatController := controller.NewChatController(userService, adminService, hub)

	chatGroups := e.Group("chat/user/ws")
	chatAdminGroups := e.Group("chat/admin/ws")

	chatGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	chatGroups.POST("/create-room", chatController.CreateRoom)
	chatGroups.GET("/join-room/:roomId", chatController.JoinRoom)
	chatGroups.GET("/get-rooms", chatController.GetRooms)
	chatGroups.GET("/get-clients/:roomId", chatController.GetClients)
	chatGroups.GET("/get-chats/:roomId", chatController.GetChats)

	chatAdminGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))

	chatAdminGroups.POST("/create-room", chatController.CreateRoom)
	chatAdminGroups.GET("/join-room/:roomId", chatController.JoinRoom)
	chatAdminGroups.GET("/get-rooms", chatController.GetRooms)
	chatAdminGroups.GET("/get-clients/:roomId", chatController.GetClients)
	chatAdminGroups.GET("/get-chats/:roomId", chatController.GetChats)

}
