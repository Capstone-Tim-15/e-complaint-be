package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentController interface {
	CreateCommentController(ctx echo.Context) error
}

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(CommentService service.CommentService) *CommentControllerImpl {
	return &CommentControllerImpl{CommentService: CommentService}
}

func (c *CommentControllerImpl) CreateCommentController(ctx echo.Context) error {
	messRequest := web.CommentCreateRequest{}
	err := ctx.Bind(&messRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)
	fmt.Println("masuk lewat extract id	")
	userResult, err := c.CommentService.CheckUser(ID)
	if err == nil {
		messRequest.Fullname = userResult.Name
		messRequest.Role = "user"
	} else {
		adminResult, err := c.CommentService.CheckAdmin(ID)
		if err != nil {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Error When Check Sender Account"))
		}
		messRequest.Fullname = adminResult.Name
		messRequest.Role = "admin"
	}

	result, err := c.CommentService.CreateComment(ctx, messRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Comment Error"))
	}

	response := res.CommentDomaintoCommentResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Comment", response))
}
