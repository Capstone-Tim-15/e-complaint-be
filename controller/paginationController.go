package controller

import (
	"ecomplaint/utils/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (categoryHandler *CategoryControllerImpl) Pagination(ctx echo.Context) error {
	params := ctx.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("params limit not valid"))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("params offset not valid"))
	}

	response, pagination, err := categoryHandler.CategoryService.Pagination(offset, limit)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Category Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return ctx.JSON(http.StatusOK, helper.PaginationResponse("Medicines Data Successfully Retrieved", response, pagination))
}
