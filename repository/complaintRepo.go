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
	FindAll(page, pageSize int) ([]domain.Complaint, int64, error)
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

func (r *ComplaintRepositoryImpl) FindAll(page, pageSize int) ([]domain.Complaint, int64, error) {
	offset := (page - 1) * pageSize

	complaints := []domain.Complaint{}
	var totalCount int64

	resultCount := r.DB.Model(&domain.Complaint{}).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}

	resultData := r.DB.Where("deleted_at IS NULL").Offset(offset).Limit(pageSize).Find(&complaints)
	if resultData.Error != nil {
		return nil, 0, resultData.Error
	}

	return complaints, totalCount, nil
}
