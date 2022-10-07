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
}

type labelConnection struct {
	connection *gorm.DB
}

func NewReportsRepository(connection *gorm.DB) ReportsRepository {
	return &DataConnection{connection: connection}
}

func (c *labelConnection) All() []domain.Reports {
	var another []domain.Reports
	c.connection.Find(&another)
	return another
}

func (c *labelConnection) Create(d domain.Reports) domain.Reports {
	c.connection.Save(&d)
	return d
}

func (c *labelConnection) Update(d domain.Reports) domain.Reports {
	c.connection.Omit("created_at").Save(&d)
	return d
}

func (c *labelConnection) Delete(d domain.Reports) {
	c.connection.Delete(&d)
}

func (c *labelConnection) FindById(id uint) (domain.Reports, error) {
	var another domain.Reports
	c.connection.Find(&another, "id = ?", id)
	if another.ID == 0 {
		return another, errors.New("id not found")
	}
	return another, nil
}
