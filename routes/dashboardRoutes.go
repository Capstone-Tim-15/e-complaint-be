package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
)

func DashboardRoutes(e *echo.Echo, db *gorm.DB) {

	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardController := controller.NewDashboardController(dashboardRepo)

	dashboardGroup := e.Group("admin/dashboard")
	dashboardGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_ADMIN"))))

	dashboardGroup.GET("/complaint", dashboardController.GetDashboardComplaintController)
	dashboardGroup.GET("/solved", dashboardController.GetDashboardSolvedController)
	dashboardGroup.GET("/user", dashboardController.GetDashboardUserController)
	dashboardGroup.GET("", dashboardController.GetDashboardAllController)

}
