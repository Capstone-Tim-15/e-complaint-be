package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func LikeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	likeRepositoryu := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepositoryu, validate)
	likeController := controller.NewLikeController(likeService)

	likeGroup := e.Group("like")
	likeGroup.POST("", likeController.CreateLikeController)
	likeGroup.GET("/search", likeController.GetLikeController)
	likeGroup.GET("", likeController.GetAllLikeController)
	likeGroup.PUT("/:id", likeController.UpdateLikeController)
	likeGroup.DELETE("/:id", likeController.DeleteLikeController)
}
