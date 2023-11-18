package web

type FaqRequest struct {
	Title       string `json:"title" validate:"required,min=5,max=50"`
	Content     string `json:"content" validate:"required,min=5,max=255"`
	Category_ID string `json:"categoryId" validate:"required"`
}
