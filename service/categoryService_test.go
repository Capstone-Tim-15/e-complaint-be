package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCategory(t *testing.T){
	mockCategoryRepository := new(mocks.CategoryRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CategoryService := &CategoryServiceImpl{
		CategoryRepository: mockCategoryRepository,
		Validate: validate,
	}

	request	:= web.CategoryRequest{
		Name: "test",
	}

	mockCategoryRepository.On("Create", mock.AnythingOfType("*domain.Category")).Return(nil, nil)

	_, err := CategoryService.CreateCategory(ctx, request)

	assert.NoError(t, err)

	mockCategoryRepository.AssertExpectations(t)
}

func TestFindByIdCategory(t *testing.T){
	mockCategoryRepository := new(mocks.CategoryRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CategoryService := &CategoryServiceImpl{
		CategoryRepository: mockCategoryRepository,
		Validate: validate,
	}
	categoryId := "123456"

	mockCategoryRepository.On("FindById", categoryId).Return(&domain.Category{}, nil)

	_, err := CategoryService.FindById(ctx, categoryId)

	assert.NoError(t, err)

	mockCategoryRepository.AssertExpectations(t)
}

func TestFindAllCategory(t *testing.T){
	mockCategoryRepository := new(mocks.CategoryRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CategoryService := &CategoryServiceImpl{
		CategoryRepository: mockCategoryRepository,
		Validate: validate,
	}

	mockCategoryRepository.On("FindAll").Return([]domain.Category{}, nil)

	_, err := CategoryService.FindAll(ctx)

	assert.NoError(t, err)

	mockCategoryRepository.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T){
	mockCategoryRepository := new(mocks.CategoryRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CategoryService := &CategoryServiceImpl{
		CategoryRepository: mockCategoryRepository,
		Validate: validate,
	}

	categoryId := "123456"
	request	:= web.CategoryRequest{
		Name: "test",
	}

	mockCategoryRepository.On("FindById", categoryId).Return(&domain.Category{}, nil)
	mockCategoryRepository.On("Update", mock.AnythingOfType("*domain.Category"), categoryId).Return(nil, nil)

	_, err := CategoryService.UpdateCategory(ctx, request, categoryId)

	assert.NoError(t, err)

	mockCategoryRepository.AssertExpectations(t)
}

