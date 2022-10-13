package web

type DailyReportRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserId      uint64 `json:"user_id" binding:"required"`
	//Location    string `json:"location" binding:"required"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}

type DailyReportUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserId      uint64 `json:"user_id" binding:"required"`
	// Location    string `json:"location" binding:"required"`
	UpdatedBy string `json:"updated_by" binding:"required"`
	DeletedBy string `json:"deleted_by"`
}
