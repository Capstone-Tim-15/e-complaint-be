package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func ComplaintSchemaToComplaintDomain(complaint *schema.Complaint) *domain.Complaint {
	return &domain.Complaint{
		ID:          complaint.ID,
		User_ID:     complaint.Admin_ID,
		Category_ID: complaint.Category_ID,
		Admin_ID:    complaint.Admin_ID,
		Title:       complaint.Title,
		Status:      string(complaint.Status),
		Attachment:  complaint.Attachment,
		URL:         complaint.URL,
	}
}

func ComplaintDomaintoComplaintResponse(complaint *domain.Complaint) web.ComplaintResponse {
	return web.ComplaintResponse{
		Title:      complaint.Title,
		Content:    complaint.Title,
		Status:     complaint.Status,
		Attachment: complaint.Attachment,
		URL:        complaint.URL,
	}
}
