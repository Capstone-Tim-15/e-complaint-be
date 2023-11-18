package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FAQRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	faqRepository := repository.NewFAQRepository(db)
	faqService := service.NewFaqService(faqRepository, validate)
	faqController := controller.NewFaqController(faqService)

	faqsGroup := e.Group("faq")

	faqsGroup.POST("", faqController.CreateFaqController)
	faqsGroup.GET("/:id", faqController.FindByIdController)
	faqsGroup.GET("", faqController.FindAllController)
	faqsGroup.GET("", faqController.FindByCategoryIDController)
	faqsGroup.PUT("/:id", faqController.UpdateFaqController)
}
