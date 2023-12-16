package repository

import (
	"ecomplaint/model/domain"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	FindAll() ([]domain.Notification, error)
}

type NotificationRepositoryImpl struct {
	DB *gorm.DB
}

func NewNotificationRepository(DB *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{DB: DB}
}

func (r *NotificationRepositoryImpl) FindAll() ([]domain.Notification, error) {
	var notifications []domain.Notification

	result := r.DB.Where("deleted_at IS NULL").Preload("Complaint").Order("created_at DESC").Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	return notifications, nil
}
