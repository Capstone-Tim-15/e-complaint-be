package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"

	"gorm.io/gorm"
)

type OTPRepository interface {
	Create(otp *domain.OTP) (*domain.OTP, error)
	FindByEmail(email string) (*domain.OTP, error)
	FindByUserId(id string) (*domain.OTP, error)
	Delete(id string) error
}

type OTPRepositoryImpl struct {
	DB *gorm.DB
}

func NewOTPRepository(DB *gorm.DB) OTPRepository {
	return &OTPRepositoryImpl{DB: DB}
}

func (r *OTPRepositoryImpl) Create(otp *domain.OTP) (*domain.OTP, error) {
	var otpDb *schema.OTP

	for {
		otpDb = req.OTPDomaintoOTPSchema(otp)
		otpDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&otp, otpDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := r.DB.Create(&otpDb)
	if result.Error != nil {
		return nil, result.Error
	}

	otp = res.OTPSchemaToOTPDomain(otpDb)

	return otp, nil
}

func (r OTPRepositoryImpl) FindByEmail(email string) (*domain.OTP, error) {
	otp := domain.OTP{}

	result := r.DB.
		Joins("JOIN users ON otp.user_id = users.id").
		Where("users.email = ?", email).
		Last(&otp)

	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) FindByUserId(id string) (*domain.OTP, error) {
	otp := domain.OTP{}

	result := r.DB.Where("user_id = ?", id).First(&otp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) Delete(id string) error {
	result := r.DB.Table("otps").Where("id = ?", id).Delete(&schema.OTP{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
