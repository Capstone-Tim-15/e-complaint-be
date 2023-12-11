package web

type FeedbackCreateRequest struct {
	User_ID string
	News_ID string `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=255"`
}

type FeedbackUpdateRequest struct {
	News_ID string `json:"newsId" form:"newsId" validate:"omitempty,min=6,max=6"`
	Content string `json:"content" form:"content" validate:"omitempty,min=1,max=255"`
}
