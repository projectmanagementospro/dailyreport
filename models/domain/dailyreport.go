package domain

import (
	"gorm.io/gorm"
)

type DailyReport struct {
	gorm.Model
	ReportsId   uint64 `json:"reports_id" gorm:"type:uint;not null, unsigned"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	UserId      uint64 `json:"user_id" gorm:"type:uint;not null"`
	Location    string `json:"location"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}
