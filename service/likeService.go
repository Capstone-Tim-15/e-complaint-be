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

type LikeService interface {
	CreateLike(ctx echo.Context, request web.LikesCreateRequest) (*domain.Likes, error)
	UpdateLike(ctx echo.Context, request web.LikesUpdateRequest, id string) (*domain.Likes, error)
	DeleteLike(ctx echo.Context, id string) error
	FindById(ctx echo.Context, id string) (*domain.Likes, error)
	FindByAll(ctx echo.Context) ([]domain.Likes, error)
}

type LikeServiceImp struct {
	LikesRepository repository.LikeRepository
	validate        *validator.Validate
}

func NewLikeService(LikesRepository repository.LikeRepository, validate *validator.Validate) *LikeServiceImp {
	return &LikeServiceImp{
		LikesRepository: LikesRepository,
		validate:        validate,
	}
}

func (ls *LikeServiceImp) CreateLike(ctx echo.Context, request web.LikesCreateRequest) (*domain.Likes, error) {
	err := ls.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	like := req.LikeCreateRequestToLikeDomain(request)
	result, err := ls.LikesRepository.Create(like)
	if err != nil {
		return nil, fmt.Errorf("error when creating Like: %s ", err.Error())
	}
	return result, nil

}

func (ls *LikeServiceImp) UpdateLike(ctx echo.Context, request web.LikesUpdateRequest, id string) (*domain.Likes, error) {
	err := ls.validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(ctx, err)
	}
	existingLike, _ := ls.LikesRepository.FindById(id)
	if existingLike == nil {
		return nil, fmt.Errorf("like not found")
	}
	like := req.LikeUpdateRequestToLikeDomain(request)
	_, err = ls.LikesRepository.Update(like, id)
	if err != nil {
		return nil, fmt.Errorf("error when updatating like: %s", err.Error())
	}
	result, err := ls.LikesRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("error when updating like: %s", err.Error())
	}
	return result, nil
}

func (ls *LikeServiceImp) DeleteLike(ctx echo.Context, id string) error {
	existingLike, _ := ls.LikesRepository.FindById(id)
	if existingLike == nil {
		return fmt.Errorf("like not found")
	}
	err := ls.LikesRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting like: %s", err)
	}
	return nil
}

func (ls *LikeServiceImp) FindById(ctx echo.Context, id string) (*domain.Likes, error) {
	existingLike, _ := ls.LikesRepository.FindById(id)
	if existingLike == nil {
		return nil, fmt.Errorf("like not found")
	}
	return existingLike, nil
}

func (ls *LikeServiceImp) FindByAll(ctx echo.Context) ([]domain.Likes, error) {
	like, err := ls.LikesRepository.FindByAll()
	if err != nil {
		return nil, fmt.Errorf("like not found")
	}
	return like, nil
}
