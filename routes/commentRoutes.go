package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CommentRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	messRepository := repository.NewCommentRepository(db)
	messService := service.NewCommentService(messRepository, validate)
	messController := controller.NewCommentController(messService)

	userMessGroups := e.Group("user/comment")
	userMessGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	userMessGroups.POST("", messController.CreateCommentController)

	adminMessGroups := e.Group("admin/comment")
	adminMessGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))
	adminMessGroups.POST("", messController.CreateCommentController)

}
