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

func TestCreateFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	request	:= web.FeedbackCreateRequest{
		Fullname: "test",
		Role: "test",
		PhotoImage: "test.png",
		News_ID: "123456",
		Content: "test",
	}

	mockFeedbackRepository.On("Create", mock.AnythingOfType("*domain.Feedback")).Return(nil, nil)

	_, err := FeedbackService.CreateFeedback(ctx, request)

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestUpdateFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	request	:= web.FeedbackUpdateRequest{
		News_ID: "123456",
		Content: "test",
	}

	mockFeedbackRepository.On("FindById", mock.AnythingOfType("string")).Return(&domain.Feedback{}, nil)
	mockFeedbackRepository.On("Update", mock.AnythingOfType("*domain.Feedback"), mock.AnythingOfType("string")).Return(nil, nil)

	_, err := FeedbackService.UpdateFeedback(ctx, request, "123456")

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestDeleteFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("FindById", mock.AnythingOfType("string")).Return(&domain.Feedback{}, nil)
	mockFeedbackRepository.On("Delete", mock.AnythingOfType("string")).Return(nil)

	err := FeedbackService.DeleteFeedback(ctx, "123456")

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestFindByIdFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("FindById", mock.AnythingOfType("string")).Return(&domain.Feedback{}, nil)

	_, err := FeedbackService.FindById(ctx, "123456")

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestFindAllFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("FindByAll",1,10).Return([]domain.Feedback{}, int64(0), nil)

	_, _, err := FeedbackService.FindByAll(ctx,1,10)

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestFindByNewsIdFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("FindByNewsId", mock.AnythingOfType("string"),1,10).Return([]domain.Feedback{}, int64(0), nil)

	_, _, err := FeedbackService.FindByNewsId(ctx,"123456",1,10)

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestCheckUserFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("CheckUser", mock.AnythingOfType("string")).Return(nil, nil)

	_, err := FeedbackService.CheckUser("123456")

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}

func TestCheckAdminFeedback(t *testing.T){
	mockFeedbackRepository := new(mocks.FeedbackRepository)
	validate := validator.New()

	FeedbackService := &FeedbackServiceImp{
		FeedbackRepository: mockFeedbackRepository,
		validate: validate,
	}

	mockFeedbackRepository.On("CheckAdmin", mock.AnythingOfType("string")).Return(nil, nil)

	_, err := FeedbackService.CheckAdmin("123456")

	assert.NoError(t, err)

	mockFeedbackRepository.AssertExpectations(t)
}
