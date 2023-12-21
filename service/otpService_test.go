package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
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
