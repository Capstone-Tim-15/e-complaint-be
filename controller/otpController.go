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
	CreateOTPUserController(ctx echo.Context) error
	CheckOTPUserController(ctx echo.Context) error
	CreateOTPAdminController(ctx echo.Context) error
	CheckOTPAdminController(ctx echo.Context) error
}

type OTPControllerImpl struct {
	OTPService service.OTPService
}

func NewOTPController(OTPService service.OTPService) *OTPControllerImpl {
	return &OTPControllerImpl{
		OTPService: OTPService,
	}
}

func (c *OTPControllerImpl) CreateOTPUserController(ctx echo.Context) error {
	otpRequest := web.OTPCreateRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.OTPService.CreateOTPUser(ctx, otpRequest)
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

func (c *OTPControllerImpl) CheckOTPUserController(ctx echo.Context) error {
	otpRequest := web.OTPCheckRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	result, err := c.OTPService.CheckOTPUser(ctx, otpRequest, ID)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "otp do not match") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("OTP Do Not Match"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Check OTP Error"))
	}

	response := res.OTPDomaintoOTPCheckResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Check OTP", response))
}

func (c *OTPControllerImpl) CreateOTPAdminController(ctx echo.Context) error {
	otpRequest := web.OTPCreateRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.OTPService.CreateOTPAdmin(ctx, otpRequest)
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

	response := res.AdminOTPDomaintoOTPResponse(result)

	token, err := middleware.GenerateTokenAdminID(response.Admin_ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	response.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Send OTP", response))
}

func (c *OTPControllerImpl) CheckOTPAdminController(ctx echo.Context) error {
	otpRequest := web.OTPCheckRequest{}
	err := ctx.Bind(&otpRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	result, err := c.OTPService.CheckOTPAdmin(ctx, otpRequest, ID)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "otp do not match") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("OTP Do Not Match"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Check OTP Error"))
	}

	response := res.AdminOTPDomaintoOTPCheckResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Check OTP", response))
}
