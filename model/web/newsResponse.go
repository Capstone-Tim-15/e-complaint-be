package web

import (
	"time"
)

type NewsResponse struct {
	ID         string             `json:"id"`
	Admin_ID   string             `json:"adminId"`
	Name       string             `json:"name"`
	PhotoImage string             `json:"photoImage"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	Date       time.Time          `json:"date"`
	Feedback   []FeedbackResponse `json:"feedback"`
	Likes      []LikesResponse    `json:"likes"`
}
