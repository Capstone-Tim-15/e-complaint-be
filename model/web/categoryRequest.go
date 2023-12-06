package web

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=3,max=255"`
}
