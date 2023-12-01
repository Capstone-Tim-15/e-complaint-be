package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminService interface {
	CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error)
	FindById(ctx echo.Context, id string) (*domain.Admin, error)
	FindAll(ctx echo.Context, page, pageSize int) ([]domain.Admin, int64, error)
	FindByName(ctx echo.Context, name string) (*domain.Admin, error)
	UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id string) (*domain.Admin, error)
	ResetPassword(ctx echo.Context, request web.AdminResetPasswordRequest, id string) (*domain.Admin, error)
	UpdatePhotoProfile(ctx echo.Context, id string, imageUrl string) (*domain.Admin, error)
	DeleteAdmin(ctx echo.Context, id string) error
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminService(AdminRepository repository.AdminRepository, validate *validator.Validate) *AdminServiceImpl {
	return &AdminServiceImpl{
		AdminRepository: AdminRepository,
		Validate:        validate,
	}
}

func (s *AdminServiceImpl) CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {

	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUsername, _ := s.AdminRepository.FindByEmail(request.Email)
	if existingUsername != nil {
		return nil, fmt.Errorf("email already exist")
	}

	existingEmail, _ := s.AdminRepository.FindByUsername(request.Username)
	if existingEmail != nil {
		return nil, fmt.Errorf("username already exist")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)

	admin.Password = helper.HashPassword(admin.Password)

	result, err := s.AdminRepository.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating Admin: %s", err.Error())
	}

	return result, nil
}

func (s *AdminServiceImpl) LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, err := s.AdminRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	admin := req.AdminLoginRequestToAdminDomain(request)

	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingAdmin, nil
}

func (s *AdminServiceImpl) FindById(ctx echo.Context, id string) (*domain.Admin, error) {

	existingAdmin, _ := s.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return existingAdmin, nil
}

func (s *AdminServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Admin, error) {
	admin, _ := s.AdminRepository.FindByName(name)
	if admin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return admin, nil
}

func (s *AdminServiceImpl) FindAll(ctx echo.Context, page, pageSize int) ([]domain.Admin, int64, error) {
	admins, totalCount, err := s.AdminRepository.FindAll(page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("admins not found")
	}

	return admins, totalCount, nil
}

func (s *AdminServiceImpl) UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id string) (*domain.Admin, error) {

	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := s.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	_, err = s.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	result, err := s.AdminRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	return result, nil
}

func (s *AdminServiceImpl) ResetPassword(ctx echo.Context, request web.AdminResetPasswordRequest, id string) (*domain.Admin, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := s.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	if request.NewPassword != request.ConfirmNewPassword {
		return nil, fmt.Errorf("new password and confirm new password do not match")
	}

	admin := req.AdminResetPasswordRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	_, err = s.AdminRepository.ResetPassword(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	result, err := s.AdminRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	return result, nil

}

func (s *AdminServiceImpl) UpdatePhotoProfile(ctx echo.Context, id string, imageUrl string) (*domain.Admin, error) {
	existingAdmin, _ := s.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	existingAdmin.ImageUrl = imageUrl

	_, err := s.AdminRepository.PhotoProfile(existingAdmin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	return existingAdmin, nil
}

func (s *AdminServiceImpl) DeleteAdmin(ctx echo.Context, id string) error {

	existingAdmin, _ := s.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}

	err := s.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting admin: %s", err)
	}

	return nil
}
