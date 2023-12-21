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

func TestCreateLike(t *testing.T){
	mockLikeRepository := new(mocks.LikeRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	LikeService := &LikeServiceImp{
		LikesRepository: mockLikeRepository,
		validate: validate,
	}

	request	:= web.LikesCreateRequest{
		User_ID: "123456",
		News_ID: "123456",
		Status: "like",
	}

	mockLikeRepository.On("Create", mock.AnythingOfType("*domain.Likes")).Return(nil, nil)

	_, err := LikeService.CreateLike(ctx, request)

	assert.NoError(t, err)

	mockLikeRepository.AssertExpectations(t)
}

func TestUpdateLike(t *testing.T){
	mockLikeRepository := new(mocks.LikeRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	LikeService := &LikeServiceImp{
		LikesRepository: mockLikeRepository,
		validate: validate,
	}

	request	:= web.LikesUpdateRequest{
		Status: "unlike",
	}

	likeId := "123456"

	mockLikeRepository.On("FindById", likeId).Return(&domain.Likes{}, nil)
	mockLikeRepository.On("Update", mock.AnythingOfType("*domain.Likes"), likeId).Return(nil, nil)

	_, err := LikeService.UpdateLike(ctx, request, likeId)

	assert.NoError(t, err)

	mockLikeRepository.AssertExpectations(t)
}

func TestDeleteLike(t *testing.T){
	mockLikeRepository := new(mocks.LikeRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	LikeService := &LikeServiceImp{
		LikesRepository: mockLikeRepository,
		validate: validate,
	}

	likeId := "123456"

	mockLikeRepository.On("FindById", likeId).Return(&domain.Likes{}, nil)
	mockLikeRepository.On("Delete", likeId).Return(nil)

	err := LikeService.DeleteLike(ctx, likeId)

	assert.NoError(t, err)

	mockLikeRepository.AssertExpectations(t)
}

func TestFindByIdLike(t *testing.T){
	mockLikeRepository := new(mocks.LikeRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	LikeService := &LikeServiceImp{
		LikesRepository: mockLikeRepository,
		validate: validate,
	}

	likeId := "123456"

	mockLikeRepository.On("FindById", likeId).Return(&domain.Likes{}, nil)

	_, err := LikeService.FindById(ctx, likeId)

	assert.NoError(t, err)

	mockLikeRepository.AssertExpectations(t)
}

func TestFindByAllLike(t *testing.T){
	mockLikeRepository := new(mocks.LikeRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	LikeService := &LikeServiceImp{
		LikesRepository: mockLikeRepository,
		validate: validate,
	}

	mockLikeRepository.On("FindByAll").Return([]domain.Likes{}, nil)

	_, err := LikeService.FindByAll(ctx)

	assert.NoError(t, err)

	mockLikeRepository.AssertExpectations(t)
}

