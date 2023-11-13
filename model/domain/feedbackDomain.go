package domain

type Feedback struct {
	ID      string
	User_ID User
	News_ID string
	Content string
}
