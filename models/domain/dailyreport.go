package domain

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DailyReport struct {
	gorm.Model
	ReportNumber uint64         `json:"report_number" gorm:"type:auto_increment"`
	Name         string         `json:"name" gorm:"type:varchar(255);not null"`
	Description  string         `json:"description" gorm:"type:varchar(255);not null"`
	User_id      uint64         `json:"user_id" gorm:"type:uint;not null"`
	Location     string         `json:"location"`
	Date         datatypes.Date `json:"date"`
	StartTime    time.Time      `json:"start_time"`
	EndTime      time.Time      `json:"end_time"`
	Duration     time.Time      `json:"duration,omitempty"`
	UpdatedBy    string         `json:"updated_by"`
	DeletedBy    string         `json:"deleted_by"`
}
