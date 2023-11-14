package domain

import "time"

type News struct {
	ID       string
	Admin_ID string
	Title    string
	Content  string
	Date     time.Time
	Feedback []Feedback
	Likes    []Likes
}