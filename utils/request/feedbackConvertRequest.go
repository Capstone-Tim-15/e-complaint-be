package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FeedbackCreateRequestToFeedbackDomain(request web.FeedbackCreateRequest) *domain.Feedback {
	return &domain.Feedback{
		Role:       request.Role,
		Fullname:   request.Fullname,
		PhotoImage: request.PhotoImage,
		News_ID:    request.News_ID,
		Content:    request.Content,
	}
}

func FeedbackUpdateRequestToFeedbackDomain(request web.FeedbackUpdateRequest) *domain.Feedback {
	return &domain.Feedback{
		News_ID: request.News_ID,
		Content: request.Content,
	}
}

func FeedbackDomaintoFeedbackSchema(request domain.Feedback) *schema.Feedback {
	return &schema.Feedback{
		Fullname:   request.Fullname,
		Role:       schema.Role(request.Role),
		PhotoImage: request.PhotoImage,
		News_ID:    request.News_ID,
		Content:    request.Content,
	}
}
