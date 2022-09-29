package repository

import (
	"dailyreport/models/domain"
	"errors"

	"gorm.io/gorm"
)

type DailyReportRepository interface {
	All() []domain.DailyReport
	Create(d domain.DailyReport) domain.DailyReport
	Update(d domain.DailyReport) domain.DailyReport
	Delete(d domain.DailyReport)
	FindById(id uint) (domain.DailyReport, error)
}

type DailyReportConnection struct {
	connection *gorm.DB
}

func NewDailyReportRepository(connection *gorm.DB) DailyReportRepository {
	return &DailyReportConnection{connection: connection}
}

func (c *DailyReportConnection) All() []domain.DailyReport {
	var dailyreports []domain.DailyReport
	c.connection.Find(&dailyreports)
	return dailyreports
}

func (c *DailyReportConnection) Create(d domain.DailyReport) domain.DailyReport {
	c.connection.Save(&d)
	return d
}

func (c *DailyReportConnection) Update(d domain.DailyReport) domain.DailyReport {
	c.connection.Save(&d)
	return d
}

func (c *DailyReportConnection) Delete(d domain.DailyReport) {
	c.connection.Delete(&d)
}

func (c *DailyReportConnection) FindById(id uint) (domain.DailyReport, error) {
	var dailyreport domain.DailyReport
	c.connection.Find(&dailyreport, "id = ?", id)
	if dailyreport.ID == 0 {
		return dailyreport, errors.New("id not found")
	}
	return dailyreport, nil
}
