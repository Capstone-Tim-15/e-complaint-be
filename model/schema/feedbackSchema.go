package schema

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User_ID   string         `gorm:"column:user_id;size:191"`
	User      User           `gorm:"foreignKey:User_ID;references:ID"`
	News_ID   string         `gorm:"column:news_id;size:191"`
	Content   string
}
