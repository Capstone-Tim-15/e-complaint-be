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

type FeedbackService interface {
	CreateFeedback(ctx echo.Context, request web.FeedbackCreateRequest) (*domain.Feedback, error)
	UpdateFeedback(ctx echo.Context, request web.FeedbackUpdateRequest, id string) (*domain.Feedback, error)
	DeleteFeedback(ctx echo.Context, id string) error
	FindById(ctx echo.Context, id string) (*domain.Feedback, error)
	FindByAll(ctx echo.Context, page, pageSize int) ([]domain.Feedback, int64, error)
	FindByNewsId(ctx echo.Context, newsID string, page, pageSize int) ([]domain.Feedback, int64, error)
	CheckUser(senderId string) (*domain.User, error)
	CheckAdmin(senderId string) (*domain.Admin, error)
}

type FeedbackServiceImp struct {
	FeedbackRepository repository.FeedbackRepository
	validate           *validator.Validate
}

func NewFeedbackService(FeedbackRepository repository.FeedbackRepository, validate *validator.Validate) *FeedbackServiceImp {
	return &FeedbackServiceImp{
		FeedbackRepository: FeedbackRepository,
		validate:           validate,
	}
}

func (fs *FeedbackServiceImp) CreateFeedback(ctx echo.Context, request web.FeedbackCreateRequest) (*domain.Feedback, error) {
	err := fs.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	feedback := req.FeedbackCreateRequestToFeedbackDomain(request)

	result, err := fs.FeedbackRepository.Create(feedback)
	if err != nil {
		return nil, fmt.Errorf("error when creating Feedback: %s ", err.Error())
	}

	return result, nil
}

func (fs *FeedbackServiceImp) UpdateFeedback(ctx echo.Context, request web.FeedbackUpdateRequest, id string) (*domain.Feedback, error) {
	err := fs.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingFeedback, _ := fs.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return nil, fmt.Errorf("feedback not found")
	}
	feedback := req.FeedbackUpdateRequestToFeedbackDomain(request)
	_, err = fs.FeedbackRepository.Update(feedback, id)
	if err != nil {
		return nil, fmt.Errorf("error when updatating feedback: %s", err.Error())
	}
	result, err := fs.FeedbackRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating feedback: %s", err.Error())
	}
	return result, nil
}

func (fs *FeedbackServiceImp) DeleteFeedback(ctx echo.Context, id string) error {
	existingFeedback, _ := fs.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return fmt.Errorf("feedback not Found")
	}
	err := fs.FeedbackRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting Feedback: %s", err)
	}
	return nil
}

func (fs *FeedbackServiceImp) FindById(ctx echo.Context, id string) (*domain.Feedback, error) {
	existingFeedback, _ := fs.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return nil, fmt.Errorf("feedback not found")
	}
	return existingFeedback, nil
}

func (fs *FeedbackServiceImp) FindByAll(ctx echo.Context, page, pageSize int) ([]domain.Feedback, int64, error) {
	feedback, totalCount, err := fs.FeedbackRepository.FindByAll(page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("feedback not found")
	}
	return feedback, totalCount, nil
}

func (fs *FeedbackServiceImp) FindByNewsId(ctx echo.Context, newsID string, page, pageSize int) ([]domain.Feedback, int64, error) {
	feedback, totalCount, err := fs.FeedbackRepository.FindByNewsId(newsID, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("feedback not found")
	}
	return feedback, totalCount, nil
}

func (fs *FeedbackServiceImp) CheckUser(senderId string) (*domain.User, error) {
	result, err := fs.FeedbackRepository.CheckUser(senderId)
	if err != nil {
		return nil, fmt.Errorf("error when check user: %s", err.Error())
	}

	return result, nil
}

func (fs *FeedbackServiceImp) CheckAdmin(senderId string) (*domain.Admin, error) {
	result, err := fs.FeedbackRepository.CheckAdmin(senderId)
	if err != nil {
		return nil, fmt.Errorf("error when check admin: %s", err.Error())
	}

	return result, nil
}
