package repository

import (
	"dailyreport/models/domain"
	"errors"

	"gorm.io/gorm"
)

type ReportsRepository interface {
	All() []domain.Reports
	Create(report domain.Reports) domain.Reports
	Update(report domain.Reports) domain.Reports
	Delete(report domain.Reports)
	FindById(id uint) (domain.Reports, error)
	IsDReportExist(id uint) (domain.DailyReport, error)
}

type ReportsConnection struct {
	dbConnect *gorm.DB
}

func NewReportsRepository(db *gorm.DB) ReportsRepository {
	return &ReportsConnection{dbConnect: db}
}

func (conn *ReportsConnection) All() []domain.Reports {
	var reports []domain.Reports
	conn.dbConnect.Preload("DailyReport").Find(&reports)
	return reports
}

func (conn *ReportsConnection) Create(reports domain.Reports) domain.Reports {
	conn.dbConnect.Save(&reports)
	return reports
}

func (conn *ReportsConnection) Update(reports domain.Reports) domain.Reports {
	conn.dbConnect.Omit("created_at").Save(&reports)
	return reports
}

func (conn *ReportsConnection) Delete(reports domain.Reports) {
	conn.dbConnect.Delete(&reports)
}

func (conn *ReportsConnection) FindById(id uint) (domain.Reports, error) {
	var reports domain.Reports
	conn.dbConnect.Preload("DailyReport").Find(&reports, "id = ?", id)
	if reports.ID == 0 {
		return reports, errors.New("id not found")
	}
	return reports, nil
}

func (conn *ReportsConnection) IsDReportExist(id uint) (domain.DailyReport, error) {
	var reports domain.DailyReport
	conn.dbConnect.Preload("DailyReport").Find(&reports, "id = ?", id)
	if reports.ID == 0 {
		return reports, errors.New("DailyReport id haven't been created yet")
	}
	return reports, nil
}
