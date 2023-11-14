package web

type FeedbackResponse struct {
	ID      string `json:"id"`
	User_ID string `json:"userId"`
	News_ID string `json:"newsId"`
	Content string `json:"content"`
}
