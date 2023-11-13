package schema

import (
	"time"

	"gorm.io/gorm"
)

type Complaint struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	User_ID     string         `gorm:"column:user_id;size:191"`
	Category_ID string         `gorm:"column:category_id;size:191"`
	Admin_ID    string         `gorm:"column:admin_id;size:191"`
	Title       string
	Status      status `gorm:"type:varchar(255)"`
	Attachment  string
	URL         string
	Message     []Message `gorm:"ForeignKey:Complaint_ID;references:ID"`
}

type status string

const (
	SOLVED    status = "SOLVED"
	UNSOLVED  status = "UNSOLVED"
	CANCELLED status = "CANCELLED"
)
