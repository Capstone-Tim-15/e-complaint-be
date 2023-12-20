package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
)

func NotificationDomainToConvertResponse(notification []domain.Notification) []web.NotificationResponse {
	var notificationResponse []web.NotificationResponse
	for _, value := range notification {
		notificationResponse = append(notificationResponse, web.NotificationResponse{
			ID:          value.ID,
			Complain_ID: value.Complaint_ID,
			Message:     value.Message,
			Status:      value.Status,
		})
	}

	return notificationResponse
}
