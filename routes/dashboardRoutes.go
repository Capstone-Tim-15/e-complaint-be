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

func DashboardRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardControl := controller.NewDashboardController()

	faqRepository := repository.NewFAQRepository(db)
	faqService := service.NewFaqService(faqRepository, validate)
	faqController := controller.NewFaqController(faqService)

	faqsGroup.POST("", faqController.CreateFaqController)
	faqsGroup.GET("/search/:id", faqController.FindController)
	faqsGroup.GET("/search", faqController.FindController)
	faqsGroup.PUT("/:id", faqController.UpdateFaqController)
}
