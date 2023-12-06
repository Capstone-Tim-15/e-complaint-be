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

type CommentService interface {
	CreateComment(ctx echo.Context, request web.CommentCreateRequest) (*domain.Comment, error)
	CheckUser(senderId string) (*domain.User, error)
	CheckAdmin(senderId string) (*domain.Admin, error)
}

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	Validate          validator.Validate
}

func NewCommentService(CommentRepository repository.CommentRepository, Validate *validator.Validate) *CommentServiceImpl {
	return &CommentServiceImpl{
		CommentRepository: CommentRepository,
		Validate:          *Validate,
	}
}

func (s *CommentServiceImpl) CreateComment(ctx echo.Context, request web.CommentCreateRequest) (*domain.Comment, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	mess := req.CommentCreateRequesttoCommentDomain(request)

	result, err := s.CommentRepository.Create(mess)
	if err != nil {
		return nil, fmt.Errorf("error when creating message: %s", err.Error())
	}

	return result, nil
}

func (s *CommentServiceImpl) CheckUser(senderId string) (*domain.User, error) {
	result, err := s.CommentRepository.CheckUser(senderId)
	if err != nil {
		return nil, fmt.Errorf("error when check user: %s", err.Error())
	}

	return result, nil
}

func (s *CommentServiceImpl) CheckAdmin(senderId string) (*domain.Admin, error) {
	result, err := s.CommentRepository.CheckAdmin(senderId)
	if err != nil {
		return nil, fmt.Errorf("error when check admin: %s", err.Error())
	}

	return result, nil
}
