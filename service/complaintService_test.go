package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"fmt"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateComplaintSuccess(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}

	request	:= web.ComplaintCreateRequest{
		User_ID: "123456",
		Category_ID: "123456",
		Title: "test",
		Content: "test",
		Address: "test",
		Status: "test",
		ImageUrl: "test.png",
	}

	mockComplaintRepository.On("Create", mock.AnythingOfType("*domain.Complaint")).Return(nil, nil)

	_, err := ComplaintService.CreateComplaint(ctx, request)

	assert.NoError(t, err)

	mockComplaintRepository.AssertExpectations(t)
}

func TestCreateComplaintFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}

	request	:= web.ComplaintCreateRequest{
		User_ID: "123456",
		Title: "test",
		Content: "test",
		Address: "test",
		Status: "test",
	}

	_, err := ComplaintService.CreateComplaint(ctx, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed: validation error on field Category_ID, tag min")

	mockComplaintRepository.AssertExpectations(t)
}

func TestCreateComplaintFailCreating(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}

	request	:= web.ComplaintCreateRequest{
		User_ID: "123456",
		Category_ID: "123456",
		Title: "test",
		Content: "test",
		Address: "test",
		Status: "test",
		ImageUrl: "test.png",
	}

	mockComplaintRepository.On("Create", mock.AnythingOfType("*domain.Complaint")).Return(nil, fmt.Errorf("error creating complaint"))

	_, err := ComplaintService.CreateComplaint(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when creating complaint: error creating complaint"))

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindById(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	role := "user"

	mockComplaintRepository.On("FindById", complaintId, role).Return(&domain.Complaint{}, nil)

	result, err := ComplaintService.FindById(ctx, complaintId, role)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByIdFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()	
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	role := "user"

	mockComplaintRepository.On("FindById", complaintId, role).Return(nil, nil)

	result, err := ComplaintService.FindById(ctx, complaintId, role)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaint not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByStatusUser(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	status := "test"
	userId := "123456"

	mockComplaintRepository.On("FindByStatusUser", status, userId).Return([]domain.Complaint{}, nil)

	result, err := ComplaintService.FindByStatusUser(ctx, status, userId)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByStatusUserFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	status := "test"
	userId := "123456"

	mockComplaintRepository.On("FindByStatusUser", status, userId).Return(nil, nil)

	result, err := ComplaintService.FindByStatusUser(ctx, status, userId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaints not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByCategory(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	category := "test"
	limit := int64(1)

	mockComplaintRepository.On("FindByCategory", category, limit).Return([]domain.Complaint{}, int64(1), nil)

	result, _, err := ComplaintService.FindByCategory(ctx, category, limit)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByCategoryFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	category := "test"
	limit := int64(1)

	mockComplaintRepository.On("FindByCategory", category, limit).Return(nil, int64(0), nil)

	result, _, err := ComplaintService.FindByCategory(ctx, category, limit)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaints not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByStatus(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	status := "test"
	page := 1
	pageSize := 1

	mockComplaintRepository.On("FindByStatus", status, page, pageSize).Return([]domain.Complaint{}, int64(1), nil)

	result, _, err := ComplaintService.FindByStatus(ctx, status, page, pageSize)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindByStatusFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	status := "test"
	page := 1
	pageSize := 1

	mockComplaintRepository.On("FindByStatus", status, page, pageSize).Return(nil, int64(0), nil)

	result, _, err := ComplaintService.FindByStatus(ctx, status, page, pageSize)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaints not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindAllUser(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	userId := "123456"

	mockComplaintRepository.On("FindAllUser", userId).Return([]domain.Complaint{}, nil)

	result, err := ComplaintService.FindAllUser(ctx, userId)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindAllUserFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	userId := "123456"

	mockComplaintRepository.On("FindAllUser", userId).Return(nil, nil)

	result, err := ComplaintService.FindAllUser(ctx, userId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaints not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindAll(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	page := 1
	pageSize := 1

	mockComplaintRepository.On("FindAll", page, pageSize).Return([]domain.Complaint{}, int64(1), nil)

	result, _, err := ComplaintService.FindAll(ctx, page, pageSize)

	assert.NoError(t, err)
	assert.NotNil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestFindAllFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	page := 1
	pageSize := 1

	mockComplaintRepository.On("FindAll", page, pageSize).Return(nil, int64(0), fmt.Errorf("complaints not found"))

	result, _, err := ComplaintService.FindAll(ctx, page, pageSize)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaints not found"))
	assert.Nil(t,result)

	mockComplaintRepository.AssertExpectations(t)
}

func TestUpdateComplaint(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	request	:= web.ComplaintUpdateRequest{
		Category_ID: "123456",
		Title: "test",
		Content: "test",
		Address: "test",
		Status: "test",
		ImageUrl: "test.png",
	}

	mockComplaintRepository.On("Update", mock.AnythingOfType("*domain.Complaint"), complaintId).Return(&domain.Complaint{}, nil)

	_, err := ComplaintService.UpdateComplaint(ctx, complaintId, request)

	assert.NoError(t, err)

	mockComplaintRepository.AssertExpectations(t)
}

func TestUpdateComplaintFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,	
		Validate: validate,
	}
	complaintId := "123456"
	request	:= web.ComplaintUpdateRequest{
		Title: "test",
		Content: "test",
		Address: "test",
		Status: "test",
	}

	mockComplaintRepository.On("Update", mock.AnythingOfType("*domain.Complaint"), complaintId).Return(nil, fmt.Errorf("complaint not found"))

	_, err := ComplaintService.UpdateComplaint(ctx, complaintId, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error when updating complaint: complaint not found")

	mockComplaintRepository.AssertExpectations(t)
}

func TestDeleteComplaint(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	role := ""

	mockComplaintRepository.On("FindById", complaintId, role).Return(&domain.Complaint{}, nil)
	mockComplaintRepository.On("Delete", mock.AnythingOfType("*domain.Complaint"), complaintId).Return(nil)

	err := ComplaintService.DeleteComplaint(ctx, complaintId)

	assert.NoError(t, err)

	mockComplaintRepository.AssertExpectations(t)
}

func TestDeleteComplaintFailed(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	role := ""

	mockComplaintRepository.On("FindById", complaintId, role).Return(nil, nil)

	err := ComplaintService.DeleteComplaint(ctx, complaintId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("complaint not found"))

	mockComplaintRepository.AssertExpectations(t)
}

func TestDeleteComplaintFailDeleting(t *testing.T){
	mockComplaintRepository := new(mocks.ComplaintRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	ComplaintService := &ComplaintServiceImpl{
		ComplaintRepository: mockComplaintRepository,
		Validate: validate,
	}
	complaintId := "123456"
	role := ""

	mockComplaintRepository.On("FindById", complaintId, role).Return(&domain.Complaint{}, nil)
	mockComplaintRepository.On("Delete", mock.AnythingOfType("*domain.Complaint"), complaintId).Return(fmt.Errorf("error deleting complaint"))

	err := ComplaintService.DeleteComplaint(ctx, complaintId)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when deleting complaint: error deleting complaint"))

	mockComplaintRepository.AssertExpectations(t)
}