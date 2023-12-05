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
	CreateOTPUser(otp *domain.OTPUser) (*domain.OTPUser, error)
	FindByUserEmail(email string) (*domain.OTPUser, error)
	FindByUserId(id string) (*domain.OTPUser, error)
	DeleteOTPUser(id string) error
	CreateOTPAdmin(otp *domain.OTPAdmin) (*domain.OTPAdmin, error)
	FindByAdminEmail(email string) (*domain.OTPAdmin, error)
	FindByAdminId(id string) (*domain.OTPAdmin, error)
	DeleteOTPAdmin(id string) error
}

type OTPRepositoryImpl struct {
	DB *gorm.DB
}

func NewOTPRepository(DB *gorm.DB) OTPRepository {
	return &OTPRepositoryImpl{DB: DB}
}

func (r *OTPRepositoryImpl) CreateOTPUser(otp *domain.OTPUser) (*domain.OTPUser, error) {
	var otpDb *schema.OTPUser

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

func (r OTPRepositoryImpl) FindByUserEmail(email string) (*domain.OTPUser, error) {
	otp := domain.OTPUser{}

	result := r.DB.
		Joins("JOIN users ON otp_users.user_id = users.id").
		Where("users.email = ?", email).
		Last(&otp)

	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) FindByUserId(id string) (*domain.OTPUser, error) {
	otp := domain.OTPUser{}

	result := r.DB.Where("user_id = ?", id).First(&otp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) DeleteOTPUser(id string) error {
	result := r.DB.Table("otp_users").Where("id = ?", id).Delete(&schema.OTPUser{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OTPRepositoryImpl) CreateOTPAdmin(otp *domain.OTPAdmin) (*domain.OTPAdmin, error) {
	var otpDb *schema.OTPAdmin

	for {
		otpDb = req.AdminOTPDomaintoOTPSchema(otp)
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

	otp = res.AdminOTPSchemaToOTPDomain(otpDb)

	return otp, nil
}

func (r OTPRepositoryImpl) FindByAdminEmail(email string) (*domain.OTPAdmin, error) {
	otp := domain.OTPAdmin{}

	result := r.DB.
		Joins("JOIN admins ON otp_admins.user_id = admins.id").
		Where("admins.email = ?", email).
		Last(&otp)

	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) FindByAdminId(id string) (*domain.OTPAdmin, error) {
	otp := domain.OTPAdmin{}

	result := r.DB.Where("admin_id = ?", id).First(&otp)
	if result.Error != nil {
		return nil, result.Error
	}

	return &otp, nil
}

func (r *OTPRepositoryImpl) DeleteOTPAdmin(id string) error {
	result := r.DB.Table("otp_admins").Where("id = ?", id).Delete(&schema.OTPAdmin{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
