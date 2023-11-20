package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	"ecomplaint/utils/request"
	"ecomplaint/utils/response"
	"gorm.io/gorm"
)

type FeedbackRepository interface {
	Create(feedback *domain.Feedback) (*domain.Feedback, error)
	Update(feedback *domain.Feedback, id string) (*domain.Feedback, error)
	Delete(id string) error
	FindById(id string) (*domain.Feedback, error)
	FindByAll() ([]domain.Feedback, error)
}

type FeedbackRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedbackRepository(DB *gorm.DB) FeedbackRepository {
	return &FeedbackRepositoryImpl{DB: DB}
}

func (repository *FeedbackRepositoryImpl) Create(feedback *domain.Feedback) (*domain.Feedback, error) {
	var feedbackDB *schema.Feedback
	if feedbackDB == nil {
		for {
			feedbackDB = request.FeedbackDomaintoFeedbackSchema(*feedback)
			feedbackDB.ID = helper.GenerateRandomString()
			result := repository.DB.First(&feedback, feedbackDB.ID)
			if result.Error != nil {
				break
			}
		}
	}
	result := repository.DB.Create(&feedbackDB)
	if result.Error != nil {
		return nil, result.Error
	}

	feedback = response.FeedbackSchemaToFeedbackDomain(feedbackDB)
	return feedback, nil
}

func (repository *FeedbackRepositoryImpl) Update(feedback *domain.Feedback, id string) (*domain.Feedback, error) {
	feedbackDb := request.FeedbackDomaintoFeedbackSchema(*feedback)
	result := repository.DB.Table("feedback").Where("id = ?", id).Updates(feedbackDb)
	if result.Error != nil {
		return nil, result.Error
	}
	feedback = response.FeedbackSchemaToFeedbackDomain(feedbackDb)

	return feedback, nil
}

func (repository *FeedbackRepositoryImpl) Delete(id string) error {
	result := repository.DB.Table("feedback").Where("id = ?", id).Delete(&schema.Feedback{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *FeedbackRepositoryImpl) FindById(id string) (*domain.Feedback, error) {
	feedback := domain.Feedback{}
	result := repository.DB.Where("deleted_at IS NULL").First(&feedback, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &feedback, nil
}

func (repository *FeedbackRepositoryImpl) FindByAll() ([]domain.Feedback, error) {
	feedback := []domain.Feedback{}
	result := repository.DB.Where("deleted_at is NULL").Find(&feedback)
	if result.Error != nil {
		return nil, result.Error
	}
	return feedback, nil
}
