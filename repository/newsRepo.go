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
	FindByAll(page, pageSize int) ([]domain.News, int64, error)
	FindByTitle(title string, page, pageSize int) ([]domain.News, int64, error)
	FindByCategory(category string, limit int64) ([]domain.News, int64, error)
}

type NewsRepositoryImpl struct {
	DB *gorm.DB
}

func NewNewsRepository(DB *gorm.DB) NewsRepository {
	return &NewsRepositoryImpl{DB: DB}
}

func (repository *NewsRepositoryImpl) Create(news *domain.News) (*domain.News, error) {
	var newsDb *schema.News
	newsDb = req.NewsDomaintoNewsSchema(*news)
	for {
		newsDb = req.NewsDomaintoNewsSchema(*news)
		newsDb.ID = helper.GenerateRandomString()
		result := repository.DB.First(&newsDb, newsDb.ID)
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
	result := repository.DB.Where("deleted_at IS NULL").Preload("Admin").Preload("Feedback.User").Preload("Likes").Preload("Category").First(&news, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &news, nil
}

func (repository *NewsRepositoryImpl) FindByAll(page, pageSize int) ([]domain.News, int64, error) {
	offset := (page - 1) * pageSize
	news := []domain.News{}
	var totalCount int64
	resultCount := repository.DB.Model(&domain.News{}).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}

	result := repository.DB.Where("deleted_at IS NULL").Preload("Admin").Preload("Feedback.User").Preload("Likes").Preload("Category").Offset(offset).Limit(pageSize).Order("created_at ASC").Find(&news)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return news, totalCount, nil
}

func (repository *NewsRepositoryImpl) FindByTitle(title string, page, pageSize int) ([]domain.News, int64, error) {
	offset := (page - 1) * pageSize
	news := []domain.News{}
	var totalCount int64
	resultCount := repository.DB.Model(&domain.News{}).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}
	result := repository.DB.Where("deleted_at IS NULL").Preload("Admin").Preload("Feedback.User").Preload("Likes").Preload("Category").Offset(offset).Limit(pageSize).Order("created_at ASC").Find(&news, "title LIKE  ?", title+"%")
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return news, totalCount, nil
}

func (repository *NewsRepositoryImpl) FindByCategory(category string, limit int64) ([]domain.News, int64, error) {
	news := []domain.News{}
	var totalCount int64

	resultCount := repository.DB.Model(&domain.News{}).Where("category_id = ?", category).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}
	result := repository.DB.Where("category_id = ?", category).Where("deleted_at IS NULL").Preload("Admin").Preload("Feedback.User").Preload("Likes").Preload("Category").Order("created_at desc").Limit(int(limit)).Find(&news)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return news, totalCount, nil
}
