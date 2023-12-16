package controller

import (
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type DashboardController interface {
	GetDashboardComplaintController(ctx echo.Context) error
}

type DashboardControllerImpl struct {
	DashboardRepo repository.DashboardRepository
}

func NewDashboardController(Repo repository.DashboardRepository) *DashboardControllerImpl {
	return &DashboardControllerImpl{DashboardRepo: Repo}
}

func (controller *DashboardControllerImpl) GetDashboardComplaintController(ctx echo.Context) error {
	result, _ := controller.DashboardRepo.CountComplaint("complaints")
	log.Println(result)
	conv := res.DashComplaintResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("success get complaint data", conv))
}

func (controller *DashboardControllerImpl) GetDashboardUserController(ctx echo.Context) error {
	result, _ := controller.DashboardRepo.CountComplaint("users")
	log.Println(result)
	conv := res.DashUserResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("success get users data", conv))
}

func (controller *DashboardControllerImpl) GetDashboardSolvedController(ctx echo.Context) error {
	result, _ := controller.DashboardRepo.CountSolved("complaints")
	log.Println(result)
	conv := res.DashUserResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("success get solved data", conv))
}

func (controller *DashboardControllerImpl) GetDashboardAllController(ctx echo.Context) error {
	result, _ := controller.DashboardRepo.TotalAll()
	log.Println(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("success get dashboard data", result))
}
