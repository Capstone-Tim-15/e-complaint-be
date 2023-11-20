package schema

import (
	"time"

	"gorm.io/gorm"
)

type Likes struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User_ID   string         `gorm:"column:user_id;size:191"`
	News_ID   string         `gorm:"column:news_id;size:191"`
	Status    string         `gorm:"type:varchar(255)"`
}
