package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController interface {
	CreateCategoryController(ctx echo.Context) error
	FindController(ctx echo.Context) error
	UpdateCategoryController(ctx echo.Context) error
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{CategoryService: categoryService}
}

func (ctrl *CategoryControllerImpl) CreateCategoryController(ctx echo.Context) error {
	CategoryReq := web.CategoryRequest{}
	err := ctx.Bind(&CategoryReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	result, err := ctrl.CategoryService.CreateCategory(ctx, CategoryReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	response := res.CategoryCreateResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Success create data", response))
}

func (ctrl *CategoryControllerImpl) FindController(ctx echo.Context) error {
	paramID := ctx.Param("id")
	queryID := ctx.QueryParam("id")
	if paramID != "" {
		result, err := ctrl.CategoryService.FindById(ctx, paramID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.CategoryResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find data", response))
	} else if queryID != "" {
		result, err := ctrl.CategoryService.FindById(ctx, queryID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.CategoryResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find data", response))
	} else if paramID == "" && queryID == "" {
		result, err := ctrl.CategoryService.FindAll(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.AllCategoryResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success find all data", response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("ID not found"))
	}
}

func (ctrl *CategoryControllerImpl) UpdateCategoryController(ctx echo.Context) error {
	paramID := ctx.Param("id")
	queryID := ctx.QueryParam("id")
	if paramID != "" {
		CategoryReq := web.CategoryRequest{}
		err := ctx.Bind(&CategoryReq)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		result, err := ctrl.CategoryService.UpdateCategory(ctx, CategoryReq, paramID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.CategoryResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success update data", response))
	} else if queryID != "" {
		CategoryReq := web.CategoryRequest{}
		err := ctx.Bind(&CategoryReq)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		result, err := ctrl.CategoryService.UpdateCategory(ctx, CategoryReq, queryID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		response := res.CategoryResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Success update data", response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("ID not found"))
	}
}
