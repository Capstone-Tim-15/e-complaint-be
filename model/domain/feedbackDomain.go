package domain

type Feedback struct {
	ID      string
	User_ID string
	User    User
	News_ID string
	Content string
}
