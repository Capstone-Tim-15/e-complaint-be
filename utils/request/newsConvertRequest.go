package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
	"time"
)

func NewsCreateRequestToNewsDomain(request web.NewsCreateRequest) *domain.News {
	dateValue := request.Date
	if dateValue == (time.Time{}) {
		now := time.Now()
		dateValue = now
	}

	return &domain.News{
		Admin_ID:    request.Admin_ID,
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Content:     request.Content,
		Date:        dateValue,
		ImageUrl:    request.ImageUrl,
	}
}

func NewsUpdateRequestToNewsDomain(request web.NewsUpdateRequest) *domain.News {
	return &domain.News{
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Content:     request.Content,
		Date:        request.Date,
		ImageUrl:    request.ImageUrl,
	}
}

func NewsDomaintoNewsSchema(request domain.News) *schema.News {
	return &schema.News{
		Admin_ID:    request.Admin_ID,
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Content:     request.Content,
		Date:        request.Date,
		ImageUrl:    request.ImageUrl,
	}
}
