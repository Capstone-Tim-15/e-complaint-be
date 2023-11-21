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

type ComplaintService interface {
	CreateComplaint(ctx echo.Context, request web.ComplaintCreateRequest) (*domain.Complaint, error)
	FindById(ctx echo.Context, id string) (*domain.Complaint, error)
	FindAll(ctx echo.Context) ([]domain.Complaint, error)
}

type ComplaintServiceImpl struct {
	ComplaintRepository repository.ComplaintRepository
	Validate            *validator.Validate
}

func NewComplaintService(ComplaintRepository repository.ComplaintRepository, Validate *validator.Validate) *ComplaintServiceImpl {
	return &ComplaintServiceImpl{
		ComplaintRepository: ComplaintRepository,
		Validate:            Validate,
	}
}

func (s *ComplaintServiceImpl) CreateComplaint(ctx echo.Context, request web.ComplaintCreateRequest) (*domain.Complaint, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}

	complaint := req.ComplaintCreateRequestToComplaintDomain(request)

	result, err := s.ComplaintRepository.Create(complaint)
	if err != nil {
		return nil, fmt.Errorf("error when creating complaint: %s", err.Error())
	}

	return result, nil
}

func (s ComplaintServiceImpl) FindById(ctx echo.Context, id string) (*domain.Complaint, error) {
	complaint, err := s.ComplaintRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if complaint == nil {
		return nil, fmt.Errorf("complaint not found")
	}

	return complaint, nil
}

func (s *ComplaintServiceImpl) FindAll(ctx echo.Context) ([]domain.Complaint, error) {
	complaints, err := s.ComplaintRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("complaints not found")
	}

	return complaints, nil
}
