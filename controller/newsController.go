package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

type NewsController interface {
	GetNewsController(ctx echo.Context) error
	GetAllNewsController(ctx echo.Context) error
	CreateNewsController(ctx echo.Context) error
	UpdateNewsController(ctx echo.Context) error
	DeleteNewsController(ctx echo.Context) error
}

type NewsControllerImpl struct {
	NewsService service.NewsService
}

func NewNewsController(newsService service.NewsService) NewsController {
	return &NewsControllerImpl{NewsService: newsService}
}

func (c *NewsControllerImpl) GetNewsController(ctx echo.Context) error {
	newsID := ctx.QueryParam("id")
	newsTitle := ctx.QueryParam("title")
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize := 10

	if newsID != "" {
		result, err := c.NewsService.FindById(ctx, newsID)
		if err != nil {
			if strings.Contains(err.Error(), "news not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("News Not Found"))
			}
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get News Error"))
		}
		response := res.NewsDomainToNewsResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get News Data", response))
	} else if newsTitle != "" {
		result, totalCount, err := c.NewsService.FindByTitle(ctx, newsTitle, page, pageSize)
		if err != nil {
			if strings.Contains(err.Error(), "news not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("News Not Found"))
			}
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get News Error"))
		}
		response := res.ConvertNewsResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get News By Title", page, pageSize, totalCount, response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Query Param Input"))
	}

}

func (c *NewsControllerImpl) GetAllNewsController(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize := 10
	result, totalCount, err := c.NewsService.FindByAll(ctx, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "news not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("News Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All News Data Error"))
	}

	response := res.ConvertNewsResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get All News Data", page, pageSize, totalCount, response))
}

func (c *NewsControllerImpl) CreateNewsController(ctx echo.Context) error {
	newsCreateRequest := web.NewsCreateRequest{}
	err := ctx.Bind(&newsCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	result, err := c.NewsService.CreateNews(ctx, newsCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create News Error"))
	}
	response := res.NewsDomainToNewsResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create News", response))

}

func (c *NewsControllerImpl) UpdateNewsController(ctx echo.Context) error {
	newsId := ctx.Param("id")
	newsUpdateRequest := web.NewsUpdateRequest{}
	err := ctx.Bind(&newsUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	result, err := c.NewsService.UpdateNews(ctx, newsUpdateRequest, newsId)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "news not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("News Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update News Error"))
	}
	response := res.NewsDomainToNewsResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated News Data", response))
}

func (c *NewsControllerImpl) DeleteNewsController(ctx echo.Context) error {
	newsId := ctx.Param("id")
	err := c.NewsService.DeleteNews(ctx, newsId)
	if err != nil {
		if strings.Contains(err.Error(), "news not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("News Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete News Data Error"))
	}
	return ctx.JSON(http.StatusNoContent, nil)
}
