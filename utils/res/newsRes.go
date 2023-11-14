package res

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
	return web.NewsResponse{
		ID:       news.ID,
		Admin_ID: news.Admin_ID,
		Title:    news.Title,
		Content:  news.Content,
		Date:     news.Date,
	}
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
		results = append(results, newsResponse)
	}
	return results
}
