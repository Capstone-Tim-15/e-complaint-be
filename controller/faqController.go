package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/res"
	"github.com/labstack/echo/v4"
	"strings"
)

type FaqController interface {
	CreateFaqController(ctx echo.Context) error
	FindByIdController(ctx echo.Context) error
	FindByCategoryIDController(ctx echo.Context) error
	FindAllController(ctx echo.Context) error
	UpdateFaqController(ctx echo.Context) error
}

type FaqControllerImpl struct {
	FaqService service.FaqService
}

func NewFaqController(faqService service.FaqService) FaqController {
	return &FaqControllerImpl{FaqService: faqService}
}

func (ctrl *FaqControllerImpl) CreateFaqController(ctx echo.Context) error {
	FaqRequest := web.FaqRequest{}
	err := ctx.Bind(&FaqRequest)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	result, err := ctrl.FaqService.CreateFaq(ctx, FaqRequest)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(201, helper.SuccessResponse("Successfully Create FAQ", response))
}
func (ctrl *FaqControllerImpl) FindByIdController(ctx echo.Context) error {
	id := ctx.Param("id")
	result, err := ctrl.FaqService.FindById(ctx, id)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(200, helper.SuccessResponse("Successfully Find FAQ", response))
}
func (ctrl *FaqControllerImpl) FindByCategoryIDController(ctx echo.Context) error {
	category := ctx.QueryParam("category")
	result, err := ctrl.FaqService.FindByCategory(ctx, category)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(200, helper.SuccessResponse("Successfully Find FAQ by category id", response))
}
func (ctrl *FaqControllerImpl) FindAllController(ctx echo.Context) error {
	result, err := ctrl.FaqService.FindAll(ctx)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	response := res.ConvertFAQResponse(result)

	return ctx.JSON(200, helper.SuccessResponse("Successfully Find All FAQ", response))
}
func (ctrl *FaqControllerImpl) UpdateFaqController(ctx echo.Context) error {
	id := ctx.Param("id")
	FaqRequest := web.FaqRequest{}
	err := ctx.Bind(&FaqRequest)
	if err != nil {
		return ctx.JSON(400, helper.ErrorResponse("Invalid Client Input"))
	}
	result, err := ctrl.FaqService.UpdateFaq(ctx, FaqRequest, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(400, helper.ErrorResponse("Invalid Client Input"))
		}
		if strings.Contains(err.Error(), "FAQ not found") {
			return ctx.JSON(400, helper.ErrorResponse("FAQ Not Found"))
		}
		return ctx.JSON(400, helper.ErrorResponse("Update Error"))
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(200, helper.SuccessResponse("Successfully Update FAQ", response))
}
