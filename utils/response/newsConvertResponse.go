package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func NewsSchemaToNewsDomain(news *schema.News) *domain.News {
	return &domain.News{
		ID:       news.ID,
		Admin_ID: news.Admin_ID,
		Title:    news.Title,
		Content:  news.Content,
		Date:     news.Date,
	}
}

func NewsDomainToNewsResponse(news *domain.News) web.NewsResponse {
	newsResponse := web.NewsResponse{
		ID:       news.ID,
		Admin_ID: news.Admin_ID,
		Title:    news.Title,
		Content:  news.Content,
		Date:     news.Date,
	}
	for _, f := range news.Feedback {
		feedbackResponse := domain.Feedback{
			ID:      f.ID,
			User_ID: f.User_ID,
			News_ID: f.News_ID,
			Content: f.Content,
		}
		newsResponse.Feedback = append(newsResponse.Feedback, feedbackResponse)
	}
	for _, l := range news.Likes {
		likesResponse := domain.Likes{
			ID:      l.ID,
			User_ID: l.User_ID,
			News_ID: l.News_ID,
			Status:  l.Status,
		}
		newsResponse.Likes = append(newsResponse.Likes, likesResponse)
	}

	return newsResponse

}

func ConvertNewsResponse(news []domain.News) []web.NewsResponse {
	var results []web.NewsResponse
	for _, n := range news {
		newsResponse := web.NewsResponse{
			ID:       n.ID,
			Admin_ID: n.Admin_ID,
			Title:    n.Title,
			Content:  n.Content,
			Date:     n.Date,
		}
		for _, f := range n.Feedback {
			feedbackResponse := domain.Feedback{
				ID:      f.ID,
				User_ID: f.User_ID,
				News_ID: f.News_ID,
				Content: f.Content,
			}
			newsResponse.Feedback = append(newsResponse.Feedback, feedbackResponse)
		}
		for _, l := range n.Likes {
			likesResponse := domain.Likes{
				ID:      l.ID,
				User_ID: l.User_ID,
				News_ID: l.News_ID,
				Status:  l.Status,
			}
			newsResponse.Likes = append(newsResponse.Likes, likesResponse)
		}

		results = append(results, newsResponse)
	}
	return results
}
