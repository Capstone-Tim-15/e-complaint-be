package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func CommentDomaintoCommentSchema(request *domain.Comment) *schema.Comment {
	return &schema.Comment{
		Complaint_ID: request.Complaint_ID,
		Fullname:     request.Fullname,
		Role:         schema.Role(request.Role),
		Message:      request.Message,
	}
}

func CommentCreateRequesttoCommentDomain(request web.CommentCreateRequest) *domain.Comment {
	return &domain.Comment{
		Complaint_ID: request.Complaint_ID,
		Fullname:     request.Fullname,
		Role:         request.Role,
		Message:      request.Message,
	}
}
