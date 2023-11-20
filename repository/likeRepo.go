package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/request"
	"ecomplaint/utils/response"
	"gorm.io/gorm"
)

type LikeRepository interface {
	Create(like *domain.Likes) (*domain.Likes, error)
	Update(like *domain.Likes, id string) (*domain.Likes, error)
	Delete(id string) error
	FindById(id string) (*domain.Likes, error)
	FindByAll() ([]domain.Likes, error)
}

type LikeRepositoryImpl struct {
	DB *gorm.DB
}

func NewLikeRepository(DB *gorm.DB) LikeRepository {
	return &LikeRepositoryImpl{DB: DB}
}

func (repository *LikeRepositoryImpl) Create(like *domain.Likes) (*domain.Likes, error) {
	var likeDb *schema.Likes

	for {
		likeDb = request.LikeDomaintoLikeSchema(*like)
		likeDb.ID = helper.GenerateRandomString()
		result := repository.DB.First(&like, likeDb.ID)
		if result.Error != nil {
			break
		}
	}
	result := repository.DB.Create(&likeDb)
	if result.Error != nil {
		return nil, result.Error
	}
	like = response.LikesSchemaToLikesDomain(likeDb)
	return like, nil
}

func (repository *LikeRepositoryImpl) Update(like *domain.Likes, id string) (*domain.Likes, error) {
	likeDb := request.LikeDomaintoLikeSchema(*like)
	result := repository.DB.Table("likes").Where("id = ?", id).Updates(likeDb)
	if result.Error != nil {
		return nil, result.Error
	}
	like = response.LikesSchemaToLikesDomain(likeDb)
	return like, nil
}

func (repository *LikeRepositoryImpl) Delete(id string) error {
	result := repository.DB.Table("likes").Where("id = ?", id).Delete(&schema.Likes{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *LikeRepositoryImpl) FindById(id string) (*domain.Likes, error) {
	like := domain.Likes{}
	result := repository.DB.Where("deleted_at IS NULL").First(&like, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &like, nil
}

func (repository *LikeRepositoryImpl) FindByAll() ([]domain.Likes, error) {
	likes := []domain.Likes{}
	result := repository.DB.Where("deleted_at is NULL").Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}
