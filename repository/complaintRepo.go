package repository

import (
	"ecomplaint/model/domain"
	"ecomplaint/model/schema"
	"ecomplaint/utils/helper"
	req "ecomplaint/utils/request"
	res "ecomplaint/utils/response"
	"log"

	"gorm.io/gorm"
)

type ComplaintRepository interface {
	Create(complaint *domain.Complaint) (*domain.Complaint, error)
	FindById(id, role string) (*domain.Complaint, error)
	FindByStatusUser(status, id string) ([]domain.Complaint, error)
	FindByCategory(category string, limit int64) ([]domain.Complaint, int64, error)
	FindByStatus(status string, page, pageSize int) ([]domain.Complaint, int64, error)
	FindAllUser(id string) ([]domain.Complaint, error)
	FindAll(page, pageSize int) ([]domain.Complaint, int64, error)
	Update(complaint *domain.Complaint, id string) (*domain.Complaint, error)
	Delete(complaint *domain.Complaint, id string) error
}

type ComplaintRepositoryImpl struct {
	DB *gorm.DB
}

func NewComplaintRepository(DB *gorm.DB) ComplaintRepository {
	return &ComplaintRepositoryImpl{DB: DB}
}

func (r *ComplaintRepositoryImpl) Create(complaint *domain.Complaint) (*domain.Complaint, error) {
	var complaintDb *schema.Complaint
	var notificationDb *schema.Notification

	notificationDb = &schema.Notification{}

	for {
		complaintDb = req.ComplaintDomaintoComplaintSchema(*complaint)
		complaintDb.ID = helper.GenerateRandomString()
		notificationDb.ID = helper.GenerateRandomString()

		result := r.DB.First(&complaint, complaintDb.ID)
		notifresult := r.DB.First(&schema.Notification{}, notificationDb.ID)
		if result.Error != nil && notifresult != nil {
			break
		}
	}
	result := r.DB.Create(&complaintDb)
	if result.Error != nil {
		return nil, result.Error
	}
	complaint = res.ComplaintSchemaToComplaintDomain(complaintDb)
	
	var getComplaint *schema.Complaint
	r.DB.Preload("User").Preload("Category").First(&getComplaint, "id = ? ",complaintDb.ID)
	log.Println(getComplaint.Category)
	notificationDb.Complaint_ID = complaintDb.ID
	notificationDb.Message = "Complaint has been created by " + getComplaint.User.Name + " with category " + getComplaint.Category.CategoryName
	notificationDb.Status = "unread"

	r.DB.Create(notificationDb)

	return complaint, nil

}

func (r *ComplaintRepositoryImpl) FindById(id, role string) (*domain.Complaint, error) {
	complaint := domain.Complaint{}
	if role == "admin" {
		result := r.DB.Where("id = ?", id).Preload("Comment", func(DB *gorm.DB) *gorm.DB{return DB.Order("created_at ASC")}).Preload("Category").Preload("User").First(&complaint)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := r.DB.Where("id = ?", id).Preload("Comment", func(DB *gorm.DB) *gorm.DB{return DB.Order("created_at DESC")}).Preload("Category").Preload("User").First(&complaint)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	var notification *schema.Notification
	notification = &schema.Notification{
		Status: "read",
	}
	r.DB.Model(&schema.Notification{}).Where("complaint_id = ?", id).Updates(notification)

	return &complaint, nil
}

func (r *ComplaintRepositoryImpl) FindByStatusUser(status, id string) ([]domain.Complaint, error) {
	complaint := []domain.Complaint{}

	result := r.DB.Debug().Where("status = ? AND user_id = ?", status, id).Where("deleted_at IS NULL").Preload("Comment").Preload("Category").Preload("User").Order("updated_at desc").Find(&complaint)
	if result.Error != nil {
		return nil, result.Error
	}

	return complaint, nil
}

func (r *ComplaintRepositoryImpl) FindByCategory(category string, limit int64) ([]domain.Complaint, int64, error) {
	complaint := []domain.Complaint{}
	var totalCount int64

	resultCount := r.DB.Model(&domain.Complaint{}).Where("category_id = ?", category).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil{
		return nil, 0, resultCount.Error
	}

	result := r.DB.Debug().Where("category_id = ?", category).Preload("Comment").Preload("Category").Preload("User").Order("created_at desc").Limit(int(limit)).Find(&complaint)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return complaint, totalCount, nil
}

func (r *ComplaintRepositoryImpl) FindByStatus(status string, page, pageSize int) ([]domain.Complaint, int64, error) {
	offset := (page - 1) * pageSize
	complaint := []domain.Complaint{}
	var totalCount int64

	resultCount := r.DB.Model(&domain.Complaint{}).Where("status = ?", status).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil{
		return nil, 0, resultCount.Error
	}

	result := r.DB.Debug().Where("status = ? AND deleted_at IS NULL", status).Preload("Comment").Preload("Category").Preload("User").Offset(offset).Limit(pageSize).Order("created_at desc").Find(&complaint)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return complaint, totalCount, nil
}

func (r *ComplaintRepositoryImpl) FindAllUser(id string) ([]domain.Complaint, error) {
	complaints := []domain.Complaint{}

	result := r.DB.Where("user_id = ?", id).Where("deleted_at IS NULL").Preload("Category").Preload("User").Order("updated_at desc").Find(&complaints)
	if result.Error != nil {
		return nil, result.Error
	}

	return complaints, nil
}

func (r *ComplaintRepositoryImpl) FindAll(page, pageSize int) ([]domain.Complaint, int64, error) {
	offset := (page - 1) * pageSize

	complaints := []domain.Complaint{}
	var totalCount int64

	resultCount := r.DB.Model(&domain.Complaint{}).Where("deleted_at IS NULL").Count(&totalCount)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}

	resultData := r.DB.Where("deleted_at IS NULL").Preload("Category").Preload("User").Offset(offset).Limit(pageSize).Order("created_at desc").Find(&complaints)

	if resultData.Error != nil {
		return nil, 0, resultData.Error
	}

	return complaints, totalCount, nil
}

func (r *ComplaintRepositoryImpl) Update(complaint *domain.Complaint, id string) (*domain.Complaint, error) {
	complaintDb := req.ComplaintDomaintoComplaintSchema(*complaint)

	result := r.DB.Table("complaints").Where("id", id).Updates(complaintDb)
	if result.Error != nil {
		return nil, result.Error
	}

	complaint = res.ComplaintSchemaToComplaintDomain(complaintDb)

	return complaint, nil
}

func (r *ComplaintRepositoryImpl) Delete(complaint *domain.Complaint, id string) error {
	complaintDb := req.ComplaintDomaintoComplaintSchema(*complaint)

	result := r.DB.Where("id", id).Delete(&complaintDb)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
