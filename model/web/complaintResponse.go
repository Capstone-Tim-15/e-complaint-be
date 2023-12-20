package web

import "time"

type ComplaintResponse struct {
	ID         string            `json:"id"`
	User_ID    string            `json:"userId"`
	Name       string            `json:"name"`
	PhotoImage string            `json:"photoImage"`
	Category   string            `json:"category"`
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	Address    string            `json:"address"`
	Status     string            `json:"status"`
	ImageUrl   string            `json:"imageUrl"`
	CreatedAt  string            `json:"createdAt"`
	UpdateAt   string            `json:"updateAt"`
	Comment    []CommentResponse `json:"comment"`
}

type ComplaintCreateResponse struct {
	ID       string `json:"id"`
	User_ID  string `json:"userId"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	ImageUrl string `json:"imageUrl"`
}

type ComplaintRealtimeResponse struct {
	ID        string           `json:"id"`
	User      UserResponse     `json:"user"`
	Category  CategoryResponse `json:"category"`
	Title     string           `json:"title"`
	Content   string           `json:"content"`
	Address   string           `json:"address"`
	Status    string           `json:"status"`
	ImageUrl  string           `json:"imageUrl"`
	CreatedAt string
	UpdatedAt time.Time
	Comment   []CommentResponse `json:"comment"`
}
