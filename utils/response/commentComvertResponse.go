package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func CommentSchematoCommentDomain(comment *schema.Comment) *domain.Comment {
	return &domain.Comment{
		ID:           comment.ID,
		Complaint_ID: comment.Complaint_ID,
		Role:         string(comment.Role),
		Message:      comment.Message,
	}
}

func CommentDomaintoCommentResponse(comment *domain.Comment) web.CommentResponse {
	return web.CommentResponse{
		ID:           comment.ID,
		Complaint_ID: comment.Complaint_ID,
		Role:         comment.Role,
		Message:      comment.Message,
	}
}
