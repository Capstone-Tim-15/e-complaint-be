package schema

import (
	"time"

	"gorm.io/gorm"
)

type FAQ struct {
	ID          string         `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Content     string
	Category_ID string `gorm:"column:category_id;size:191"`
}
