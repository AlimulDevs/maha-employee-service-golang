package model

import "time"

// Department model
type Department struct {
	ID       uint64     `json:"id" gorm:"primaryKey"`
	JobTitle []JobTitle `json:"job_title" gorm:"foreignKey:DepartmentID"`
	DepartmentRequest
	IsDeleted bool       `json:"is_deleted" gorm:"column:is_deleted"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at;softDelete"`
}

// TableName overrides the default table name behavior
func (Department) TableName() string {
	return "departments" // Custom table name
}

type DepartmentResponse struct {
	ID             uint64             `json:"id" gorm:"primaryKey"`
	DepartmentCode string             `json:"department_code"  gorm:"column:department_code;type:varchar(50);index:unique,where:deleted_at is null;not null"`
	DepartmentName string             `json:"department_name" gorm:"column:department_name"`
	IsSub          int                `json:"is_sub" gorm:"column:is_sub"`
	GmNum          int                `json:"gm_num" gorm:"column:gm_num"`
	JobTitle       []JobTitleResponse `json:"job_title" gorm:"foreignKey:DepartmentID"`
}

type DepartmentRequest struct {
	DepartmentCode string `json:"department_code" gorm:"column:department_code;size:50;not null;uniqueIndex:idx_department_code_active"`
	DepartmentName string `json:"department_name" gorm:"column:department_name"`
	IsSub          int    `json:"is_sub"`
	GmNum          int    `json:"gm_num"`
}
