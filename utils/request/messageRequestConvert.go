package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func MessDomaintoMessSchema(request *domain.Comment) *schema.Comment {
	return &schema.Comment{
		Complaint_ID: request.Complaint_ID,
		Role:         schema.Role(request.Role),
		Message:      request.Message,
	}
}

func MessCreateRequesttoMessDomain(request web.CommentCreateRequest) *domain.Comment {
	return &domain.Comment{
		Complaint_ID: request.Complaint_ID,
		Role:         request.Role,
		Message:      request.Message,
	}
}
