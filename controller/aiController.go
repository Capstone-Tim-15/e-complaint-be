package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AIController interface {
	AIRecomController(ctx echo.Context) error
}

type AIControllerImpl struct {
	AIService service.AIService
}

func NewAIController(AIService service.AIService) *AIControllerImpl {
	return &AIControllerImpl{AIService: AIService}
}

func (c *AIControllerImpl) AIRecomController(ctx echo.Context) error {
	aiCreateRequest := web.AICreateRequest{}
	err := ctx.Bind(&aiCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	if aiCreateRequest.Message == "" {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Complaint message cannot be empty"))
	}

	result, err := c.AIService.ResolveComplaint(ctx, aiCreateRequest.Message)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Generating Recommendation"))
	}

	response := map[string]interface{}{
		"complaint":      aiCreateRequest.Message,
		"recommendation": result,
		"timestamp":      time.Now().Format(time.RFC3339),
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Generated Recommendation", response))
}
