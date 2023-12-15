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

func NotificationRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	NotificationRepository := repository.NewNotificationRepository(db)
	NotificationService := service.NewNotificationService(NotificationRepository, validate)
	NotificationController := controller.NewNotificationController(NotificationService)

	notificationGroup := e.Group("admin/notification")
	notificationGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))
	notificationGroup.GET("", NotificationController.FindAllNotification)
}
