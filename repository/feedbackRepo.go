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
	FindByAll(page, pageSize int) ([]domain.Feedback, int64, error)
	FindByNewsId(newsID string, page, pageSize int) ([]domain.Feedback, int64, error)
	CheckUser(senderId string) (*domain.User, error)
	CheckAdmin(senderId string) (*domain.Admin, error)
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
	result := repository.DB.Table("feedbacks").Where("id = ?", id).Updates(feedbackDb)
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

func (repository *FeedbackRepositoryImpl) FindByAll(page, pageSize int) ([]domain.Feedback, int64, error) {
	offset := (page - 1) * pageSize
	feedback := []domain.Feedback{}
	var totalCount int64
	resultCount := repository.DB.Model(&domain.Feedback{}).Where("deleted_at is NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}
	result := repository.DB.Where("deleted_at is NULL").Offset(offset).Limit(pageSize).Order("created_at ASC").Find(&feedback)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return feedback, totalCount, nil
}

func (repository *FeedbackRepositoryImpl) FindByNewsId(newsID string, page, pageSize int) ([]domain.Feedback, int64, error) {
	offset := (page - 1) * pageSize
	feedback := []domain.Feedback{}
	var totalCount int64
	resultCount := repository.DB.Model(&domain.Feedback{}).Where("deleted_at is NULL AND news_id = ?", newsID).Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}
	result := repository.DB.Where("deleted_at is NULL").Where("news_id = ?", newsID).Offset(offset).Limit(pageSize).Order("created_at ASC").Find(&feedback)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return feedback, totalCount, nil
}

func (repository *FeedbackRepositoryImpl) CheckUser(senderId string) (*domain.User, error) {
	user := domain.User{}
	result := repository.DB.Where("id = ?", senderId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *FeedbackRepositoryImpl) CheckAdmin(senderId string) (*domain.Admin, error) {
	admin := domain.Admin{}
	result := repository.DB.Where("id = ?", senderId).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
