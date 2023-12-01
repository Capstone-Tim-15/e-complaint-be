package web

type ComplaintResponse struct {
	ID         string            `json:"id"`
	User_ID    string            `json:"userId"`
	Name       string            `json:"name"`
	PhotoImage string            `json:"photoImage"`
	Category   string            `json:"category"`
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	Status     string            `json:"status"`
	ImageUrl   string            `json:"imageUrl"`
	Comment    []CommentResponse `json:"comment"`
}

type ComplaintCreateResponse struct {
	ID       string `json:"id"`
	User_ID  string `json:"userId"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   string `json:"status"`
	ImageUrl string `json:"imageUrl"`
}
