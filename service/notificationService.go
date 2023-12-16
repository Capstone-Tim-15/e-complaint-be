package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/repository"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type NotificationService interface {
	FindAllNotification(ctx echo.Context) ([]domain.Notification, error)
}

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
	Validate               *validator.Validate
}

func NewNotificationService(NotificationRepository repository.NotificationRepository, Validate *validator.Validate) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		NotificationRepository: NotificationRepository,
		Validate:               Validate,
	}
}

func (s NotificationServiceImpl) FindAllNotification(ctx echo.Context) ([]domain.Notification, error) {
	notifications, err := s.NotificationRepository.FindAll()
	if err != nil {
		return nil, err
	}

	if notifications == nil {
		return nil, fmt.Errorf("notifications not found")
	}

	return notifications, nil
}
