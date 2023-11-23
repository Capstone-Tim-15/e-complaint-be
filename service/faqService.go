package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/repository"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/req"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type FaqService interface {
	CreateFaq(ctx echo.Context, request web.FaqRequest) (*domain.FAQ, error)
	FindById(ctx echo.Context, id string) (*domain.FAQ, error)
	FindByCategory(ctx echo.Context, category string) (*domain.FAQ, error)
	FindAll(ctx echo.Context) ([]domain.FAQ, error)
	UpdateFaq(ctx echo.Context, request web.FaqUpdateRequest, id string) (*domain.FAQ, error)
}

type FaqServiceImpl struct {
	FaqRepository repository.FaqRepository
	Validate      *validator.Validate
}

func NewFaqService(FaqRepository repository.FaqRepository, validate *validator.Validate) *FaqServiceImpl {
	return &FaqServiceImpl{
		FaqRepository: FaqRepository,
		Validate:      validate,
	}
}

func (svc *FaqServiceImpl) CreateFaq(ctx echo.Context, request web.FaqRequest) (*domain.FAQ, error) {
	err := svc.Validate.Struct(request)
	if err != nil {
		return nil, err
	}
	faq := req.FAQRequestToFAQDomain(request)

	result, err := svc.FaqRepository.Create(faq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FaqServiceImpl) FindById(ctx echo.Context, id string) (*domain.FAQ, error) {
	result, err := svc.FaqRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FaqServiceImpl) FindByCategory(ctx echo.Context, category string) (*domain.FAQ, error) {
	result, err := svc.FaqRepository.FindByCategory(category)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FaqServiceImpl) FindAll(ctx echo.Context) ([]domain.FAQ, error) {
	result, err := svc.FaqRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (svc *FaqServiceImpl) UpdateFaq(ctx echo.Context, request web.FaqUpdateRequest, id string) (*domain.FAQ, error) {
	err := svc.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	ExistingFAQ, _ := svc.FaqRepository.FindById(id)
	if ExistingFAQ == nil {
		return nil, fmt.Errorf("FAQ not found")
	}

	faq := req.FAQUpdateToFAQDomain(request)

	_, err = svc.FaqRepository.Update(faq, id)
	if err != nil {
		return nil, err
	}

	result, err := svc.FaqRepository.FindById(id)
	return result, nil
}
