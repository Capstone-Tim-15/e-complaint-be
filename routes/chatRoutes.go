package routes

import (
	"ecomplaint/controller"

	"github.com/labstack/echo/v4"
)

func ChatRoutes(e *echo.Echo) {

	chatGroups := e.Group("chat")

	chatGroups.GET("/ws", controller.SendChat)
}
