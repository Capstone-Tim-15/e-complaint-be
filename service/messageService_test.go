package service

import (
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"fmt"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateComment(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}

	request	:= web.CommentCreateRequest{
		Complaint_ID: "123456",
		Fullname: "test",
		Role: "user",
		Message: "test",
	}

	mockCommentRepository.On("Create", mock.AnythingOfType("*domain.Comment")).Return(nil, nil)

	_, err := CommentService.CreateComment(ctx, request)

	assert.NoError(t, err)

	mockCommentRepository.AssertExpectations(t)
}

func TestCreateCommentFailedValidate(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}

	request	:= web.CommentCreateRequest{
		Fullname: "test",
		Role: "user",
		Message: "test",
	}

	_, err := CommentService.CreateComment(ctx, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed: validation error on field Complaint_ID, tag required")

	mockCommentRepository.AssertExpectations(t)
}

func TestCreateCommentFailCreating(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}

	request	:= web.CommentCreateRequest{
		Complaint_ID: "123456",
		Fullname: "test",
		Role: "user",
		Message: "test",
	}

	mockCommentRepository.On("Create", mock.AnythingOfType("*domain.Comment")).Return(nil, fmt.Errorf("error creating comment"))

	_, err := CommentService.CreateComment(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when creating message: error creating comment"))

	mockCommentRepository.AssertExpectations(t)
}

func TestCheckUser(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}
	senderId := "123456"

	mockCommentRepository.On("CheckUser", senderId).Return(nil, nil)

	_, err := CommentService.CheckUser(senderId)

	assert.NoError(t, err)

	mockCommentRepository.AssertExpectations(t)
}

func TestCheckUserFailed(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}
	senderId := "123456"

	mockCommentRepository.On("CheckUser", senderId).Return(nil, fmt.Errorf("error check user"))

	_, err := CommentService.CheckUser(senderId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when check user: error check user"))

	mockCommentRepository.AssertExpectations(t)
}

func TestCheckAdmin(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}
	senderId := "123456"

	mockCommentRepository.On("CheckAdmin", senderId).Return(nil, nil)

	_, err := CommentService.CheckAdmin(senderId)

	assert.NoError(t, err)

	mockCommentRepository.AssertExpectations(t)
}

func TestCheckAdminFailed(t *testing.T){
	mockCommentRepository := new(mocks.CommentRepository)
	validate := validator.New()

	CommentService := &CommentServiceImpl{
		CommentRepository: mockCommentRepository,
		Validate: *validate,
	}
	senderId := "123456"

	mockCommentRepository.On("CheckAdmin", senderId).Return(nil, fmt.Errorf("error check admin"))

	_, err := CommentService.CheckAdmin(senderId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when check admin: error check admin"))

	mockCommentRepository.AssertExpectations(t)
}