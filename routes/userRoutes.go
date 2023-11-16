package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"ecomplaint/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	usersGroup := e.Group("user")

	usersGroup.POST("/register", userController.RegisterUserController)
	usersGroup.POST("/login", userController.LoginUserController)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.GET("/search", userController.GetUserController)
	usersGroup.GET("", userController.GetUsersController)
	usersGroup.PUT("/:id", userController.UpdateUserController)
	usersGroup.PUT("/reset-password", userController.ResetPasswordController)
	usersGroup.DELETE("/:id", userController.DeleteUserController)

}
