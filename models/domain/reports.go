package domain

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	gorm.Model
	DailyReport DailyReport `json:"dailyreportid" gorm:"foreignKey:DailyReportId"`
	Name        string      `json:"name" gorm:"type: varchar (200)"`
	Description string      `json:"description" gorm:"type:varchar(255);not null"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
	UserId      uint64      `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string      `json:"updated_by"`
	DeletedBy   string      `json:"deleted_by"`
}
