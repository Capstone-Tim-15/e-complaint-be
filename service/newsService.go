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

type NewsService interface {
	CreateNews(ctx echo.Context, request web.NewsCreateRequest) (*domain.News, error)
	UpdateNews(ctx echo.Context, request web.NewsUpdateRequest, id string) (*domain.News, error)
	DeleteNews(ctx echo.Context, id string) error
	FindById(ctx echo.Context, id string) (*domain.News, error)
	FindByAll(ctx echo.Context, page, pageSize int) ([]domain.News, int64, error)
	FindByTitle(ctx echo.Context, title string, page, pageSize int) ([]domain.News, int64, error)
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

func (ns *NewsServiceImp) CreateNews(ctx echo.Context, request web.NewsCreateRequest) (*domain.News, error) {
	err := ns.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	news := req.NewsCreateRequestToNewsDomain(request)

	result, err := ns.NewsRepository.Create(news)
	if err != nil {
		return nil, fmt.Errorf("error when creating News: %s", err.Error())
	}
	return result, nil
}

func (ns *NewsServiceImp) UpdateNews(ctx echo.Context, request web.NewsUpdateRequest, id string) (*domain.News, error) {
	err := ns.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingNews, _ := ns.NewsRepository.FindById(id)
	if existingNews == nil {
		return nil, fmt.Errorf("news not found")
	}

	news := req.NewsUpdateRequestToNewsDomain(request)
	_, err = ns.NewsRepository.Update(news, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating news: %s", err.Error())
	}
	result, err := ns.NewsRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating news: %s", err.Error())
	}
	return result, nil

}

func (ns *NewsServiceImp) DeleteNews(ctx echo.Context, id string) error {
	existingNews, _ := ns.NewsRepository.FindById(id)
	if existingNews == nil {
		return fmt.Errorf("news not found")
	}
	err := ns.NewsRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting news: %s", err)
	}
	return nil
}

func (ns *NewsServiceImp) FindById(ctx echo.Context, id string) (*domain.News, error) {
	existingNews, _ := ns.NewsRepository.FindById(id)
	if existingNews == nil {
		return nil, fmt.Errorf("news not found")
	}
	return existingNews, nil
}

func (ns *NewsServiceImp) FindByAll(ctx echo.Context, page, pageSize int) ([]domain.News, int64, error) {
	news, totalCount, err := ns.NewsRepository.FindByAll(page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("news not found")
	}
	return news, totalCount, nil
}

func (ns *NewsServiceImp) FindByTitle(ctx echo.Context, title string, page, pageSize int) ([]domain.News, int64, error) {
	news, totalCount, err := ns.NewsRepository.FindByTitle(title, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("title not found")
	}
	return news, totalCount, nil
}
