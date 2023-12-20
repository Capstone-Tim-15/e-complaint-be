package domain

import "time"

type News struct {
	ID          string
	Admin_ID    string
	Admin       Admin
	Category_ID string
	Category    Category
	Title       string
	Content     string
	Date        time.Time
	ImageUrl    string
	Feedback    []Feedback
	Likes       []Likes
}
