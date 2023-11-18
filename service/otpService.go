package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type OTPService interface {
	CreateOTP(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTP, error)
	CheckOTP(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTP, error)
}

type OTPServiceImpl struct {
	OTPRepository  repository.OTPRepository
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewOTPService(OTPRepository repository.OTPRepository, UserRepository repository.UserRepository, Validate *validator.Validate) *OTPServiceImpl {
	return &OTPServiceImpl{
		OTPRepository:  OTPRepository,
		UserRepository: UserRepository,
		Validate:       Validate,
	}
}

func (s *OTPServiceImpl) CreateOTP(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTP, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	user, _ := s.UserRepository.FindByEmail(request.Email)
	if user == nil {
		return nil, fmt.Errorf("email not found")
	}

	otpKey := helper.GenerateOTP()

	err = helper.SendOTP(user.Email, otpKey)
	if err != nil {
		return nil, fmt.Errorf("send otp error")
	}

	otp := res.OTPCreateRequesttoOTPDomain(user.ID, otpKey)
	result, err := s.OTPRepository.Create(otp)
	if err != nil {
		return nil, fmt.Errorf("error when creating user: %s", err.Error())
	}

	return result, nil
}

func (s *OTPServiceImpl) CheckOTP(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTP, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	result, err := s.OTPRepository.FindByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("error when retrieving OTP: %s", err.Error())
	}

	if result != nil && result.OTP == request.OTP {
		s.OTPRepository.Delete(result.ID)

		return result, nil
	}

	return result, fmt.Errorf("invalid OTP")
}
