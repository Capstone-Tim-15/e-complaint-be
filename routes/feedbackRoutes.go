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

func FeedbackRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	feedbackRepository := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepository, validate)
	feedbackController := controller.NewFeedbackController(feedbackService)

	feedbackGroup := e.Group("feedback")
	feedbackGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	feedbackGroup.POST("", feedbackController.CreateFeedbackController)
	feedbackGroup.GET("/search", feedbackController.GetFeedbackController)
	feedbackGroup.GET("", feedbackController.GetAllFeedbackController)
	feedbackGroup.PUT("/:id", feedbackController.UpdateFeedbackController)
	feedbackGroup.DELETE("/:id", feedbackController.DeleteFeedbackController)

}
