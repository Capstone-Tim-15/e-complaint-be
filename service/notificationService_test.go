package service

import (
	"ecomplaint/model/domain"
	"ecomplaint/test/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFindAllNotification(t *testing.T) {
	mockNotificationRepository := new(mocks.NotificationRepository)
	validate := validator.New()
	e := echo.New()
	ctx := e.AcquireContext()

	NotificationService := &NotificationServiceImpl{
		NotificationRepository: mockNotificationRepository,
		Validate:               validate,
	}

	mockNotificationRepository.On("FindAll").Return([]domain.Notification{}, nil)

	_, err := NotificationService.FindAllNotification(ctx)

	assert.NoError(t, err)

	mockNotificationRepository.AssertExpectations(t)

}