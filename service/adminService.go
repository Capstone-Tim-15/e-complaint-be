package service

import (
	"fmt"
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/req"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminService interface {
	CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error)
	UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error)
	ResetPassword(ctx echo.Context, request web.AdminResetPasswordRequest) (*domain.Admin, error)
	FindById(ctx echo.Context, id int) (*domain.Admin, error)
	FindAll(ctx echo.Context) ([]domain.Admin, error)
	FindByName(ctx echo.Context, name string) (*domain.Admin, error)
	DeleteAdmin(ctx echo.Context, id int) error
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

func (context *AdminServiceImpl) CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("email already exist")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)

	admin.Password = helper.HashPassword(admin.Password)

	result, err := context.AdminRepository.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating Admin: %s", err.Error())
	}

	return result, nil
}

func (context *AdminServiceImpl) LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, err := context.AdminRepository.FindByEmail(request.Email)
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

func (context *AdminServiceImpl) FindById(ctx echo.Context, id int) (*domain.Admin, error) {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return existingAdmin, nil
}

func (context *AdminServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Admin, error) {
	admin, _ := context.AdminRepository.FindByName(name)
	if admin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return admin, nil
}

func (context *AdminServiceImpl) FindAll(ctx echo.Context) ([]domain.Admin, error) {
	admin, err := context.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}

	return admin, nil
}

func (context *AdminServiceImpl) UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error) {

	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	_, err = context.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	result, err := context.AdminRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	return result, nil
}

func (context *AdminServiceImpl) ResetPassword(ctx echo.Context, request web.AdminResetPasswordRequest) (*domain.Admin, error) {
	err := context.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingAdmin, _ := context.AdminRepository.FindByEmail(request.Email)
	if existingAdmin == nil {
		return nil, fmt.Errorf("user not found")
	}

	if request.NewPassword != request.ConfirmNewPassword {
		return nil, fmt.Errorf("new password and confirm new password do not match")
	}

	user := req.AdminResetPasswordRequestToAdminDomain(request)
	user.Password = helper.HashPassword(user.Password)

	_, err = context.AdminRepository.ResetPassword(user, request.Email)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	result, err := context.AdminRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	return result, nil

}

func (context *AdminServiceImpl) DeleteAdmin(ctx echo.Context, id int) error {

	existingAdmin, _ := context.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}

	err := context.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting admin: %s", err)
	}

	return nil
}
