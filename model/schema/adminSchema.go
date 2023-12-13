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
	Username  string
	Email     string
	Phone     string
	Password  string
	ImageUrl  string
	OTP       OTPAdmin `gorm:"ForeignKey:Admin_ID;references:ID"`
	News      []News   `gorm:"ForeignKey:Admin_ID;references:ID"`
}
