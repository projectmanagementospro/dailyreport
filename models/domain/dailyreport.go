package domain

import (
	"time"

	"gorm.io/gorm"
)

type DailyReport struct {
	gorm.Model
	ReportNumber uint      `json:"report_number" gorm:"primaryKey;autoIncrement:true"`
	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
	Description  string    `json:"description" gorm:"type:varchar(255);not null"`
	User_id      uint64    `json:"user_id" gorm:"type:uint;not null"`
	Location     string    `json:"location"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	UpdatedBy    string    `json:"updated_by"`
	DeletedBy    string    `json:"deleted_by"`
}
