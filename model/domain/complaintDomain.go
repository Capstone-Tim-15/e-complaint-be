package domain

import "time"

type Complaint struct {
	ID          string
	User_ID     string
	User        User
	Category_ID string
	Category    Category
	Title       string
	Content     string
	Address     string
	Status      string
	ImageUrl    string
	CreatedAt   string
	UpdatedAt   time.Time
	Comment     []Comment
}
