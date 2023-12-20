package controller

import (
	"ecomplaint/service"
	res "ecomplaint/utils/response"
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

type NotificationController interface {
	FindAllNotification(ctx echo.Context) error
}

type NotificationControllerImpl struct {
	NotificationService service.NotificationService
}

func NewNotificationController(NotificationService service.NotificationService) NotificationController {
	return &NotificationControllerImpl{
		NotificationService: NotificationService,
	}
}

func (c *NotificationControllerImpl) FindAllNotification(ctx echo.Context) error {
	ctx.Response().Header().Set("Content-Type", "text/event-stream")
	ctx.Response().Header().Set("Cache-Control", "no-cache")
	ctx.Response().Header().Set("Connection", "keep-alive")

	var lastUpdate time.Time

	messageChan := make(chan string)

	for {
		select {
		case <-ctx.Request().Context().Done():
			close(messageChan)
			return nil
		default:
			result, _ := c.NotificationService.FindAllNotification(ctx)
			if len(result) == 0 {
				message := fmt.Sprintf("data: %s\n\n", "null")
				fmt.Fprintf(ctx.Response(), message)
				ctx.Response().Flush()
				lastUpdate = time.Time{}
			}
			if len(result) > 0 && result[0].CreatedAt != lastUpdate {
				results := res.NotificationDomainToConvertResponse(result)
				data, _ := json.Marshal(results)
				message := fmt.Sprintf("data: %s\n\n", data)
				fmt.Fprintf(ctx.Response(), message)
				ctx.Response().Flush()
				lastUpdate = result[0].CreatedAt
			}
		}
		time.Sleep(2 * time.Second)
	}
}
