package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func CategoryRequestToCategoryDomain(request web.CategoryRequest) *domain.Category {
	return &domain.Category{
		CategoryName: request.Name,
	}
}

func CategoryDomaintoCategorySchema(request domain.Category) *schema.Category {
	return &schema.Category{
		ID:           request.ID,
		CategoryName: request.CategoryName,
	}
}
