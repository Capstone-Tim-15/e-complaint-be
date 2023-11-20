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

type UserService interface {
	CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error)
	LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error)
	FindById(ctx echo.Context, id string) (*domain.User, error)
	FindAll(ctx echo.Context) ([]domain.User, error)
	FindByName(ctx echo.Context, name string) (*domain.User, error)
	UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id string) (*domain.User, error)
	ResetPassword(ctx echo.Context, request web.UserResetPasswordRequest, id string) (*domain.User, error)
	DeleteUser(ctx echo.Context, id string) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) CreateUser(ctx echo.Context, request web.UserCreateRequest) (*domain.User, error) {

	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUserName, _ := s.UserRepository.FindByUsername(request.Username)	
	if existingUserName != nil {
		return nil, fmt.Errorf("username already exist")
	}

	existingEmail, _ := s.UserRepository.FindByEmail(request.Email)
	if existingEmail != nil {
		return nil, fmt.Errorf("email already exist")
	}

	user := req.UserCreateRequestToUserDomain(request)

	user.Password = helper.HashPassword(user.Password)

	result, err := s.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error when creating user: %s", err.Error())
	}

	return result, nil
}

func (s *UserServiceImpl) LoginUser(ctx echo.Context, request web.UserLoginRequest) (*domain.User, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, err := s.UserRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	user := req.UserLoginRequestToUserDomain(request)

	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingUser, nil
}

func (s *UserServiceImpl) FindById(ctx echo.Context, id string) (*domain.User, error) {

	existingUser, err := s.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	return existingUser, nil
}

func (s *UserServiceImpl) FindAll(ctx echo.Context) ([]domain.User, error) {
	users, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("users not found")
	}

	return users, nil
}

func (r *UserServiceImpl) FindByName(ctx echo.Context, name string) (*domain.User, error) {
	user, _ := r.UserRepository.FindByName(name)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *UserServiceImpl) UpdateUser(ctx echo.Context, request web.UserUpdateRequest, id string) (*domain.User, error) {

	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	user := req.UserUpdateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	_, err = s.UserRepository.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}
	result, err := s.UserRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	return result, nil
}

func (s *UserServiceImpl) ResetPassword(ctx echo.Context, request web.UserResetPasswordRequest, id string) (*domain.User, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	existingUser, _ := s.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	if request.NewPassword != request.ConfirmNewPassword {
		return nil, fmt.Errorf("new password and confirm new password do not match")
	}

	user := req.UserResetPasswordRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	_, err = s.UserRepository.ResetPassword(user, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	result, err := s.UserRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	return result, nil

}

func (s *UserServiceImpl) DeleteUser(ctx echo.Context, id string) error {

	existingUser, _ := s.UserRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	err := s.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting user: %s", err)
	}

	return nil
}
