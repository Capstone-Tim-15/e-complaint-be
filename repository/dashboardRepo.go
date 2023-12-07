package repository

import (
	"ecomplaint/model/domain"
	"gorm.io/gorm"
)

type DashboardRepository interface {
	CountUser() (totalUser int64, err error)
	CountComplaint() (totalComplaint int64, err error)
	CountResolved() (totalResolved int64, err error)
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: DB}
}

func (repository *DashboardRepositoryImpl) CountUser() (totalUser int64, err error) {
	var totalCount int64
	result := repository.DB.Model(&domain.User{}).Where("deleted_at IS NULL").Count(&totalCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalCount, nil
}

func (repository *DashboardRepositoryImpl) CountComplaint() (totalComplaint int64, err error) {
	var totalCount int64
	result := repository.DB.Model(&domain.Complaint{}).Where("deleted_at IS NULL").Count(&totalCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalCount, nil
}

func (repository *DashboardRepositoryImpl) CountResolved() (totalResolved int64, err error) {
	var totalCount int64
	result := repository.DB.Model(&domain.Complaint{}).Where("status = ?", "solved").Count(&totalCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalCount, nil
}
