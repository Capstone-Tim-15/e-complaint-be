package schema

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           string         `gorm:"primaryKey"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CategoryName string
	FAQ          []FAQ `gorm:"ForeignKey:Category_ID;references:ID"`
}
