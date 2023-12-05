package web

type FeedbackResponse struct {
	ID         string `json:"id"`
	User_ID    string `json:"userId"`
	Name       string `json:"name"`
	PhotoImage string `json:"photoImage"`
	News_ID    string `json:"newsId"`
	Content    string `json:"content"`
}
