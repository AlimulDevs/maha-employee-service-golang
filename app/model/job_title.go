package model

import (
	"time"
)

// JobTitle model
type JobTitle struct {
	ID           uint64     `json:"id"`
	Name         string     `json:"name"`
	DepartmentID uint64     `json:"department_id"`
	SubDept      *string    `json:"sub_dept,omitempty"`
	Role         int        `json:"role"`
	GmNum        *int       `json:"gm_num,omitempty"`
	IsDaily      int        `json:"is_daily"`
	DailyLevel   int        `json:"daily_level"`
	IsDeleted    bool       `json:"is_deleted"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// TableName overrides the default table name behavior
func (JobTitle) TableName() string {
	return "job_titles" // Custom table name
}

// JobTitleResponse struct used for the API response
type JobTitleResponse struct {
	ID           uint64  `json:"id"`
	Name         string  `json:"name"`
	DepartmentID uint64  `json:"department_id"`
	SubDept      *string `json:"sub_dept,omitempty"`
	Role         int     `json:"role"`
	GmNum        *int    `json:"gm_num,omitempty"`
	IsDaily      int     `json:"is_daily"`
	DailyLevel   int     `json:"daily_level"`
}

// JobTitleRequest struct used for API requests (e.g., when creating or updating a job title)
type JobTitleRequest struct {
	Name         string  `json:"name"`
	DepartmentID uint64  `json:"department_id"`
	SubDept      *string `json:"sub_dept,omitempty"`
	Role         int     `json:"role"`
	GmNum        *int    `json:"gm_num,omitempty"`
	IsDaily      int     `json:"is_daily"`
	DailyLevel   int     `json:"daily_level"`
}
