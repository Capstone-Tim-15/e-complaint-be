package web

type FaqRequest struct {
	Title       string `json:"title" form:"title" validate:"required,min=1,max=50"`
	Content     string `json:"content" form:"content" validate:"required,min=1,max=255"`
	Category_ID string `json:"category_id" form:"category_id" validate:"required"`
}

type FaqUpdateRequest struct {
	Title       string `json:"title" form:"title" validate:"omitempty,min=1,max=50"`
	Content     string `json:"content" form:"content" validate:"omitempty,min=1,max=255"`
	Category_ID string `json:"category_id" form:"category_id" validate:"omitempty"`
}
