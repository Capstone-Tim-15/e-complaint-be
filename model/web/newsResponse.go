package web

import (
	"ecomplaint/model/domain"
	"time"
)

type NewsResponse struct {
	ID       string            `json:"id"`
	Admin_ID string            `json:"adminId"`
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	Date     time.Time         `json:"date"`
	Feedback []domain.Feedback `json:"feedback"`
	Likes    []domain.Likes    `json:"likes"`
}
