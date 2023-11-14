package web

import "time"

type NewsResponse struct {
	ID       string    `json:"id"`
	Admin_ID string    `json:"adminId"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
}
