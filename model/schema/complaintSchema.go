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
	User        User           `gorm:"foreignKey:User_ID;references:ID"`
	Category_ID string         `gorm:"column:category_id;size:191"`
	Category    Category       `gorm:"foreignKey:Category_ID;references:ID"`
	Title       string
	Content     string
	Status      Status `gorm:"type:varchar(255)"`
	ImageUrl    string
	Comment     []Comment `gorm:"ForeignKey:Complaint_ID;references:ID"`
}

type Status string

const (
	SOLVED    Status = "SOLVED"
	UNSOLVED  Status = "UNSOLVED"
	SEND Status = "SEND"
)
