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

func OTPRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(db)
	OTPRepository := repository.NewOTPRepository(db)
	OTPService := service.NewOTPService(OTPRepository, userRepository, validate)
	OTPController := controller.NewOTPController(OTPService)

	otpGroups := e.Group("otp")
	
	otpGroups.POST("/send-otp", OTPController.CreateOTPController)

	otpGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	otpGroups.POST("/check-otp", OTPController.CheckOTPController)
}
