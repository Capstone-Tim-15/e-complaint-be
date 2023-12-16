package schema

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID           string         `gorm:"primaryKey"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Complaint_ID string			`gorm:"column:complaint_id;size:191"`
	Complaint    Complaint		`gorm:"foreignKey:Complaint_ID;references:ID"`
	Message      string
	Status       string
}