package departmentDto

type DepartmentRequest struct {
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`
	IsSub          int64  `json:"is_sub"`
	GmNum          *int64 `json:"gm_num"`
	IsDeleted      int64  `json:"is_deleted"`
}
