package web

type FeedbackCreateRequest struct {
	User_ID string `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=255"`
}

type FeedbackUpdateRequest struct {
	User_ID string `json:"userId" form:"userId" validate:"required,min=1,max=255"`
	News_ID string `json:"newsId" form:"newsId" validate:"required,min=1,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=255"`
}