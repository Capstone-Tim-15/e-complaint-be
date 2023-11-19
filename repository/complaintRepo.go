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

	complaintDb.URL = "halo"

	result := r.DB.Create(&complaintDb)
	if result.Error != nil {
		return nil, result.Error
	}

	complaint = res.ComplaintSchemaToComplaintDomain(complaintDb)

	return complaint, nil

}
