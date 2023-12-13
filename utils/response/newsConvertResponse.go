package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func NewsSchemaToNewsDomain(news *schema.News) *domain.News {
	return &domain.News{
		ID:          news.ID,
		Admin_ID:    news.Admin_ID,
		Category_ID: news.Category_ID,
		Title:       news.Title,
		Content:     news.Content,
		Date:        news.Date,
		ImageUrl:    news.ImageUrl,
	}
}

func NewsDomainToNewsResponse(news *domain.News) web.NewsCreateResponse {
	newsResponse := web.NewsCreateResponse{
		ID:       news.ID,
		Admin_ID: news.Admin_ID,
		Category: news.Category_ID,
		Title:    news.Title,
		Content:  news.Content,
		Date:     news.Date,
		ImageUrl: news.ImageUrl,
	}
	return newsResponse

}

func ConvertNewsResponse(news []domain.News) []web.NewsResponse {
	var results []web.NewsResponse
	for _, n := range news {
		newsResponse := web.NewsResponse{
			ID:         n.ID,
			Admin_ID:   n.Admin_ID,
			Category:   n.Category.CategoryName,
			Name:       n.Admin.Name,
			PhotoImage: n.Admin.ImageUrl,
			Title:      n.Title,
			Content:    n.Content,
			Date:       n.Date,
			ImageUrl:   n.ImageUrl,
			Feedback:   ConvertFeedbackResponse(n.Feedback),
			Likes:      ConvertLikesResponse(n.Likes),
		}
		results = append(results, newsResponse)
	}
	return results
}

func FindNewsDomainToNewsResponse(news *domain.News) web.NewsResponse {
	newsResponse := web.NewsResponse{
		ID:         news.ID,
		Admin_ID:   news.Admin_ID,
		Category:   news.Category.CategoryName,
		Name:       news.Admin.Name,
		PhotoImage: news.Admin.ImageUrl,
		Title:      news.Title,
		Content:    news.Content,
		Date:       news.Date,
		ImageUrl:   news.ImageUrl,
		Feedback:   ConvertFeedbackResponse(news.Feedback),
		Likes:      ConvertLikesResponse(news.Likes),
	}
	return newsResponse
}
