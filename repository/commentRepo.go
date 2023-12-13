package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"
	"fmt"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(mess *domain.Comment) (*domain.Comment, error)
	CheckUser(senderId string) (*domain.User, error)
	CheckAdmin(senderId string) (*domain.Admin, error)
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(DB *gorm.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: DB}
}

func (r *CommentRepositoryImpl) Create(mess *domain.Comment) (*domain.Comment, error) {
	var messDb *schema.Comment

	for {
		messDb = req.CommentDomaintoCommentSchema(mess)
		messDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&mess, messDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := r.DB.Create(&messDb)
	if result.Error != nil {
		return nil, result.Error
	}

	mess = res.CommentSchematoCommentDomain(messDb)

	return mess, nil
}

func (r *CommentRepositoryImpl) CheckUser(senderId string) (*domain.User, error) {
	user := domain.User{}
	fmt.Println("masok 1")
	result := r.DB.Where("id = ?", senderId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *CommentRepositoryImpl) CheckAdmin(senderId string) (*domain.Admin, error) {
	admin := domain.Admin{}
	fmt.Println("masok 2")
	result := r.DB.Where("id = ?", senderId).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
