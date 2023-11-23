package web

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
}
