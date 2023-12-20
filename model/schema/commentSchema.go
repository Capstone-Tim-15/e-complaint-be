package schema

import (
	"time"
)

type Comment struct {
	ID           string    `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	Complaint_ID string    `gorm:"column:complaint_id;size:191"`
	Fullname     string    `gorm:"type:varchar(255)"`
	Role         Role      `gorm:"type:varchar(255)"`
	Message      string
}

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)
