package controller

import (
	"net/http"
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/res"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
)

type AdminController interface {
	RegisterAdminController(ctx echo.Context) error
	LoginAdminController(ctx echo.Context) error
	UpdateAdminController(ctx echo.Context) error
	ResetPasswordController(ctx echo.Context) error
	GetAdminController(ctx echo.Context) error
	GetAdminsController(ctx echo.Context) error
	GetAdminByNameController(ctx echo.Context) error
	DeleteAdminController(ctx echo.Context) error
}

type AdminControllerImpl struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{AdminService: adminService}
}

func (c *AdminControllerImpl) RegisterAdminController(ctx echo.Context) error {
	adminCreateRequest := web.AdminCreateRequest{}
	err := ctx.Bind(&adminCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminService.CreateAdmin(ctx, adminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (c *AdminControllerImpl) LoginAdminController(ctx echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	err := ctx.Bind(&adminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.AdminService.LoginAdmin(ctx, adminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := helper.GenerateAdminToken(&adminLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	adminLoginResponse.Token = token

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

func (c *AdminControllerImpl) GetAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.AdminService.FindById(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin Data", response))
}

func (c *AdminControllerImpl) GetAdminsController(ctx echo.Context) error {
	result, err := c.AdminService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "admins not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Admins Data Error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Admin Data", response))
}

func (c *AdminControllerImpl) GetAdminByNameController(ctx echo.Context) error {
	adminName := ctx.QueryParam("name")

	result, err := c.AdminService.FindByName(ctx, adminName)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data By Name Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin Data By Name", response))
}

func (c *AdminControllerImpl) UpdateAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	adminUpdateRequest := web.AdminUpdateRequest{}
	err = ctx.Bind(&adminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminService.UpdateAdmin(ctx, adminUpdateRequest, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated Admin Data", response))
}

func (c *AdminControllerImpl) ResetPasswordController(ctx echo.Context) error {
	resetPasswordRequest := web.AdminResetPasswordRequest{}
	err := ctx.Bind(&resetPasswordRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminService.ResetPassword(ctx, resetPasswordRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "new password and confirm new password do not match") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("New Password & Confirm New Password Do Not Match"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Reset Password", response))
}

func (c *AdminControllerImpl) DeleteAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.AdminService.DeleteAdmin(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
