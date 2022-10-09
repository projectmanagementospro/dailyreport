package web

import (
	"time"
)

type ReportsRequest struct {
	DailyReportId uint64    `json:"dailyreport_id" binding:"required,numeric"`
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	User_id       uint64    `json:"user_id" binding:"required"`
	StartTime     time.Time `json:"start_time" binding:"required"`
	EndTime       time.Time `json:"end_time" binding:"required"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
}

type ReportsUpdateRequest struct {
	ID            uint      `json:"id" binding:"required"`
	DailyReportId uint64    `json:"dailyreport_id" binding:"required,numeric"`
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	User_id       uint64    `json:"user_id" binding:"required"`
	StartTime     time.Time `json:"start_time" binding:"required"`
	EndTime       time.Time `json:"end_time" binding:"required"`
	UpdatedBy     string    `json:"updated_by" binding:"required"`
	DeletedBy     string    `json:"deleted_by"`
}
