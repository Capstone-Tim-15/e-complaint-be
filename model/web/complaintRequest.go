package web

type ComplaintCreateRequest struct {
	User_ID     string
	Category_ID string `json:"categoryId" form:"categoryId" validate:"min=6,max=6"`
	Title       string `json:"title" form:"title" validate:"min=1,max=255"`
	Status      string `json:"status" form:"status" validate:"min=1,max=255"`
	Content     string `json:"content" form:"content" validate:"min=1,max=255"`
	ImageUrl    string
}
