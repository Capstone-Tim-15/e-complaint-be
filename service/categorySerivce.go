package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/req"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryService interface {
	CreateCategory(ctx echo.Context, request web.CategoryRequest) (*domain.Category, error)
	FindById(ctx echo.Context, id string) (*domain.Category, error)
	FindAll(ctx echo.Context) ([]domain.Category, error)
	UpdateCategory(ctx echo.Context, request web.CategoryRequest, id string) (*domain.Category, error)
	Pagination(offset int, limit int) ([]domain.Category, *web.Pagination, error)
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: repo,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx echo.Context, request web.CategoryRequest) (*domain.Category, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	cat := req.CategoryRequestToCategoryDomain(request)

	result, err := service.CategoryRepository.Create(cat)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *CategoryServiceImpl) FindById(ctx echo.Context, id string) (*domain.Category, error) {
	result, err := svc.CategoryRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *CategoryServiceImpl) FindAll(ctx echo.Context) ([]domain.Category, error) {
	result, err := svc.CategoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *CategoryServiceImpl) UpdateCategory(ctx echo.Context, request web.CategoryRequest, id string) (*domain.Category, error) {
	err := svc.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	cat := req.CategoryRequestToCategoryDomain(request)

	_, err = svc.CategoryRepository.Update(cat, id)
	if err != nil {
		return nil, err
	}
	result, err := svc.CategoryRepository.FindById(id)

	return result, nil
}
