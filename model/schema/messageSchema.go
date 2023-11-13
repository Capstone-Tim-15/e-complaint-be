package schema

import (
	"time"
)

type Message struct {
	ID           string    `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Complaint_ID string    `gorm:"column:complaint_id;size:191"`
	Role         role      `gorm:"type:varchar(255)"`
	Message      string
}

type role string

const (
	USER  role = "USER"
	ADMIN role = "ADMIN"
)
