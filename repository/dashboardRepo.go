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
	TotalAll() (landing *web.LandingPage, err error)
}

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRepository(DB *gorm.DB) DashboardRepository {
	return &DashboardRepositoryImpl{DB: DB}
}

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
func (repository *DashboardRepositoryImpl) TotalAll() (landing *web.LandingPage, err error) {
	var total *web.LandingPage
	var user int64
	var complaint int64
	var solved int64
	resultUsers := repository.DB.Table("users").Count(&user)
	if resultUsers.Error != nil {
		return nil, resultUsers.Error
	}
	resultComplaint := repository.DB.Table("complaints").Count(&complaint)
	if resultComplaint.Error != nil {
		return nil, resultComplaint.Error
	}
	resultSolved := repository.DB.Table("complaints").Where("status = ?", "solved").Count(&solved)
	if resultSolved.Error != nil {
		return nil, resultSolved.Error
	}
	total = &web.LandingPage{
		TotalUser:      user,
		TotalComplaint: complaint,
		TotalResolved:  solved,
	}

	return total, nil
}
