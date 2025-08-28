package employeeDto

import (
	"time"
)

type EmployeeResponse struct {
	ID          uint64 `json:"id"`
	Nik         string `json:"nik"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	JobTitleID  int    `json:"job_title_id"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
	// PhotoEmbedding           string      `json:"photo_embedding"`
	DepartmentID             int         `json:"department_id"`
	BranchCode               string      `json:"branch_code"`
	Password                 string      `json:"password"`
	Signature                string      `json:"signature"`
	IntegrityPactNum         string      `json:"integrity_pact_num"`
	IntegrityPactCheck       int         `json:"integrity_pact_check"`
	IntegrityPactCheckDate   *time.Time  `json:"integrity_pact_check_date"`
	StatementLetterCheck     int         `json:"statement_letter_check"`
	StatementLetterCheckDate *time.Time  `json:"statement_letter_check_date"`
	ContractID               int         `json:"contract_id"`
	OldContractID            int         `json:"old_contract_id"`
	EmployeeStatus           string      `json:"employee_status"`
	Salary                   float64     `json:"salary"`
	ShowContract             int         `json:"show_contract"`
	EmployeeLetterCode       string      `json:"employee_letter_code"`
	BiodataConfirm           int         `json:"biodata_confirm"`
	BiodataConfirmDate       *time.Time  `json:"biodata_confirm_date"`
	CurrentAddress           string      `json:"current_address"`
	BankAccountNumber        string      `json:"bank_account_number"`
	RoleID                   int         `json:"role_id"`
	Status                   int         `json:"status"`
	StatementRejected        string      `json:"statement_rejected"`
	IsDaily                  int         `json:"is_daily"`
	IsFlexibleAbsent         int         `json:"is_flexible_absent"`
	IsOvertime               int         `json:"is_overtime"`
	OvertimeLimit            int         `json:"overtime_limit"`
	FlexibleAbsentEndDate    string      `json:"flexible_absent_end_date"`
	DeviceToken              string      `json:"device_token"`
	EmailVerifiedAt          *time.Time  `json:"email_verified_at"`
	StatusContractUpdate     int         `json:"status__contract_update"`
	PhotoURL                 string      `json:"photo_url"`
	SignatureURL             interface{} `json:"signature_url"`
	StatusLabel              string      `json:"status_label"`
	IsBpjs                   bool        `json:"is_bpjs"`
	IsBpjsContribution       bool        `json:"is_bpjs_contribution"`
	BpjsWages                string      `json:"bpjs_wages"`
	BpjsContributionWages    string      `json:"bpjs_contribution_wages"`
	JobTitle                 JobTitle    `json:"job_title"`
	Department               Department  `json:"department"`
	Branch                   Branch      `json:"branch"`
}

type Branch struct {
	ID               int64       `json:"id"`
	BranchCode       string      `json:"branch_code"`
	BranchLetterCode string      `json:"branch_letter_code"`
	BranchName       string      `json:"branch_name"`
	BranchTime       string      `json:"branch_time"`
	BranchLocation   string      `json:"branch_location"`
	BranchRadius     int64       `json:"branch_radius"`
	IsProject        int64       `json:"is_project"`
	IsSub            int64       `json:"is_sub"`
	BranchParentCode interface{} `json:"branch_parent_code"`
	Meal             int64       `json:"meal"`
	IsActive         int64       `json:"is_active"`
	IsDeleted        int64       `json:"is_deleted"`
	DeletedAt        interface{} `json:"deleted_at"`
}

type Department struct {
	ID             int64       `json:"id"`
	DepartmentCode string      `json:"department_code"`
	DepartmentName string      `json:"department_name"`
	IsSub          int64       `json:"is_sub"`
	GmNum          int64       `json:"gm_num"`
	IsDeleted      int64       `json:"is_deleted"`
	DeletedAt      interface{} `json:"deleted_at"`
}

type JobTitle struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	DepartmentID  int64       `json:"department_id"`
	SubDept       interface{} `json:"sub_dept"`
	Role          int64       `json:"role"`
	GmNum         interface{} `json:"gm_num"`
	IsDaily       int64       `json:"is_daily"`
	DailyLevel    int64       `json:"daily_level"`
	IsDeleted     int64       `json:"is_deleted"`
	DeletedAt     interface{} `json:"deleted_at"`
	TotalEmployee int64       `json:"total_employee"`
}
