package response

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/model/web"
)

func ComplaintSchemaToComplaintDomain(complaint *schema.Complaint) *domain.Complaint {
	return &domain.Complaint{
		ID:          complaint.ID,
		User_ID:     complaint.User_ID,
		Category_ID: complaint.Category_ID,
		Title:       complaint.Title,
		Content:     complaint.Content,
		Address:     complaint.Address,
		Status:      string(complaint.Status),
		ImageUrl:    complaint.ImageUrl,
		CreatedAt:   complaint.CreatedAt.String(),
	}
}

func ComplaintDomaintoComplaintResponse(complaint *domain.Complaint) web.ComplaintCreateResponse {
	return web.ComplaintCreateResponse{
		ID:       complaint.ID,
		User_ID:  complaint.User_ID,
		Category: complaint.Category_ID,
		Title:    complaint.Title,
		Content:  complaint.Content,
		Address:  complaint.Address,
		Status:   complaint.Status,
		ImageUrl: complaint.ImageUrl,
	}
}

func FindComplaintDomaintoComplaintResponse(complaint *domain.Complaint) web.ComplaintResponse {
	complaintResponse := web.ComplaintResponse{
		ID:        complaint.ID,
		User_ID:   complaint.User_ID,
		Name:      complaint.User.Name,
		Category:  complaint.Category.CategoryName,
		Title:     complaint.Title,
		Content:   complaint.Content,
		Address:   complaint.Address,
		Status:    complaint.Status,
		ImageUrl:  complaint.ImageUrl,
		CreatedAt: complaint.CreatedAt,

		Comment: make([]web.CommentResponse, len(complaint.Comment)),
	}

	for i, comment := range complaint.Comment {
		commentResponse := web.CommentResponse{
			ID:           comment.ID,
			Complaint_ID: comment.Complaint_ID,
			Fullname: 	  comment.Fullname,
			Role:         comment.Role,
			Message:      comment.Message,
		}
		complaintResponse.Comment[i] = commentResponse
	}

	return complaintResponse
}

func FindStatusComplaintDomaintoComplaintResponse(complaints []domain.Complaint) []web.ComplaintResponse {
	var resultComplaintResponse []web.ComplaintResponse
	for _, complaint := range complaints {
		complaintResponse := web.ComplaintResponse{
			ID:         complaint.ID,
			User_ID:    complaint.User_ID,
			Name:       complaint.User.Name,
			PhotoImage: complaint.User.ImageUrl,
			Category:   complaint.Category.CategoryName,
			Title:      complaint.Title,
			Content:    complaint.Content,
			Address:    complaint.Address,
			Status:     complaint.Status,
			ImageUrl:   complaint.ImageUrl,
			CreatedAt:  complaint.CreatedAt,
			UpdateAt:   complaint.UpdatedAt.String(),

			Comment: make([]web.CommentResponse, len(complaint.Comment)),
		}

		for i, comment := range complaint.Comment {
			commentResponse := web.CommentResponse{
				ID:           comment.ID,
				Complaint_ID: comment.Complaint_ID,
				Role:         comment.Role,
				Message:      comment.Message,
			}
			complaintResponse.Comment[i] = commentResponse
		}
		resultComplaintResponse = append(resultComplaintResponse, complaintResponse)
	}
	return resultComplaintResponse
}

func FindCategoryComplaintDomaintoComplaintResponse(complaints []domain.Complaint) []web.ComplaintResponse {
	var resultComplaintResponse []web.ComplaintResponse
	for _, complaint := range complaints {
		complaintResponse := web.ComplaintResponse{
			ID:         complaint.ID,
			User_ID:    complaint.User_ID,
			Name:       complaint.User.Name,
			PhotoImage: complaint.User.ImageUrl,
			Category:   complaint.Category.CategoryName,
			Title:      complaint.Title,
			Content:    complaint.Content,
			Address:    complaint.Address,
			Status:     complaint.Status,
			ImageUrl:   complaint.ImageUrl,

			Comment: make([]web.CommentResponse, len(complaint.Comment)),
		}

		for i, comment := range complaint.Comment {
			commentResponse := web.CommentResponse{
				ID:           comment.ID,
				Complaint_ID: comment.Complaint_ID,
				Role:         comment.Role,
				Message:      comment.Message,
			}
			complaintResponse.Comment[i] = commentResponse
		}
		resultComplaintResponse = append(resultComplaintResponse, complaintResponse)
	}
	return resultComplaintResponse
}

func ConvertComplaintResponse(complaints []domain.Complaint) []web.ComplaintResponse {
	var results []web.ComplaintResponse
	for _, complaint := range complaints {
		complaintResponse := web.ComplaintResponse{
			ID:        complaint.ID,
			User_ID:   complaint.User_ID,
			Name:      complaint.User.Name,
			Category:  complaint.Category.CategoryName,
			Title:     complaint.Title,
			Content:   complaint.Content,
			Address:   complaint.Address,
			Status:    complaint.Status,
			ImageUrl:  complaint.ImageUrl,
			CreatedAt: complaint.CreatedAt,
		}
		results = append(results, complaintResponse)
	}
	return results
}
