package schema

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Email     string
	Phone     string
	Password  string
	Complaint []Complaint `gorm:"ForeignKey:Admin_ID;references:ID"`
	News      []News      `gorm:"ForeignKey:Admin_ID;references:ID"`
}
