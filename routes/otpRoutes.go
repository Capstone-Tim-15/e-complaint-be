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
	adminRepository := repository.NewAdminRepository(db)
	OTPRepository := repository.NewOTPRepository(db)
	OTPService := service.NewOTPService(OTPRepository, userRepository, adminRepository, validate)
	OTPController := controller.NewOTPController(OTPService)

	otpGroups := e.Group("otp/user")
	otpAdminGroups := e.Group("otp/admin")

	otpGroups.POST("/send-otp", OTPController.CreateOTPUserController)

	otpGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	otpGroups.POST("/check-otp", OTPController.CheckOTPUserController)

	otpAdminGroups.POST("/send-otp", OTPController.CreateOTPAdminController)

	otpAdminGroups.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))

	otpAdminGroups.POST("/check-otp", OTPController.CheckOTPAdminController)
}
