package jobTitleDto

import "time"

type JobTitleRequest struct {
	ID            *int64     `json:"id"`
	Name          string     `json:"name"`
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
