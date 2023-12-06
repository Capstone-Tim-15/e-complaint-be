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

type UserController interface {
	RegisterUserController(ctx echo.Context) error
	LoginUserController(ctx echo.Context) error
	GetUserController(ctx echo.Context) error
	GetUsersController(ctx echo.Context) error
	UpdateUserController(ctx echo.Context) error
	ResetPasswordController(ctx echo.Context) error
	UpdatePhotoProfileController(ctx echo.Context) error
	DeleteUserController(ctx echo.Context) error
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (c *UserControllerImpl) RegisterUserController(ctx echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.UserService.CreateUser(ctx, userCreateRequest)
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

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (c *UserControllerImpl) LoginUserController(ctx echo.Context) error {
	userLoginRequest := web.UserLoginRequest{}
	err := ctx.Bind(&userLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.UserService.LoginUser(ctx, userLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	userLoginResponse := res.UserDomainToUserLoginResponse(response)

	token, err := middleware.GenerateToken(&userLoginResponse, response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	userLoginResponse.Token = token

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign In", userLoginResponse))
}

func (c *UserControllerImpl) GetUserController(ctx echo.Context) error {
	idQueryParam := ctx.QueryParam("id")
	nameQueryParam := ctx.QueryParam("name")

	if idQueryParam != "" && nameQueryParam != "" {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Both 'id' and 'name' query params are provided, please provide only one"))
	}

	if idQueryParam != "" {

		result, err := c.UserService.FindById(ctx, idQueryParam)
		if err != nil {
			if strings.Contains(err.Error(), "users not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Id Error"))
		}

		response := res.UserDomaintoUserResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User Data By Id", response))

	} else if nameQueryParam != "" {
		result, err := c.UserService.FindByName(ctx, nameQueryParam)
		if err != nil {
			if strings.Contains(err.Error(), "user not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Name Error"))
		}

		response := res.UserDomaintoUserResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User Data By Name", response))
	} else {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Query Param Input"))
	}
}

func (c *UserControllerImpl) GetUsersController(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize := 10

	result, totalCount, err := c.UserService.FindAll(ctx, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Users Data Error"))
	}

	response := res.ConvertUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get User Data", page, pageSize, totalCount, response))
}

func (c *UserControllerImpl) UpdateUserController(ctx echo.Context) error {
	userId := ctx.Param("id")

	userUpdateRequest := web.UserUpdateRequest{}
	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.UserService.UpdateUser(ctx, userUpdateRequest, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated User Data", response))
}

func (c *UserControllerImpl) ResetPasswordController(ctx echo.Context) error {
	resetPasswordRequest := web.UserResetPasswordRequest{}
	err := ctx.Bind(&resetPasswordRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := (claims["id"].(string))

	result, err := c.UserService.ResetPassword(ctx, resetPasswordRequest, ID)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "new password and confirm new password do not match") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("New Password & Confirm New Password Do Not Match"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Reset Password User Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Reset Password", response))
}

func (c *UserControllerImpl) UpdatePhotoProfileController(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Missing Image"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Opening File"))
	}
	defer file.Close()

	cldService, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Initializing Cloudinary"))
	}

	uploadParams := uploader.UploadParams{}
	resp, err := cldService.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Uploading File to Cloudinary"))
	}

	fileName := path.Base(resp.SecureURL)

	result, err := c.UserService.UpdatePhotoProfile(ctx, ID, fileName)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Photo Profile Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Update Photo Profile", response))

}

func (c *UserControllerImpl) DeleteUserController(ctx echo.Context) error {
	userId := ctx.Param("id")

	err := c.UserService.DeleteUser(ctx, userId)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete User Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
