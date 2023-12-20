package schema

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID         string         `gorm:"primaryKey"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Fullname   string         `gorm:"type:varchar(255)"`
	Role       Role           `gorm:"type:varchar(255)"`
	PhotoImage string
	News_ID    string `gorm:"column:news_id;size:191"`
	Content    string
}
