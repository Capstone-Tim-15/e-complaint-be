package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"
	"fmt"
	"gorm.io/gorm"
)

type FaqRepository interface {
	Create(faq *domain.FAQ) (*domain.FAQ, error)
	FindById(id string) (*domain.FAQ, error)
	FindByCategory(category string) (*domain.FAQ, error)
	FindAll() ([]domain.FAQ, error)
	Update(FAQ *domain.FAQ, id string) (*domain.FAQ, error)
}

type FAQRepositoryImpl struct {
	DB *gorm.DB
}

func NewFAQRepository(DB *gorm.DB) FaqRepository {
	return &FAQRepositoryImpl{DB: DB}
}

func (repository *FAQRepositoryImpl) Create(faq *domain.FAQ) (*domain.FAQ, error) {
	var faqDb *schema.FAQ

	for {
		faqDb = req.FAQDomaintoAdminSchema(*faq)
		faqDb.ID = helper.GenerateRandomString()

		result := repository.DB.First(&faq, faqDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := repository.DB.Create(&faqDb)
	if result.Error != nil {
		return nil, result.Error
	}
	faq = res.FAQSchemaIntoDomain(faqDb)

	return faq, nil
}

func (repository *FAQRepositoryImpl) FindById(id string) (*domain.FAQ, error) {
	faq := domain.FAQ{}

	result := repository.DB.Where("id = ?", id).First(&faq)
	if result.Error != nil {
		return nil, result.Error
	}

	return &faq, nil
}

func (repository *FAQRepositoryImpl) FindByCategory(category string) (*domain.FAQ, error) {
	faq := domain.FAQ{}

	result := repository.DB.Where("category_id = ?", category).First(&faq)
	if result.Error != nil {
		return nil, result.Error
	}

	return &faq, nil
}

func (repository *FAQRepositoryImpl) FindAll() ([]domain.FAQ, error) {
	var faqs []domain.FAQ

	result := repository.DB.Find(&faqs)
	if result.Error != nil {
		return nil, result.Error
	}

	return faqs, nil
}

func (repository *FAQRepositoryImpl) Update(faq *domain.FAQ, id string) (*domain.FAQ, error) {
	faqDb := req.FAQDomaintoAdminSchema(*faq)
	result := repository.DB.Table("faqs").Where("id = ?", id).Updates(faqDb)
	if result.Error != nil {
		return nil, result.Error
	}

	faq = res.FAQSchemaIntoDomain(faqDb)
	fmt.Println(id)

	fmt.Println(faq)

	return faq, nil
}
