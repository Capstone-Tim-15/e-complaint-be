package schema

import (
	"time"

	"gorm.io/gorm"
)

type OTPUser struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User_ID   string         `gorm:"column:user_id;size:191"`
	OTP       string
}

type OTPAdmin struct {
	ID        string         `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User_ID   string         `gorm:"column:user_id;size:191"`
	OTP       string
}
