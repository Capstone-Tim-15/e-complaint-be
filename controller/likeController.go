package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type LikeController interface {
	GetLikeController(ctx echo.Context) error
	GetAllLikeController(ctx echo.Context) error
	CreateLikeController(ctx echo.Context) error
	UpdateLikeController(ctx echo.Context) error
	DeleteLikeController(ctx echo.Context) error
}

type LikeControllerImpl struct {
	LikeService service.LikeService
}

func NewLikeController(likeService service.LikeService) LikeController {
	return &LikeControllerImpl{LikeService: likeService}
}

func (c *LikeControllerImpl) GetLikeController(ctx echo.Context) error {

	likeID := ctx.QueryParam("id")
	result, err := c.LikeService.FindById(ctx, likeID)
	if err != nil {
		if strings.Contains(err.Error(), "like not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Like Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Like Error"))
	}
	if result == nil {
		return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Like not found"))
	}
	response := res.LikesDomainToLikesResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Like Data", response))
}

func (c *LikeControllerImpl) GetAllLikeController(ctx echo.Context) error {
	result, err := c.LikeService.FindByAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "like not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Like Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Like Data Error"))
	}
	response := res.ConvertLikesResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Like Data", response))
}

func (c *LikeControllerImpl) CreateLikeController(ctx echo.Context) error {
	likeCreateRequest := web.LikesCreateRequest{}
	err := ctx.Bind(&likeCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
	}
	result, err := c.LikeService.CreateLike(ctx, likeCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Like Error"))
	}
	response := res.LikesDomainToLikesResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Create Like Data", response))
}

func (c *LikeControllerImpl) UpdateLikeController(ctx echo.Context) error {
	likeID := ctx.Param("id")
	likeUpdateRequest := web.LikesUpdateRequest{}
	err := ctx.Bind(&likeUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
	}
	result, err := c.LikeService.UpdateLike(ctx, likeUpdateRequest, likeID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Like Error"))
	}
	response := res.LikesDomainToLikesResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update Like Data", response))
}

func (c *LikeControllerImpl) DeleteLikeController(ctx echo.Context) error {
	likeID := ctx.Param("id")
	err := c.LikeService.DeleteLike(ctx, likeID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Like Error"))
	}
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Like Data", nil))
}
