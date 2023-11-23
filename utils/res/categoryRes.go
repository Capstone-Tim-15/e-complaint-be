package res

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/web"
)

func CategoryCreateResponse(domain *domain.Category) *web.CategoryCreateResponse {
	return &web.CategoryCreateResponse{
		Id:           domain.ID,
		CategoryName: domain.CategoryName,
	}
}

func CategoryResponse(request *domain.Category) web.CategoryResponse {
	catResponse := web.CategoryResponse{
		Id:           request.ID,
		CategoryName: request.CategoryName,
	}
	for _, faq := range request.FAQ {
		faqResponse := web.FaqCategoryResponse{
			ID:      faq.ID,
			Title:   faq.Title,
			Content: faq.Content,
		}
		catResponse.FAQ = append(catResponse.FAQ, faqResponse)
	}
	return catResponse
}

func AllCategoryResponse(request []domain.Category) []web.CategoryResponse {
	var results []web.CategoryResponse
	for _, category := range request {
		catResponse := web.CategoryResponse{
			Id:           category.ID,
			CategoryName: category.CategoryName,
		}
		for _, faq := range category.FAQ {
			faqResponse := web.FaqCategoryResponse{
				ID:      faq.ID,
				Title:   faq.Title,
				Content: faq.Content,
			}
			catResponse.FAQ = append(catResponse.FAQ, faqResponse)
		}
		results = append(results, catResponse)
	}

	return results
}
