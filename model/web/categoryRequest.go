package web

type CategoryCreateRequest struct {
	CategoryName string `json:"CategoryName" form:"name" validate:"required,min=1,max=255"`
}
