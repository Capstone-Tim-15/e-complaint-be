package domain

type Category struct {
	ID        string
	Name      string
	FAQ       []FAQ
	Complaint []Complaint
}
