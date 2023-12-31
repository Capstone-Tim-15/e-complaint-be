package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FeedbackSchemaToFeedbackDomain(feedback *schema.Feedback) *domain.Feedback {
	return &domain.Feedback{
		ID:         feedback.ID,
		Fullname:   feedback.Fullname,
		Role:       string(feedback.Role),
		PhotoImage: feedback.PhotoImage,
		News_ID:    feedback.News_ID,
		Content:    feedback.Content,
	}
}

func FeedbackDomainToFeedbackResponse(feedback *domain.Feedback) web.FeedbackCreateResponse {
	return web.FeedbackCreateResponse{
		ID:         feedback.ID,
		Fullname:   feedback.Fullname,
		Role:       feedback.Role,
		PhotoImage: feedback.PhotoImage,
		News_ID:    feedback.News_ID,
		Content:    feedback.Content,
	}
}

func ConvertFeedbackResponse(feedback []domain.Feedback) []web.FeedbackResponse {
	var results []web.FeedbackResponse
	for _, f := range feedback {
		feedbackResponse := web.FeedbackResponse{
			ID:         f.ID,
			Fullname:   f.Fullname,
			Role:       f.Role,
			PhotoImage: f.PhotoImage,
			News_ID:    f.News_ID,
			Content:    f.Content,
		}
		results = append(results, feedbackResponse)
	}
	return results
}

func FindFeedbackDomainToFeedbackResponse(feedback *domain.Feedback) web.FeedbackResponse {
	return web.FeedbackResponse{
		ID:         feedback.ID,
		Fullname:   feedback.Fullname,
		Role:       feedback.Role,
		PhotoImage: feedback.PhotoImage,
		News_ID:    feedback.News_ID,
		Content:    feedback.Content,
	}
}
