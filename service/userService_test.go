package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	request	:= web.UserCreateRequest{
		Name: "test",
		Username: "test",
		Email: "test@test.com",
		Phone: "123456789",
		Password: "testing1",
	}

	mockUserRepository.On("FindByEmail", "test@test.com").Return(nil, nil)
	mockUserRepository.On("FindByUsername", request.Username).Return(nil, nil)
	mockUserRepository.On("Create", mock.AnythingOfType("*domain.User")).Return(nil, nil)

	_, err := UserService.CreateUser(ctx, request)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestFindByUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}
	userId := "123456"

	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)

	_, err := UserService.FindById(ctx, userId)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestFindAllUsers(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	mockUserRepository.On("FindAll", 1, 10).Return([]domain.User{}, int64(0), nil)

	_, _, err := UserService.FindAll(ctx, 1, 10)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func FindByNameUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	mockUserRepository.On("FindByName", "test").Return([]domain.User{}, int64(0), nil)

	_, err := UserService.FindByName(ctx, "test")

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	request	:= web.UserUpdateRequest{
		Name: "test",
		Username: "test",
		Email: "test@test.com",
		Phone: "123456789",
	}
	userId := "123456"

	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)
	mockUserRepository.On("Update", mock.AnythingOfType("*domain.User"),userId).Return(nil, nil)
	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)

	_, err := UserService.UpdateUser(ctx, request, userId)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)

}

func TestDeleteUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	userId := "123456"

	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)
	mockUserRepository.On("Delete", userId).Return(nil)

	err := UserService.DeleteUser(ctx, userId)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestResetPasswordUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	request	:= web.UserResetPasswordRequest{
		NewPassword: "testing1",
		ConfirmNewPassword: "testing1",
	}
	userId := "123456"

	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)
	mockUserRepository.On("ResetPassword", mock.AnythingOfType("*domain.User"),userId).Return(nil, nil)
	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)

	_, err := UserService.ResetPassword(ctx, request, userId)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestUpdatePhotoProfileUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	userId := "123456"

	mockUserRepository.On("FindById", userId).Return(&domain.User{}, nil)
	mockUserRepository.On("PhotoProfile", mock.AnythingOfType("*domain.User"),userId).Return(nil, nil)

	_, err := UserService.UpdatePhotoProfile(ctx, userId, "test")

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)
}

func TestLoginUser(t *testing.T){
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	UserService := &UserServiceImpl{
		UserRepository: mockUserRepository,
		Validate: validate,
	}

	request	:= web.UserLoginRequest{
		Username: "test",
		Password: "12345678",
	}

	mockUserRepository.On("FindByUsername", request.Username).Return(&domain.User{ID: "123456", Name: "test", Username: "test", Email: "test@test.com", Phone: "12345678910", Password: "$2a$10$ETTLBf7J9zl2WPBNjkriW.4mpISY2Z.8VJHXYAp8RkwYFTGFpj4au", ImageUrl: "test"}, nil)

	_, err := UserService.LoginUser(ctx, request)

	assert.NoError(t, err)

	mockUserRepository.AssertExpectations(t)

}
