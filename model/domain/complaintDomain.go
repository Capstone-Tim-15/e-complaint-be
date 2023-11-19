package domain

type Complaint struct {
	ID          string
	User_ID     string
	Category_ID string
	Admin_ID    string
	Title       string
	Content     string
	Status      string
	Attachment  string
	URL         string
	Message     []Message
}
