package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryGroup := e.Group("category")
	categoryGroup.POST("/create", categoryController.CreateCategoryController)
	categoryGroup.GET("/find/:id", categoryController.FindController)
	categoryGroup.GET("/find", categoryController.FindController)
	categoryGroup.PUT("/update/:id", categoryController.UpdateCategoryController)
}
