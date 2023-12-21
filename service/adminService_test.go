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

func TestCreateAdmin(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}

	request	:= web.AdminCreateRequest{
		Name: "test",
		Username: "test",
		Email: "test@test.com",
		Phone: "123456789",
		Password: "testing1",
	}

	mockAdminRepository.On("FindByEmail", "test@test.com").Return(nil, nil)
	mockAdminRepository.On("FindByUsername", request.Username).Return(nil, nil)
	mockAdminRepository.On("Create", mock.AnythingOfType("*domain.Admin")).Return(nil, nil)

	_, err := AdminService.CreateAdmin(ctx, request)

	assert.NoError(t, err)
	
	mockAdminRepository.AssertExpectations(t)
}

func TestFindByID(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}
	userId := "123456"

	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)

	_, err := AdminService.FindById(ctx, userId)

	assert.NoError(t, err)
	
	mockAdminRepository.AssertExpectations(t)
}

func TestFindAll(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}

	mockAdminRepository.On("FindAll", 1, 10).Return([]domain.Admin{}, int64(0), nil)

	_, _, err := AdminService.FindAll(ctx, 1, 10)

	assert.NoError(t, err)
	
	mockAdminRepository.AssertExpectations(t)
}

func TestFindByName(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}
	name := "test"

	mockAdminRepository.On("FindByName", name).Return(&domain.Admin{}, nil)

	_, err := AdminService.FindByName(ctx, name)

	assert.NoError(t, err)
	
	mockAdminRepository.AssertExpectations(t)
}

func TestUpdateAdmin(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}

	request	:= web.AdminUpdateRequest{
		Name: "test",
		Username: "test",
		Email: "test@test.test",	
		Phone: "123456789",
		Password: "testing1",
	}
	userId := "123456"

	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("Update", mock.AnythingOfType("*domain.Admin"),userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)

	_, err := AdminService.UpdateAdmin(ctx, request, userId)

	assert.NoError(t, err)

	mockAdminRepository.AssertExpectations(t)
}

func TestDeleteAdmin(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}
	userId := "123456"

	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("Delete", userId).Return(nil)

	err := AdminService.DeleteAdmin(ctx, userId)

	assert.NoError(t, err)

	mockAdminRepository.AssertExpectations(t)
}

func TestResetPassword(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}

	request	:= web.AdminResetPasswordRequest{
		NewPassword: "testing1",
		ConfirmNewPassword: "testing1",
	}
	userId := "123456"

	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("ResetPassword", mock.AnythingOfType("*domain.Admin"),userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)

	_, err := AdminService.ResetPassword(ctx, request, userId)

	assert.NoError(t, err)

	mockAdminRepository.AssertExpectations(t)
}

func TestUpdatePhotoProfile(t *testing.T){
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	AdminService := &AdminServiceImpl{
		AdminRepository: mockAdminRepository,
		Validate: validate,
	}
	userId := "123456"
	imageUrl := "test"

	mockAdminRepository.On("FindById", userId).Return(&domain.Admin{}, nil)
	mockAdminRepository.On("PhotoProfile", mock.AnythingOfType("*domain.Admin"),userId).Return(&domain.Admin{}, nil)

	_, err := AdminService.UpdatePhotoProfile(ctx, userId, imageUrl)

	assert.NoError(t, err)

	mockAdminRepository.AssertExpectations(t)
}