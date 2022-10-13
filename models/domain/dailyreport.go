package domain

import (
	"gorm.io/gorm"
)

type DailyReport struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	UserId      uint64 `json:"user_id" gorm:"type:uint;not null"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}

// type Location struct {
// 	name string
// 	X, Y int
// }

// func (loc Location) GormDataType() string {
// 	return "geometry"
// }
