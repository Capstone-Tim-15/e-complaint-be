package controller

import (
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/helper/middleware"
	res "ecomplaint/utils/response"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OTPController interface {
	CreateOTPController(ctx echo.Context) error
}

type OTPControllerImpl struct {
	OTPService service.OTPService
}

func NewOTPController(OTPService service.OTPService) *OTPControllerImpl {
	return &OTPControllerImpl{
		OTPService: OTPService,
	}
}

func (c *OTPControllerImpl) CreateOTPController(ctx echo.Context) error {
	otpRequest := web.OTPCreateRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.OTPService.CreateOTP(ctx, otpRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "email not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Email Not Found"))
		}

		if strings.Contains(err.Error(), "send otp error") {
			return ctx.JSON(http.StatusBadGateway, helper.ErrorResponse("Send OTP Error"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create OTP Error"))
	}

	response := res.OTPDomaintoOTPResponse(result)

	token, err := middleware.GenerateTokenUserID(response.User_ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	response.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Send OTP", response))
}

func (c *OTPControllerImpl) CheckOTPController(ctx echo.Context) error {
	otpRequest := web.OTPCheckRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	result, err := c.OTPService.CheckOTP(ctx, otpRequest, ID)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "failed to extract user ID from token") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Failed Extract JWT Token User Id"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Check OTP Error"))
	}

	response := res.OTPDomaintoOTPCheckResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Check OTP", response))
}
