package web

import (
	"time"
)

type DailyReportRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	User_id     uint64    `json:"user_id" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
}

type DailyReportUpdateRequest struct {
	ID uint `json:"id" binding:"required"`
	// ReportNumber uint      `json:"report_number" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	User_id     uint64    `json:"user_id" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	UpdatedBy   string    `json:"updated_by" binding:"required"`
	DeletedBy   string    `json:"deleted_by"`
}
