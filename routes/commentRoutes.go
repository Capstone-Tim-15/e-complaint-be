package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	messRepository := repository.NewCommentRepository(db)
	messService := service.NewCommentService(messRepository, validate)
	messController := controller.NewCommentController(messService)

	messGroups := e.Group("comment")

	messGroups.POST("", messController.CreateCommentController)
}
