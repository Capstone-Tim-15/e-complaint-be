package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FeedbackRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	feedbackRepository := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepository, validate)
	feedbackController := controller.NewFeedbackController(feedbackService)

	feedbackGroup := e.Group("feedback")
	feedbackGroup.POST("", feedbackController.CreateFeedbackController)
	feedbackGroup.GET("/search", feedbackController.GetFeedbackController)
	feedbackGroup.GET("", feedbackController.GetAllFeedbackController)
	feedbackGroup.PUT("/:id", feedbackController.UpdateFeedbackController)
	feedbackGroup.DELETE("/:id", feedbackController.DeleteFeedbackController)

}
