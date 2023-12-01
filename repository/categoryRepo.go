package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(cat *domain.Category) (*domain.Category, error)
	FindById(id string) (*domain.Category, error)
	FindAll() ([]domain.Category, error)
	Update(cat *domain.Category, id string) (*domain.Category, error)
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: DB}
}

func (repository *CategoryRepositoryImpl) FindAll() ([]domain.Category, error) {
	var categories []domain.Category

	result := repository.DB.Preload("FAQ").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (repository *CategoryRepositoryImpl) FindById(id string) (*domain.Category, error) {
	faq := domain.Category{}

	result := repository.DB.Where("id = ?", id).First(&faq)
	if result.Error != nil {
		return nil, result.Error
	}

	return &faq, nil
}

func (repository *CategoryRepositoryImpl) Create(cat *domain.Category) (*domain.Category, error) {
	var catDB *schema.Category
	for {
		catDB = req.CategoryDomaintoCategorySchema(*cat)
		catDB.ID = helper.GenerateRandomString()

		result := repository.DB.First(&cat, catDB.ID)
		if result.Error != nil {
			break
		}
	}

	result := repository.DB.Create(catDB)
	if result.Error != nil {
		return nil, result.Error
	}

	return cat, nil
}

func (repository *CategoryRepositoryImpl) Update(cat *domain.Category, id string) (*domain.Category, error) {
	result := repository.DB.Where("id = ?", id).Updates(&cat)
	if result.Error != nil {
		return nil, result.Error
	}

	return cat, nil
}
