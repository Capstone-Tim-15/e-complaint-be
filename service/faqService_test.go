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

func TestCreateFaq(t *testing.T){
	mockFaqRepository := new(mocks.FaqRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FaqService := &FaqServiceImpl{
		FaqRepository: mockFaqRepository,
		Validate: validate,
	}

	request	:= web.FaqRequest{
		Title: "test",
		Content: "test",
		Category_ID: "123456",
	}

	mockFaqRepository.On("Create", mock.AnythingOfType("*domain.FAQ")).Return(nil, nil)

	_, err := FaqService.CreateFaq(ctx, request)

	assert.NoError(t, err)

	mockFaqRepository.AssertExpectations(t)

}

func TestFindByIdFaq(t *testing.T){
	mockFaqRepository := new(mocks.FaqRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FaqService := &FaqServiceImpl{
		FaqRepository: mockFaqRepository,
		Validate: validate,
	}
	faqId := "123456"

	mockFaqRepository.On("FindById", faqId).Return(&domain.FAQ{}, nil)

	_, err := FaqService.FindById(ctx, faqId)

	assert.NoError(t, err)

	mockFaqRepository.AssertExpectations(t)
}

func TestFindByCategoryFaq(t *testing.T){
	mockFaqRepository := new(mocks.FaqRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FaqService := &FaqServiceImpl{
		FaqRepository: mockFaqRepository,
		Validate: validate,
	}
	category := "test"

	mockFaqRepository.On("FindByCategory", category).Return(&domain.FAQ{}, nil)

	_, err := FaqService.FindByCategory(ctx, category)

	assert.NoError(t, err)

	mockFaqRepository.AssertExpectations(t)
}

func TestFindAllFaq(t *testing.T){
	mockFaqRepository := new(mocks.FaqRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FaqService := &FaqServiceImpl{
		FaqRepository: mockFaqRepository,
		Validate: validate,
	}

	mockFaqRepository.On("FindAll").Return([]domain.FAQ{}, nil)

	_, err := FaqService.FindAll(ctx)

	assert.NoError(t, err)

	mockFaqRepository.AssertExpectations(t)
}

func TestUpdateFaq(t *testing.T){
	mockFaqRepository := new(mocks.FaqRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FaqService := &FaqServiceImpl{
		FaqRepository: mockFaqRepository,
		Validate: validate,
	}

	request	:= web.FaqUpdateRequest{
		Title: "test",
		Content: "test",
		Category_ID: "123456",
	}

	mockFaqRepository.On("FindById", mock.AnythingOfType("string")).Return(&domain.FAQ{}, nil)
	mockFaqRepository.On("Update", mock.AnythingOfType("*domain.FAQ"), mock.AnythingOfType("string")).Return(nil, nil)

	_, err := FaqService.UpdateFaq(ctx, request, "123456")

	assert.NoError(t, err)

	mockFaqRepository.AssertExpectations(t)
}