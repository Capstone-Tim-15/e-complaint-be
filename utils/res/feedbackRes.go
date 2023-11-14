package res

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FeedbackSchemaToFeedbackDomain(feedback *schema.Feedback) *domain.Feedback {
	return &domain.Feedback{
		ID:      feedback.ID,
		User_ID: feedback.User_ID,
		News_ID: feedback.News_ID,
		Content: feedback.Content,
	}
}

func FeedbackDomainToFeedbackResponse(feedback *domain.Feedback) web.FeedbackResponse {
	return web.FeedbackResponse{
		ID:      feedback.ID,
		User_ID: feedback.User_ID,
		News_ID: feedback.News_ID,
		Content: feedback.Content,
	}
}

func ConvertFeedbackResponse(feedback []domain.Feedback) []web.FeedbackResponse {
	var results []web.FeedbackResponse
	for _, f := range feedback {
		feedbackResponse := web.FeedbackResponse{
			ID:      f.ID,
			User_ID: f.User_ID,
			News_ID: f.News_ID,
			Content: f.Content,
		}
		results = append(results, feedbackResponse)
	}
	return results
}
