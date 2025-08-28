package departmentDto

import (
	"time"
)

type DepartmentResponse struct {
	ID             *int64             `json:"id"`
	DepartmentCode *string            `json:"department_code"`
	DepartmentName *string            `json:"department_name"`
	IsSub          *int64             `json:"is_sub"`
	GmNum          *int64             `json:"gm_num"`
	IsDeleted      *int64             `json:"is_deleted"`
	DeletedAt      *time.Time         `json:"deleted_at"`
	JobTitle       []JobTitleResponse `json:"job_title"`
}

type JobTitleResponse struct {
	ID            *int64     `json:"id"`
	Name          *string    `json:"name"`
	DepartmentID  *int64     `json:"department_id"`
	SubDept       *int       `json:"sub_dept"`
	Role          *int64     `json:"role"`
	GmNum         *int       `json:"gm_num"`
	IsDaily       *int64     `json:"is_daily"`
	DailyLevel    *int64     `json:"daily_level"`
	IsDeleted     *int64     `json:"is_deleted"`
	DeletedAt     *time.Time `json:"deleted_at"`
	TotalEmployee int64      `json:"total_employee"`
}
