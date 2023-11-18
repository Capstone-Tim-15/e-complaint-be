package res

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FAQSchemaIntoDomain(schema *schema.FAQ) *domain.FAQ {
	return &domain.FAQ{
		ID:          schema.ID,
		Title:       schema.Title,
		Content:     schema.Content,
		Category_ID: schema.Category_ID,
	}
}

func FAQDomaintoResponse(request *domain.FAQ) *web.FaqResponse {
	return &web.FaqResponse{
		ID:          request.ID,
		Title:       request.Title,
		Content:     request.Content,
		Category_ID: request.Category_ID,
	}
}
func ConvertFAQResponse(request []domain.FAQ) []web.FaqResponse {
	var results []web.FaqResponse
	for _, faq := range request {
		faqResponse := web.FaqResponse{
			ID:          faq.ID,
			Title:       faq.Title,
			Content:     faq.Content,
			Category_ID: faq.Category_ID,
		}
		results = append(results, faqResponse)
	}
	return results
}
