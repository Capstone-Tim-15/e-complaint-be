package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
)

func NewsRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	newsRepository := repository.NewNewsRepository(db)
	newsService := service.NewNewsService(newsRepository, validate)
	newsController := controller.NewNewsController(newsService)

	newsGroup := e.Group("news")
	newsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	newsGroup.POST("", newsController.CreateNewsController)
	newsGroup.GET("/search", newsController.GetNewsController)
	newsGroup.GET("", newsController.GetAllNewsController)
	newsGroup.PUT("/:id", newsController.UpdateNewsController)
	newsGroup.DELETE("/:id", newsController.DeleteNewsController)

}
