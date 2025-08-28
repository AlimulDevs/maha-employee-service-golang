package departmentModel

import (
	"api/app/models/model/jobTitleModel"
	"time"
)

type DepartmentModel struct {
	ID             *int64                        `json:"id"`
	DepartmentCode *string                       `json:"department_code"`
	DepartmentName *string                       `json:"department_name"`
	IsSub          *int64                        `json:"is_sub"`
	GmNum          *int64                        `json:"gm_num"`
	IsDeleted      *int64                        `json:"is_deleted"`
	DeletedAt      *time.Time                    `json:"deleted_at"`
	JobTitle       []jobTitleModel.JobTitleModel `json:"job_title" gorm:"foreignKey:DepartmentID;references:ID"`
}

func (DepartmentModel) TableName() string {
	return "departments" // Custom table name
}
