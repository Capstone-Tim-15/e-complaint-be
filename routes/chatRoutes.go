package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChatRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	chatController := controller.NewChatController(userService)

	chatGroups := e.Group("chat")

	chatGroups.GET("/ws/:id", chatController.HandleWebsocket)
}
