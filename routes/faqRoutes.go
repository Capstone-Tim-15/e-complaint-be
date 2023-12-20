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

func FAQRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	faqRepository := repository.NewFAQRepository(db)
	faqService := service.NewFaqService(faqRepository, validate)
	faqController := controller.NewFaqController(faqService)

	faqsGroup := e.Group("admin/faq")
	faqsUser := e.Group("user/faq")
	faqsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))

	faqsGroup.POST("", faqController.CreateFaqController)
	faqsGroup.GET("/search/:id", faqController.FindController)
	faqsGroup.GET("/search", faqController.FindController)
	faqsGroup.PUT("/:id", faqController.UpdateFaqController)

	faqsUser.GET("/search/:id", faqController.FindController)
	faqsUser.GET("/search", faqController.FindController)
	faqsUser.GET("", faqController.FindController)
}
