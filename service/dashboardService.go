package service

import (
	"ecomplaint/repository"
	"github.com/labstack/echo/v4"
)

type DashboardService interface {
	GetDashboardService(ctx echo.Context) error
}

type DashboardServiceImpl struct {
	DashboardRepository repository.DashboardRepository
}
