package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(mess *domain.Comment) (*domain.Comment, error)
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessRepository(DB *gorm.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: DB}
}

func (r *CommentRepositoryImpl) Create(mess *domain.Comment) (*domain.Comment, error) {
	var messDb *schema.Comment

	for {
		messDb = req.MessDomaintoMessSchema(mess)
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

	mess = res.MessSchematoMessDomain(messDb)

	return mess, nil
}
