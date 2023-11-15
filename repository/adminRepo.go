package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	FindById(id string) (*domain.Admin, error)
	FindByEmail(email string) (*domain.Admin, error)
	FindAll() ([]domain.Admin, error)
	FindByName(name string) (*domain.Admin, error)
	Update(admin *domain.Admin, id string) (*domain.Admin, error)
	ResetPassword(admin *domain.Admin, email string) (*domain.Admin, error)
	Delete(id string) error
}

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: DB}
}

func (r *AdminRepositoryImpl) Create(admin *domain.Admin) (*domain.Admin, error) {
	var adminDb *schema.Admin

	for {
		adminDb = req.AdminDomaintoAdminSchema(*admin)
		adminDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&admin, adminDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := r.DB.Create(&adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	admin = res.AdminSchemaToAdminDomain(adminDb)

	return admin, nil
}

func (r *AdminRepositoryImpl) FindById(id string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := r.DB.Where("id = ?", id).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) FindByEmail(email string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := r.DB.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (r *AdminRepositoryImpl) FindAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := r.DB.Where("deleted_at IS NULL").Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) FindByName(name string) (*domain.Admin, error) {
	author := domain.Admin{}

	result := r.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&author)

	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func (r *AdminRepositoryImpl) Update(admin *domain.Admin, id string) (*domain.Admin, error) {
	adminDb := req.AdminDomaintoAdminSchema(*admin)

	result := r.DB.Table("admins").Where("id = ?", id).Updates(adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	admin = res.AdminSchemaToAdminDomain(adminDb)

	return admin, nil
}

func (r *AdminRepositoryImpl) ResetPassword(admin *domain.Admin, email string) (*domain.Admin, error) {
	adminDb := req.AdminDomaintoAdminSchema(*admin)

	result := r.DB.Table("admins").Where("email = ?", email).Updates(adminDb)
	if result.Error != nil {
		return nil, result.Error
	}

	admin = res.AdminSchemaToAdminDomain(adminDb)

	return admin, nil
}

func (r *AdminRepositoryImpl) Delete(id string) error {
	result := r.DB.Table("admins").Where("id = ?", id).Delete(&schema.Admin{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
