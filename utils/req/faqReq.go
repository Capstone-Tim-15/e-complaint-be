package req

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func FAQRequestToFAQDomain(request web.FaqRequest) *domain.FAQ {
	return &domain.FAQ{
		Title:       request.Title,
		Content:     request.Content,
		Category_ID: request.Category_ID,
	}
}

func FAQDomaintoAdminSchema(request domain.FAQ) *schema.FAQ {
	return &schema.FAQ{
		ID:          request.ID,
		Title:       request.Title,
		Content:     request.Content,
		Category_ID: request.Category_ID,
	}
}
