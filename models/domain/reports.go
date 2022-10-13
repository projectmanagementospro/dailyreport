package domain

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	gorm.Model
	DailyReportId uint        `json:"dailyreport_id"`
	DailyReport   DailyReport `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string      `json:"name" gorm:"type: varchar (200)"`
	LocationId    uint        `json:"location_id"`
	Location      Location    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description   string      `json:"description" gorm:"type:varchar(255);not null"`
	StartTime     time.Time   `json:"start_time"`
	EndTime       time.Time   `json:"end_time"`
	UserId        uint64      `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy     string      `json:"updated_by"`
	DeletedBy     string      `json:"deleted_by"`
}

type Location struct {
	Id   uint   `json:"id" gorm:"primaryKey:autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Lng  int    `json:"lng" gorm:"type:int;not null"`
	Lat  int    `json:"lat" gorm:"type:int;not null"`
}
