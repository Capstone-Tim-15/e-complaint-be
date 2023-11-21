package domain

type Complaint struct {
	ID          string
	User_ID     string
	Category_ID string
	Category    Category
	Title       string
	Content     string
	Status      string
	ImageUrl  string
	Comment     []Comment
}
