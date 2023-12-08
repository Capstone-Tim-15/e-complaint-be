package repository

import (
	"ecomplaint/model/web"
	"gorm.io/gorm"
	"log"
)

type DashboardRepository interface {
	//CountUser() (totalUser int64, err error)
	CountComplaint(table string) (monthly []web.Monthly, err error)
	CountSolved(table string) (monthly []web.Monthly, err error)
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) DashboardRepository {
	return &DashboardRepositoryImpl{DB: DB}
}

//func (repository *DashboardRepositoryImpl) CountUser() (totalUser int64, err error) {
//	var totalCount int64
//	result := repository.DB.Model(&domain.User{}).Where("deleted_at IS NULL").Count(&totalCount)
//	if result.Error != nil {
//		return 0, result.Error
//	}
//	return totalCount, nil
//}

func (repository *DashboardRepositoryImpl) CountComplaint(table string) (monthly []web.Monthly, err error) {
	var monthCount []web.Monthly
	result := repository.DB.Table(table).Select("DATE_FORMAT(created_at, '%Y-%m') AS month, COUNT(*) AS total").
		Group("month").
		Order("month").Find(&monthCount)
	if result.Error != nil {
		return nil, result.Error
	}

	return monthCount, nil
}

func (repository *DashboardRepositoryImpl) CountSolved(table string) (monthly []web.Monthly, err error) {
	var monthCount []web.Monthly
	log.Println(table)
	result := repository.DB.Table(table).Select("DATE_FORMAT(created_at, '%Y-%m') AS month, COUNT(*) AS total").
		Where("status = ?", "solved").
		Group("month").
		Order("month").Find(&monthCount)
	if result.Error != nil {
		return nil, result.Error
	}

	return monthCount, nil
}

//func (repository *DashboardRepositoryImpl) CountResolved() (totalResolved int64, err error) {
//	var totalCount int64
//	result := repository.DB.Model(&domain.Complaint{}).Where("status = ?", "solved").Count(&totalCount)
//	if result.Error != nil {
//		return 0, result.Error
//	}
//	return totalCount, nil
//}
