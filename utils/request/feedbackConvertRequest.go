package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FeedbackCreateRequestToFeedbackDomain(request web.FeedbackCreateRequest) *domain.Feedback {
	return &domain.Feedback{
		User_ID: request.User_ID,
		News_ID: request.News_ID,
		Content: request.Content,
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
		User_ID: request.User_ID,
		News_ID: request.News_ID,
		Content: request.Content,
	}
}
