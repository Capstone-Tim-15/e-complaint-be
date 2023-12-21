package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
	"ecomplaint/test/mocks"
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsCreateRequest{
		Title:       "Test",
		Admin_ID:    "1",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	mockNewsRepository.On("Create", mock.AnythingOfType("*domain.News")).Return(nil, nil)

	_, err := newsService.CreateNews(ctx, request)

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestCreateNewsFailValidate(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsCreateRequest{
		Admin_ID:    "1",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	_, err := newsService.CreateNews(ctx, request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed: validation error on field Title, tag required")

	mockNewsRepository.AssertExpectations(t)
}

func TestCreateNewsFailCreating(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsCreateRequest{
		Title:       "Test",
		Admin_ID:    "1",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	mockNewsRepository.On("Create", mock.AnythingOfType("*domain.News")).Return(nil, fmt.Errorf("error when creating news"))

	_, err := newsService.CreateNews(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when creating News: error when creating news"))

	mockNewsRepository.AssertExpectations(t)
}

func TestUpdateNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsUpdateRequest{
		Title:       "Test",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	mockNewsRepository.On("FindById", "123456").Return(&domain.News{}, nil)
	mockNewsRepository.On("Update", mock.AnythingOfType("*domain.News"), "123456").Return(nil, nil)

	_, err := newsService.UpdateNews(ctx, request, "123456")

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestUpdateNewsFailValidate(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsUpdateRequest{
		Title:       "Test",
		Content:     "Test",
		Category_ID: "1234",
		Date:        time,
		ImageUrl:   "Test",
	}

	_, err := newsService.UpdateNews(ctx, request, "123456")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed: validation error on field Category_ID, tag min")

	mockNewsRepository.AssertExpectations(t)
}

func TestUpdateNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsUpdateRequest{
		Title:       "Test",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	mockNewsRepository.On("FindById", "123456").Return(nil, nil)

	_, err := newsService.UpdateNews(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestUpdateNewsFailUpdating(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}
	time, _ := time.Parse("2006-01-02", "2021-01-01")
	request := web.NewsUpdateRequest{
		Title:       "Test",
		Content:     "Test",
		Category_ID: "123456",
		Date:        time,
		ImageUrl:   "Test",
	}

	mockNewsRepository.On("FindById", "123456").Return(&domain.News{}, nil)
	mockNewsRepository.On("Update", mock.AnythingOfType("*domain.News"), "123456").Return(nil, fmt.Errorf("error when updating news"))

	_, err := newsService.UpdateNews(ctx, request, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when updating news: error when updating news"))

	mockNewsRepository.AssertExpectations(t)
}

func TestDeleteNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindById", "123456").Return(&domain.News{}, nil)
	mockNewsRepository.On("Delete", "123456").Return(nil)

	err := newsService.DeleteNews(ctx, "123456")

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestDeleteNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindById", "123456").Return(nil, nil)

	err := newsService.DeleteNews(ctx, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestDeleteNewsFailDeleting(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindById", "123456").Return(&domain.News{}, nil)
	mockNewsRepository.On("Delete", "123456").Return(fmt.Errorf("error when deleting news"))

	err := newsService.DeleteNews(ctx, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("error when deleting news: error when deleting news"))

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByIdNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindById", "123456").Return(&domain.News{}, nil)

	_, err := newsService.FindById(ctx, "123456")

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByIdNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindById", "123456").Return(nil, nil)

	_, err := newsService.FindById(ctx, "123456")

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestFindAllNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByAll",1,10).Return([]domain.News{}, int64(0), nil)

	_, _, err := newsService.FindByAll(ctx, 1, 10)

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestFindAllNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByAll",1,10).Return([]domain.News{}, int64(0), fmt.Errorf("news not found"))

	_, _, err := newsService.FindByAll(ctx, 1, 10)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByTitleNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByTitle","Test",1,10).Return([]domain.News{}, int64(0), nil)

	_, _, err := newsService.FindByTitle(ctx, "Test", 1, 10)

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByTitleNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByTitle","Test",1,10).Return([]domain.News{}, int64(0), fmt.Errorf("news not found"))

	_, _, err := newsService.FindByTitle(ctx, "Test", 1, 10)

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByCategoryNews(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByCategory","Test", int64(1)).Return([]domain.News{}, int64(0), nil)

	_, _, err := newsService.FindByCategory(ctx, "Test", int64(1))

	assert.NoError(t, err)

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByCategoryNewsFail(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByCategory","Test", int64(1)).Return([]domain.News{}, int64(0), fmt.Errorf("news not found"))

	_, _, err := newsService.FindByCategory(ctx, "Test", int64(1))

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}

func TestFindByCategoryNewsFailGet(t *testing.T){
	mockNewsRepository := new(mocks.NewsRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	newsService := &NewsServiceImp{
		NewsRepository: mockNewsRepository,
		validate: validate,
	}

	mockNewsRepository.On("FindByCategory","Test", int64(1)).Return(nil, int64(0), nil)

	_, _, err := newsService.FindByCategory(ctx, "Test", int64(1))

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("news not found"))

	mockNewsRepository.AssertExpectations(t)
}