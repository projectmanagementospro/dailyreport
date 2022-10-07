package repository

import (
	"dailyreport/models/domain"
	"errors"

	"gorm.io/gorm"
)

type DailyReportRepository interface {
	All() []domain.DailyReport
	Create(dReport domain.DailyReport) domain.DailyReport
	Update(dReport domain.DailyReport) domain.DailyReport
	Delete(dReport domain.DailyReport)
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

func (conn *DailyReportConnection) Create(dReport domain.DailyReport) domain.DailyReport {
	conn.dbConnect.Save(&dReport)
	return dReport
}

func (conn *DailyReportConnection) Update(dReport domain.DailyReport) domain.DailyReport {
	conn.dbConnect.Omit("created_at").Save(&dReport)
	return dReport
}

func (conn *DailyReportConnection) Delete(dReport domain.DailyReport) {
	conn.dbConnect.Delete(&dReport)
}

func (conn *DailyReportConnection) FindById(id uint) (domain.DailyReport, error) {
	var dReport domain.DailyReport
	conn.dbConnect.Find(&dReport, "id = ?", id)
	if dReport.ID == 0 {
		return dReport, errors.New("id not found")
	}
	return dReport, nil
}
