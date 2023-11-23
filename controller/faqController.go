package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type FaqController interface {
	CreateFaqController(ctx echo.Context) error
	FindController(ctx echo.Context) error
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
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := ctrl.FaqService.CreateFaq(ctx, FaqRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(201, helper.SuccessResponse("Successfully Create FAQ", response))
}
func (ctrl *FaqControllerImpl) FindController(ctx echo.Context) error {
	paramID := ctx.Param("id")
	queryID := ctx.QueryParam("id")
	categoryID := ctx.QueryParam("category")
	if categoryID != "" {
		result, err := ctrl.FaqService.FindByCategory(ctx, categoryID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}
		response := res.FAQDomaintoResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Find FAQ by category id", response))
	} else if paramID != "" {
		result, err := ctrl.FaqService.FindById(ctx, paramID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.FAQDomaintoResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find data", response))
	} else if queryID != "" {
		result, err := ctrl.FaqService.FindById(ctx, queryID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.FAQDomaintoResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find data", response))
	} else if paramID == "" && queryID == "" && categoryID == "" {
		result, err := ctrl.FaqService.FindAll(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.ConvertFAQResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find all data", response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("ID not found"))
	}
}

func (ctrl *FaqControllerImpl) UpdateFaqController(ctx echo.Context) error {
	id := ctx.Param("id")
	FaqRequest := web.FaqUpdateRequest{}
	err := ctx.Bind(&FaqRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	result, err := ctrl.FaqService.UpdateFaq(ctx, FaqRequest, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
		}
		if strings.Contains(err.Error(), "FAQ not found") {
			return ctx.JSON(400, helper.ErrorResponse("FAQ Not Found"))
		}
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Update Error"))
	}
	response := res.FAQDomaintoResponse(result)

	return ctx.JSON(200, helper.SuccessResponse("Successfully Update FAQ", response))
}
