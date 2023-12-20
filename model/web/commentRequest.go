package web

type CommentCreateRequest struct {
	Complaint_ID string `json:"complaintId" form:"complaintId" validate:"required,min=6,max=6"`
	Fullname     string `json:"fullname" form:"fullname" validate:"omitempty,min=1,max=255"`
	Role         string `json:"role" form:"role" validate:"omitempty,min=1,max=255"`
	Message      string `json:"message" form:"message" validate:"required,min=1,max=255"`
}
