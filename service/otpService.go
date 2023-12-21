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
	CreateOTPUser(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTPUser, error)
	CheckOTPUser(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTPUser, error)
	CreateOTPAdmin(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTPAdmin, error)
	CheckOTPAdmin(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTPAdmin, error)
}

type OTPServiceImpl struct {
	OTPRepository   repository.OTPRepository
	UserRepository  repository.UserRepository
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewOTPService(OTPRepository repository.OTPRepository, UserRepository repository.UserRepository, AdminRepository repository.AdminRepository, Validate *validator.Validate) *OTPServiceImpl {
	return &OTPServiceImpl{
		OTPRepository:   OTPRepository,
		UserRepository:  UserRepository,
		AdminRepository: AdminRepository,
		Validate:        Validate,
	}
}

func (s *OTPServiceImpl) CreateOTPUser(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTPUser, error) {
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
		fmt.Println(err)
		return nil, fmt.Errorf("send otp error")
	}

	otp := res.OTPCreateRequesttoOTPDomain(user.ID, otpKey)
	result, err := s.OTPRepository.CreateOTPUser(otp)
	if err != nil {
		return nil, fmt.Errorf("error when creating user: %s", err.Error())
	}

	return result, nil
}

func (s *OTPServiceImpl) CheckOTPUser(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTPUser, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	result, err := s.OTPRepository.FindByUserId(id)
	if err != nil {
		return nil, fmt.Errorf("error when retrieving OTP: %s", err.Error())
	}

	if result != nil && result.OTP != request.OTP {
		return nil, fmt.Errorf("otp do not match")
	}

	if result != nil && result.OTP == request.OTP {
		s.OTPRepository.DeleteOTPUser(result.ID)

		return result, nil
	}

	return result, fmt.Errorf("invalid OTP")
}

func (s *OTPServiceImpl) CreateOTPAdmin(ctx echo.Context, request web.OTPCreateRequest) (*domain.OTPAdmin, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	admin, _ := s.AdminRepository.FindByEmail(request.Email)
	if admin == nil {
		return nil, fmt.Errorf("email not found")
	}

	otpKey := helper.GenerateOTP()

	err = helper.SendOTP(admin.Email, otpKey)
	if err != nil {
		return nil, fmt.Errorf("send otp error")
	}

	otp := res.AdminOTPCreateRequesttoOTPDomain(admin.ID, otpKey)
	result, err := s.OTPRepository.CreateOTPAdmin(otp)
	if err != nil {
		return nil, fmt.Errorf("error when creating otp admin: %s", err.Error())
	}

	return result, nil
}

func (s *OTPServiceImpl) CheckOTPAdmin(ctx echo.Context, request web.OTPCheckRequest, id string) (*domain.OTPAdmin, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	result, err := s.OTPRepository.FindByAdminId(id)
	if err != nil {
		return nil, fmt.Errorf("error when retrieving OTP: %s", err.Error())
	}

	if result != nil && result.OTP != request.OTP {
		return nil, fmt.Errorf("otp do not match")
	}

	if result != nil && result.OTP == request.OTP {
		s.OTPRepository.DeleteOTPAdmin(result.ID)

		return result, nil
	}

	return result, fmt.Errorf("invalid OTP")
}
