package schema

type Message struct {
	ID           uint `gorm:"primaryKey"`
	RoomID       string
	Message      string
}