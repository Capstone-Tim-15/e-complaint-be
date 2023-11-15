package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *domain.News) (*domain.News, error)
	Update(news *domain.News, id string) (*domain.News, error)
	Delete(id string) error
	FindById(id string) (*domain.News, error)
	FindByAll() ([]domain.News, error)
}

type NewsRepositoryImpl struct {
	DB *gorm.DB
}

func NewNewsRepository(DB *gorm.DB) NewsRepository {
	return &NewsRepositoryImpl{DB: DB}
}

func (repository *NewsRepositoryImpl) Create(news *domain.News) (*domain.News, error) {
	var newsDb *schema.News
	for {
		newsDb = req.NewsDomaintoNewsSchema(*news)
		newsDb.ID = helper.GenerateRandomString()

		result := repository.DB.First(&news, newsDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := repository.DB.Create(&newsDb)

	if result.Error != nil {
		return nil, result.Error
	}

	news = res.NewsSchemaToNewsDomain(newsDb)
	return news, nil

}

func (repository *NewsRepositoryImpl) Update(news *domain.News, id string) (*domain.News, error) {
	newsDb := req.NewsDomaintoNewsSchema(*news)
	result := repository.DB.Table("news").Where("id = ?", id).Updates(newsDb)
	if result.Error != nil {
		return nil, result.Error
	}
	news = res.NewsSchemaToNewsDomain(newsDb)
	return news, nil
}

func (repository *NewsRepositoryImpl) Delete(id string) error {
	result := repository.DB.Table("news").Where("id = ?", id).Delete(&schema.News{})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (repository *NewsRepositoryImpl) FindById(id string) (*domain.News, error) {
	news := domain.News{}
	result := repository.DB.First(&news, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	println(result)
	return &news, nil
}

func (repository *NewsRepositoryImpl) FindByAll() ([]domain.News, error) {
	news := []domain.News{}
	result := repository.DB.Where("deleted_at IS NULL").Find(&news)
	if result.Error != nil {
		return nil, result.Error
	}
	return news, nil
}
