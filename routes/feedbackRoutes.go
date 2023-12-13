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

	feedbackGroup := e.Group("/users/news/feedback")
	feedbackGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	feedbackGroup.GET("/search", feedbackController.GetFeedbackController)
	feedbackGroup.GET("", feedbackController.GetAllFeedbackController)
	feedbackGroup.POST("", feedbackController.CreateFeedbackController)

	adminFeedbackGroup := e.Group("/admin/news/feedback")
	adminFeedbackGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))
	adminFeedbackGroup.POST("", feedbackController.CreateFeedbackController)
	adminFeedbackGroup.GET("/search", feedbackController.GetFeedbackController)
	adminFeedbackGroup.GET("", feedbackController.GetAllFeedbackController)
	adminFeedbackGroup.PUT("/:id", feedbackController.UpdateFeedbackController)
	adminFeedbackGroup.DELETE("/:id", feedbackController.DeleteFeedbackController)

}
