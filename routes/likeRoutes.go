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

func LikeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	likeRepositoryu := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepositoryu, validate)
	likeController := controller.NewLikeController(likeService)

	likeGroup := e.Group("/users/news/like")
	likeGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	likeGroup.POST("", likeController.CreateLikeController)
	likeGroup.GET("/search", likeController.GetLikeController)
	likeGroup.GET("", likeController.GetAllLikeController)
	likeGroup.PUT("/:id", likeController.UpdateLikeController)
	likeGroup.DELETE("/:id", likeController.DeleteLikeController)

	adminLikeGroup := e.Group("/admin/news/like")
	adminLikeGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))
	adminLikeGroup.GET("/search", likeController.GetLikeController)
	adminLikeGroup.GET("", likeController.GetAllLikeController)

}
