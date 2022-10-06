package repository

import (
	"dailyreport/models/domain"
	"errors"

	"gorm.io/gorm"
)

type DailyReportRepository interface {
	All() []domain.DailyReport
	Create(dailyreport domain.DailyReport) domain.DailyReport
	Update(dailyreport domain.DailyReport) domain.DailyReport
	Delete(dailyreport domain.DailyReport)
	FindById(id uint) (domain.DailyReport, error)
}

type DailyReportConnection struct {
	dbConnect *gorm.DB //connect to database
}

func NewDailyReportRepository(db *gorm.DB) DailyReportRepository {
	return &DailyReportConnection{dbConnect: db} //connect database to interface
}

func (conn *DailyReportConnection) All() []domain.DailyReport {
	var dailyreports []domain.DailyReport
	conn.dbConnect.Find(&dailyreports)
	return dailyreports
}

func (conn *DailyReportConnection) Create(dailyreport domain.DailyReport) domain.DailyReport {
	conn.dbConnect.Save(&dailyreport)
	return dailyreport
}

func (conn *DailyReportConnection) Update(dailyreport domain.DailyReport) domain.DailyReport {
	conn.dbConnect.Omit("created_at").Save(&dailyreport)
	return dailyreport
}

func (conn *DailyReportConnection) Delete(dailyreport domain.DailyReport) {
	conn.dbConnect.Delete(&dailyreport)
}

func (conn *DailyReportConnection) FindById(id uint) (domain.DailyReport, error) {
	var dailyreport domain.DailyReport
	conn.dbConnect.Find(&dailyreport, "id = ?", id)
	if dailyreport.ID == 0 {
		return dailyreport, errors.New("id not found")
	}
	return dailyreport, nil
}
