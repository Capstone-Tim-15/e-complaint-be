package schema

type Rooms struct {
	ID           string `gorm:"unique"`
	Name         string
	PhotoProfile string
	LastMessage  string
}