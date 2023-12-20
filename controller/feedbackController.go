package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type FeedbackController interface {
	GetFeedbackController(ctx echo.Context) error
	GetAllFeedbackController(ctx echo.Context) error
	CreateFeedbackController(ctx echo.Context) error
	UpdateFeedbackController(ctx echo.Context) error
	DeleteFeedbackController(ctx echo.Context) error
}

type FeedbackControllerImpl struct {
	FeedbackService service.FeedbackService
}

func NewFeedbackController(feedbackService service.FeedbackService) FeedbackController {
	return &FeedbackControllerImpl{FeedbackService: feedbackService}
}

func (c *FeedbackControllerImpl) GetFeedbackController(ctx echo.Context) error {
	feedbackID := ctx.QueryParam("id")
	feedbackNewsId := ctx.QueryParam("news_id")
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize := 10
	if feedbackID != "" {
		result, err := c.FeedbackService.FindById(ctx, feedbackID)
		if err != nil {
			if strings.Contains(err.Error(), "feedback not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Feedback Not Found"))
			}
		}

		response := res.FindFeedbackDomainToFeedbackResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Feedback Data", response))
	} else if feedbackNewsId != "" {
		result, totalCount, err := c.FeedbackService.FindByNewsId(ctx, feedbackNewsId, page, pageSize)
		if err != nil {
			if strings.Contains(err.Error(), "feedback not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Feedback Not Found"))
			}
		}

		response := res.ConvertFeedbackResponse(result)
		return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Feedback By news_id", page, pageSize, totalCount, response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Query Param Input"))
	}
}

func (c *FeedbackControllerImpl) GetAllFeedbackController(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	pageSize := 10
	result, totalCount, err := c.FeedbackService.FindByAll(ctx, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "feedback not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Feedback Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Feedback Data Error"))
	}
	response := res.ConvertFeedbackResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get All Feedback Data", page, pageSize, totalCount, response))
}

func (c *FeedbackControllerImpl) CreateFeedbackController(ctx echo.Context) error {
	feedbackCreateRequest := web.FeedbackCreateRequest{}
	err := ctx.Bind(&feedbackCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)
	userResult, err := c.FeedbackService.CheckUser(ID)
	if err == nil {
		feedbackCreateRequest.Fullname = userResult.Name
		feedbackCreateRequest.PhotoImage = userResult.ImageUrl
		feedbackCreateRequest.Role = "user"
	} else {
		adminResult, err := c.FeedbackService.CheckAdmin(ID)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Error When Check Sender Account"))
		}
		feedbackCreateRequest.Fullname = adminResult.Name
		feedbackCreateRequest.PhotoImage = adminResult.ImageUrl
		feedbackCreateRequest.Role = "admin"
	}

	result, err := c.FeedbackService.CreateFeedback(ctx, feedbackCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Feedback Error"))
	}
	response := res.FeedbackDomainToFeedbackResponse(result)
	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Feedback", response))
}

func (c *FeedbackControllerImpl) UpdateFeedbackController(ctx echo.Context) error {
	feedbackID := ctx.Param("id")
	feedbackUpdateRequest := web.FeedbackUpdateRequest{}
	err := ctx.Bind(&feedbackUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	result, err := c.FeedbackService.UpdateFeedback(ctx, feedbackUpdateRequest, feedbackID)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "error when updating feedback") {
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error When Updating Feedback"))
		}
		if strings.Contains(err.Error(), "feedback not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Feedback Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Feedback Error"))
	}
	response := res.FeedbackDomainToFeedbackResponse(result)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update Feedback", response))
}

func (c *FeedbackControllerImpl) DeleteFeedbackController(ctx echo.Context) error {
	feedbackID := ctx.Param("id")
	err := c.FeedbackService.DeleteFeedback(ctx, feedbackID)
	if err != nil {
		if strings.Contains(err.Error(), "feedback not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Feedback Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Feedback Error"))
	}
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Feedback", nil))
}
