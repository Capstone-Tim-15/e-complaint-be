package routes

import (
	"ecomplaint/controller"
	"ecomplaint/service"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func AIRoutes(e *echo.Echo, client *openai.Client) {
	aiService := service.NewAIService(client)
	aiController := controller.NewAIController(aiService)

	aiGroups := e.Group("ai")

	aiGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	aiGroups.POST("", aiController.AIRecomController)
}
