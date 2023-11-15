package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByName(name string) (*domain.User, error)
	Update(user *domain.User, id string) (*domain.User, error)
	ResetPassword(user *domain.User, email string) (*domain.User, error)
	Delete(id string) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (r *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	var userDb *schema.User

	for {
		userDb = req.UserDomaintoUserSchema(*user)
		userDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&user, userDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := r.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (r *UserRepositoryImpl) FindById(id string) (*domain.User, error) {
	user := domain.User{}

	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	user := []domain.User{}

	result := r.DB.Where("deleted_at IS NULL").Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByName(name string) (*domain.User, error) {
	user := domain.User{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := r.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *domain.User, id string) (*domain.User, error) {
	userDb := req.UserDomaintoUserSchema(*user)

	result := r.DB.Table("users").Where("id = ?", id).Updates(userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (r *UserRepositoryImpl) ResetPassword(user *domain.User, email string) (*domain.User, error) {
	userDb := req.UserDomaintoUserSchema(*user)

	result := r.DB.Table("users").Where("email = ?", email).Updates(userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (r *UserRepositoryImpl) Delete(id string) error {
	result := r.DB.Table("users").Where("id = ?", id).Delete(&schema.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
