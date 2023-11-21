package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"

	"gorm.io/gorm"
)

type ComplaintRepository interface {
	Create(complaint *domain.Complaint) (*domain.Complaint, error)
	FindById(id string) (*domain.Complaint, error)
	FindAll() ([]domain.Complaint, error)
}

type ComplaintRepositoryImpl struct {
	DB *gorm.DB
}

func NewComplaintRepository(DB *gorm.DB) ComplaintRepository {
	return &ComplaintRepositoryImpl{DB: DB}
}

func (r *ComplaintRepositoryImpl) Create(complaint *domain.Complaint) (*domain.Complaint, error) {
	var complaintDb *schema.Complaint

	for {
		complaintDb = req.ComplaintDomaintoComplaintSchema(*complaint)
		complaintDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&complaint, complaintDb.ID)
		if result.Error != nil {
			break
		}
	}

	result := r.DB.Create(&complaintDb)
	if result.Error != nil {
		return nil, result.Error
	}

	complaint = res.ComplaintSchemaToComplaintDomain(complaintDb)

	return complaint, nil

}

func (r *ComplaintRepositoryImpl) FindById(id string) (*domain.Complaint, error) {
	complaint := domain.Complaint{}

	result := r.DB.Where("id = ?", id).Preload("Comment").Preload("Category").First(&complaint)
	if result.Error != nil {
		return nil, result.Error
	}

	return &complaint, nil
}

func (r *ComplaintRepositoryImpl) FindAll() ([]domain.Complaint, error) {
	complaint := []domain.Complaint{}

	result := r.DB.Where("deleted_at IS NULL").Preload("Category").Find(&complaint)
	if result.Error != nil {
		return nil, result.Error
	}

	return complaint, nil
}