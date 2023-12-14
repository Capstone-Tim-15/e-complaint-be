package schema

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Admin_ID    string         `gorm:"column:admin_id;size:191"`
	Admin       Admin          `gorm:"foreignKey:Admin_ID;references:ID"`
	Category_ID string         `gorm:"column:category_id;size:191"`
	Category    Category       `gorm:"foreignKey:Category_ID;references:ID"`
	Title       string
	Content     string
	Date        time.Time
	ImageUrl    string
	Feedback    []Feedback `gorm:"ForeignKey:News_ID;references:ID"`
	Likes       []Likes    `gorm:"ForeignKey:News_ID;references:ID"`
}
