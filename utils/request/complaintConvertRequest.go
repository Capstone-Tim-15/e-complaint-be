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
		Title:       request.Title,
		Content:     request.Content,
		Address:     request.Address,
		Status:      schema.Status(request.Status),
		ImageUrl:    request.ImageUrl,
	}
}

func ComplaintCreateRequestToComplaintDomain(request web.ComplaintCreateRequest) *domain.Complaint {
	return &domain.Complaint{
		User_ID:     request.User_ID,
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Content:     request.Content,
		Address:     request.Address,
		Status:      request.Status,
		ImageUrl:    request.ImageUrl,
	}
}

func ComplaintUpdateRequestToComplaintDomain(request web.ComplaintUpdateRequest) *domain.Complaint {
	return &domain.Complaint{
		Category_ID: request.Category_ID,
		Title:       request.Title,
		Status:      request.Status,
		Content:     request.Content,
		Address:     request.Address,
	}
}
