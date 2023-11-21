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
}

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	Validate          validator.Validate
}

func NewMessService(CommentRepository repository.CommentRepository, Validate *validator.Validate) *CommentServiceImpl {
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

	mess := req.MessCreateRequesttoMessDomain(request)

	result, err := s.CommentRepository.Create(mess)
	if err != nil {
		return nil, fmt.Errorf("error when creating message: %s", err.Error())
	}

	return result, nil
}
