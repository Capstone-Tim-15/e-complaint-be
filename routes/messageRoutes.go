package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MessRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	messRepository := repository.NewMessRepository(db)
	messService := service.NewMessService(messRepository, validate)
	messController := controller.NewMessController(messService)

	messGroups := e.Group("comment")

	messGroups.POST("", messController.CreateCommentController)
}
