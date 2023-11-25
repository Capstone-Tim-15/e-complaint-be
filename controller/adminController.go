package controller

import (
	"context"
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/helper/middleware"
	res "ecomplaint/utils/response"
	"net/http"
	"os"
	"path"
	"strconv"

	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminController interface {
	RegisterAdminController(ctx echo.Context) error
	LoginAdminController(ctx echo.Context) error
	GetAdminController(ctx echo.Context) error
	GetAdminsController(ctx echo.Context) error
	UpdateAdminController(ctx echo.Context) error
	ResetPasswordController(ctx echo.Context) error
	UpdatePhotoProfileController(ctx echo.Context) error
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

		if strings.Contains(err.Error(), "username already exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Username Already Exist"))

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

	token, err := middleware.GenerateAdminToken(&adminLoginResponse, response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	adminLoginResponse.Token = token

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

func (c *AdminControllerImpl) GetAdminController(ctx echo.Context) error {
	idQueryParam := ctx.QueryParam("id")
	nameQueryParam := ctx.QueryParam("name")

	if idQueryParam != "" && nameQueryParam != "" {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Both 'id' and 'name' query params are provided, please provide only one"))
	}

	if idQueryParam != "" {
		result, err := c.AdminService.FindById(ctx, idQueryParam)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data By Id Error"))
		}

		response := res.AdminDomaintoAdminResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin Data By Id", response))

	} else if nameQueryParam != "" {
		result, err := c.AdminService.FindByName(ctx, nameQueryParam)
		if err != nil {
			if strings.Contains(err.Error(), "user not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data By Name Error"))
		}

		response := res.AdminDomaintoAdminResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin Data By Name", response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Query Param Input"))
	}
}

func (c *AdminControllerImpl) GetAdminsController(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize := 10

	result, totalCount, err := c.AdminService.FindAll(ctx, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admins Data Error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Admin Data", page, pageSize, totalCount, response))
}

func (c *AdminControllerImpl) UpdateAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")

	adminUpdateRequest := web.AdminUpdateRequest{}
	err := ctx.Bind(&adminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminService.UpdateAdmin(ctx, adminUpdateRequest, adminId)
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

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	result, err := c.AdminService.ResetPassword(ctx, resetPasswordRequest, ID)
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

func (c *AdminControllerImpl) UpdatePhotoProfileController(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Missing attachment"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error opening file"))
	}
	defer file.Close()

	cldService, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error initializing Cloudinary"))
	}

	uploadParams := uploader.UploadParams{}
	resp, err := cldService.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error uploading file to Cloudinary"))
	}

	fileName := path.Base(resp.SecureURL)

	result, err := c.AdminService.UpdatePhotoProfile(ctx, ID, fileName)
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

	err := c.AdminService.DeleteAdmin(ctx, adminId)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
