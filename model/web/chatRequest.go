package web

type Message struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Content  string `json:"content"`
}
