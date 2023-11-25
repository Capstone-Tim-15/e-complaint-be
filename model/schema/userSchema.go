package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Username  string
	Email     string
	Phone     string
	Password  string
	ImageUrl  string
	OTP       OTP         `gorm:"ForeignKey:User_ID;references:ID"`
	Complaint []Complaint `gorm:"ForeignKey:User_ID;references:ID"`
	Feedback  []Feedback  `gorm:"ForeignKey:User_ID;references:ID"`
	Likes     []Likes     `gorm:"ForeignKey:User_ID;references:ID"`
}
