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

type UserController interface {
	RegisterUserController(ctx echo.Context) error
	LoginUserController(ctx echo.Context) error
	UpdateUserController(ctx echo.Context) error
	ResetPasswordController(ctx echo.Context) error
	GetUserController(ctx echo.Context) error
	GetUsersController(ctx echo.Context) error
	GetUserByNameController(ctx echo.Context) error
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

	token, err := helper.GenerateToken(&userLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	userLoginResponse.Token = token

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign In", userLoginResponse))
}

func (c *UserControllerImpl) UpdateUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	userUpdateRequest := web.UserUpdateRequest{}
	err = ctx.Bind(&userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.UserService.UpdateUser(ctx, userUpdateRequest, userIdInt)
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

	result, err := c.UserService.ResetPassword(ctx, resetPasswordRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "new password and confirm new password do not match") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("New Password & Confirm New Password Do Not Match"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Reset Password", response))
}

func (c *UserControllerImpl) GetUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.UserService.FindById(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Users Data Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Users Data", response))
}

func (c *UserControllerImpl) GetUsersController(ctx echo.Context) error {
	result, err := c.UserService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Users Data Error"))
	}

	response := res.ConvertUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User Data", response))
}

func (c *UserControllerImpl) GetUserByNameController(ctx echo.Context) error {
	userName := ctx.QueryParam("name")

	result, err := c.UserService.FindByName(ctx, userName)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data By Name Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User Data By Name", response))
}

func (c *UserControllerImpl) DeleteUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.UserService.DeleteUser(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete User Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
