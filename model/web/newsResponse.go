package web

import (
	"time"
)

type NewsResponse struct {
	ID         string             `json:"id"`
	Admin_ID   string             `json:"adminId"`
	Category   string             `json:"category"`
	Name       string             `json:"name"`
	PhotoImage string             `json:"photoImage"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Date       time.Time          `json:"date"`
	ImageUrl   string             `json:"imageUrl"`
	Feedback   []FeedbackResponse `json:"feedback"`
	Likes      []LikesResponse    `json:"likes"`
}

type NewsCreateResponse struct {
	ID       string    `json:"id"`
	Admin_ID string    `json:"adminId"`
	Category string    `json:"category"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	ImageUrl string    `json:"imageUrl"`
}
