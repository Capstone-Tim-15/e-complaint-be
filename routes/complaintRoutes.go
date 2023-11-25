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

func ComplaintRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	complaintRepository := repository.NewComplaintRepository(db)
	complaintService := service.NewComplaintService(complaintRepository, validate)
	complaintController := controller.NewComplaintController(complaintService)

	complaintGroups := e.Group("complaint")

	complaintGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	complaintGroups.POST("", complaintController.CreateComplaintController)
	complaintGroups.GET("/search", complaintController.GetComplaintController)
	complaintGroups.GET("", complaintController.GetComplaintsController)
}
