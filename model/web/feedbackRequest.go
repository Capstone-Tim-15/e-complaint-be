package web

type FeedbackCreateRequest struct {
	User_ID string `json:"user_id" form:"user_id" validate:"required,min=1,max=255"`
	News_ID string `json:"news_id" form:"news_id" validate:"required,min=1,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=255"`
}

type FeedbackUpdateRequest struct {
	User_ID string `json:"user_id" form:"user_id" validate:"required,min=1,max=255"`
	News_ID string `json:"news_id" form:"news_id" validate:"required,min=1,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=255"`
}
