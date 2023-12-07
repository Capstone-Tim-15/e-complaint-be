package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DashboardController interface {
	GetDashboardController(ctx echo.Context) error
}

type DashboardControllerImpl struct {
	ComplaintRepo repository.DashboardRepository
}
func NewDashboardController(Repo repository.DashboardRepository) *DashboardControllerImpl {
	return &DashboardControllerImpl{ComplaintRepo: Repo}
}

func (controller *DashboardControllerImpl) GetDashboardController(ctx echo.Context) error {

	totalUser, err := controller.ComplaintRepo.CountUser()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	totalComplaint, err := controller.ComplaintRepo.CountComplaint()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	totalResolved, err := controller.ComplaintRepo.CountResolved()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	var dashboard web.DashboardResponse
	dashboard.TotalUser = totalUser
	dashboard.TotalComplaint = totalComplaint
	dashboard.TotalResolved = totalResolved

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("success get dashboard data", dashboard))
}
