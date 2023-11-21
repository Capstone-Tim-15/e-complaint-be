package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func MessSchematoMessDomain(mess *schema.Comment) *domain.Comment {
	return &domain.Comment{
		ID:           mess.ID,
		Complaint_ID: mess.Complaint_ID,
		Role:         string(mess.Role),
		Message:      mess.Message,
	}
}

func MessDomaintoMessResponse(mess *domain.Comment) web.CommentResponse {
	return web.CommentResponse{
		ID:           mess.ID,
		Complaint_ID: mess.Complaint_ID,
		Role:         mess.Role,
		Message:      mess.Message,
	}
}
