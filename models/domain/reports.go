package domain

import (
	"time"

	"gorm.io/gorm"
)

type Reports struct {
	gorm.Model
	Name        string    `json:"name" gorm:"type: varchar (200)"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
