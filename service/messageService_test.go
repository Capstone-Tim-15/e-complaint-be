package service

import (
	"ecomplaint/test/mocks"
	"ecomplaint/model/web"
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