package web

type FeedbackResponse struct {
	ID         string `json:"id"`
	Fullname   string `json:"fullname"`
	Role       string `json:"role"`
	PhotoImage string `json:"photoImage"`
	News_ID    string `json:"newsId"`
	Content    string `json:"content"`
}

type FeedbackCreateResponse struct {
	ID         string `json:"id"`
	Fullname   string `json:"fullname"`
	Role       string `json:"role"`
	PhotoImage string `json:"photoImage"`
	News_ID    string `json:"newsId"`
	Content    string `json:"content"`
}
