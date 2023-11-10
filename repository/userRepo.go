package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/req"
	"ecomplaint/utils/res"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindById(id int) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]domain.User, error)
	FindByName(name string) (*domain.User, error)
	Update(user *domain.User, id int) (*domain.User, error)
	ResetPassword(user *domain.User, email string) (*domain.User, error)
	Delete(id int) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	userDb := req.UserDomaintoUserSchema(*user)
	result := repository.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (repository *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	user := domain.User{}

	result := repository.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	user := []domain.User{}

	result := repository.DB.Where("deleted_at IS NULL").Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByName(name string) (*domain.User, error) {
	user := domain.User{}

	// Menggunakan query LIKE yang tidak case-sensitive
	result := repository.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Update(user *domain.User, id int) (*domain.User, error) {
	userDb := req.UserDomaintoUserSchema(*user)

	result := repository.DB.Table("users").Where("id = ?", id).Updates(userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (repository *UserRepositoryImpl) ResetPassword(user *domain.User, email string) (*domain.User, error) {
	userDb := req.UserDomaintoUserSchema(*user)

	result := repository.DB.Table("users").Where("email = ?", email).Updates(userDb)
	if result.Error != nil {
		return nil, result.Error
	}

	user = res.UserSchemaToUserDomain(userDb)

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("users").Where("id = ?", id).Delete(&schema.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
