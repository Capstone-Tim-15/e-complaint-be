package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"fmt"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestCheckOTPUser(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:  mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:       validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByUserId", "123456").Return(&domain.OTPUser{ID: "123456", User_ID: "123456", Admin_ID: "", OTP: "123456"}, nil)
	mockOTPRepository.On("DeleteOTPUser", "123456").Return(nil)

	_, err := OTPService.CheckOTPUser(ctx, request, "123456")

	assert.NoError(t, err)

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPUserFailValidate(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:  mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:       validate,
	}

	request := web.OTPCheckRequest{
		OTP: "12345",
	}

	_, err := OTPService.CheckOTPUser(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("validation failed: validation error on field OTP, tag min"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPUserFailGet(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:  mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:       validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByUserId", "123456").Return(&domain.OTPUser{ID: "123456", User_ID: "123456", Admin_ID: "", OTP: "123456"}, fmt.Errorf("error"))

	_, err := OTPService.CheckOTPUser(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when retrieving OTP: error"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPUserFailNoMatch(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:  mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:       validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByUserId", "123456").Return(&domain.OTPUser{ID: "123456", User_ID: "123456", Admin_ID: "", OTP: "123457"}, nil)

	_, err := OTPService.CheckOTPUser(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("otp do not match"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPUserFail(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:  mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:       validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByUserId", "123456").Return(nil, nil)

	_, err := OTPService.CheckOTPUser(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("invalid OTP"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPAdmin(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByAdminId", "123456").Return(&domain.OTPAdmin{ID: "123456", User_ID: "", Admin_ID: "123456", OTP: "123456"}, nil)
	mockOTPRepository.On("DeleteOTPAdmin", "123456").Return(nil)

	_, err := OTPService.CheckOTPAdmin(ctx, request, "123456")

	assert.NoError(t, err)

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPAdminFailValidate(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCheckRequest{
		OTP: "12345",
	}

	_, err := OTPService.CheckOTPAdmin(ctx, request, "123456")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed: validation error on field OTP, tag min")

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPAdminFailGet(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByAdminId", "123456").Return(&domain.OTPAdmin{ID: "123456", User_ID: "", Admin_ID: "123456", OTP: "123456"}, fmt.Errorf("error"))

	_, err := OTPService.CheckOTPAdmin(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when retrieving OTP: error"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPAdminFailNoMatch(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByAdminId", "123456").Return(&domain.OTPAdmin{ID: "123456", User_ID: "", Admin_ID: "123456", OTP: "123457"}, nil)

	_, err := OTPService.CheckOTPAdmin(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("otp do not match"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCheckOTPAdminFail(t *testing.T) {
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCheckRequest{
		OTP: "123456",
	}

	mockOTPRepository.On("FindByAdminId", "123456").Return(nil, nil)

	_, err := OTPService.CheckOTPAdmin(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("invalid OTP"))

	mockOTPRepository.AssertExpectations(t)

}

func TestCreateOTPUserFailSend(t *testing.T){
	mockOTPRepository := new(mocks.OTPRepository)
	mockUserRepository := new(mocks.UserRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		UserRepository: mockUserRepository,
		Validate:        validate,
	}

	request := web.OTPCreateRequest{
		Email: "innakadyleexd@gmail.com",
	}

	mockUserRepository.On("FindByEmail", request.Email).Return(&domain.User{ID: "123456", Name: "test", Username: "test", Email: "innakadylexd@upi.edu", Phone: "12345678910", Password: "testing1", ImageUrl: "test.png"}, nil)

	_, err := OTPService.CreateOTPUser(ctx, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "send otp error")

	mockOTPRepository.AssertExpectations(t)
}

func TestCreateOTPAdminFailSend(t *testing.T){
	mockOTPRepository := new(mocks.OTPRepository)
	mockAdminRepository := new(mocks.AdminRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	OTPService := &OTPServiceImpl{
		OTPRepository:   mockOTPRepository,
		AdminRepository: mockAdminRepository,
		Validate:        validate,
	}

	request := web.OTPCreateRequest{
		Email: "innakadyleexd@gmail.com",
	}

	mockAdminRepository.On("FindByEmail", request.Email).Return(&domain.Admin{ID: "123456", Name: "test", Username: "test", Email: "innakadylexd@upi.edu", Phone: "12345678910", Password: "testing1", ImageUrl: "test.png"}, nil)

	_, err := OTPService.CreateOTPAdmin(ctx, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "send otp error")

	mockOTPRepository.AssertExpectations(t)
}