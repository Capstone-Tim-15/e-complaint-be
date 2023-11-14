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

type NewsService interface {
	CreateNews(ctx echo.Context, request web.NewsCreateRequest) (*domain.News, error)
	UpdateNews(ctx echo.Context, request web.NewsUpdateRequest, id string) (*domain.News, error)
	DeleteNews(ctx echo.Context, id string) error
	FindById(ctx echo.Context, id string) (*domain.News, error)
	FindByAll(ctx echo.Context) ([]domain.News, error)
}

type NewsServiceImp struct {
	NewsRepository repository.NewsRepository
	validate       *validator.Validate
}

func NewNewsService(NewsRepository repository.NewsRepository, validate *validator.Validate) *NewsServiceImp {
	return &NewsServiceImp{
		NewsRepository: NewsRepository,
		validate:       validate,
	}
}

func (context *NewsServiceImp) CreateNews(ctx echo.Context, request web.NewsCreateRequest) (*domain.News, error) {
	err := context.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	news := req.NewsCreateRequestToNewsDomain(request)

	result, err := context.NewsRepository.Create(news)
	if err != nil {
		return nil, fmt.Errorf("error when creating News: %s", err.Error())
	}
	return result, nil
}

func (context *NewsServiceImp) UpdateNews(ctx echo.Context, request web.NewsUpdateRequest, id string) (*domain.News, error) {
	err := context.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingNews, _ := context.NewsRepository.FindById(id)
	if existingNews == nil {
		return nil, fmt.Errorf("news not found")
	}

	news := req.NewsUpdateRequestToNewsDomain(request)
	_, err = context.NewsRepository.Update(news, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating news: %s", err.Error())
	}
	result, err := context.NewsRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating news: %s", err.Error())
	}
	return result, nil

}

func (context *NewsServiceImp) DeleteNews(ctx echo.Context, id string) error {
	existingNews, _ := context.NewsRepository.FindById(id)
	if existingNews == nil {
		return fmt.Errorf("News not Found")
	}
	err := context.NewsRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting news: %s", err)
	}
	return nil
}

func (context *NewsServiceImp) FindById(ctx echo.Context, id string) (*domain.News, error) {
	existingNews, _ := context.NewsRepository.FindById(id)
	fmt.Println(existingNews)
	if existingNews == nil {
		return nil, fmt.Errorf("news not found")
	}
	return existingNews, nil
}

func (context *NewsServiceImp) FindByAll(ctx echo.Context) ([]domain.News, error) {
	news, err := context.NewsRepository.FindByAll()
	if err != nil {
		return nil, fmt.Errorf("news not found")
	}
	return news, nil
}
