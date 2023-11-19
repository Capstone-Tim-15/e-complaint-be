package request

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func ComplaintDomaintoComplaintSchema(request domain.Complaint) *schema.Complaint {
	return &schema.Complaint{
		ID:          request.ID,
		User_ID:     request.User_ID,
		Category_ID: request.Category_ID,
		Admin_ID:    request.Admin_ID,
		Title:       request.Title,
		Status:      schema.Status(request.Status),
		Attachment:  request.Attachment,
		URL:         request.URL,
	}
}

func ComplaintCreateRequestToComplaintDomain(request web.ComplaintCreateRequest) *domain.Complaint {
	return &domain.Complaint{
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Content:     request.Content,
		Status:      request.Status,
	}
}
