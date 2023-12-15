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

func CategoryRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryGroup := e.Group("admin/category")
	categoryUser := e.Group("user/category")
	categoryGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))

	categoryGroup.POST("", categoryController.CreateCategoryController)
	categoryGroup.GET("/:id", categoryController.FindController)
	categoryGroup.GET("", categoryController.FindController)
	categoryGroup.PUT("/:id", categoryController.UpdateCategoryController)

	// for mobile usage
	categoryUser.GET("/:id", categoryController.FindController)
	categoryUser.GET("", categoryController.FindController)

}
