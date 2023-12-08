package routes

import (
	"ecomplaint/controller"
	"ecomplaint/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DashboardRoutes(e *echo.Echo, db *gorm.DB) {

	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardController := controller.NewDashboardController(dashboardRepo)

	dashboardGroup := e.Group("admin/dashboard")

	dashboardGroup.GET("/complaint", dashboardController.GetDashboardComplaintController)
	dashboardGroup.GET("/solved", dashboardController.GetDashboardSolvedController)
	dashboardGroup.GET("/user", dashboardController.GetDashboardUserController)

}
