package model

import "time"

// Employee model
type Employee struct {
	ID                       uint64     `json:"id"`
	Nik                      string     `json:"nik"`
	Fullname                 string     `json:"fullname"`
	Email                    string     `json:"email"`
	JobTitleID               int        `json:"job_title_id"`
	PhoneNumber              string     `json:"phone_number"`
	Photo                    string     `json:"photo"`
	PhotoEmbedding           string     `json:"photo_embedding"`
	DepartmentID             int        `json:"department_id"`
	BranchCode               string     `json:"branch_code"`
	Password                 string     `json:"password"`
	Signature                string     `json:"signature"`
	IntegrityPactNum         string     `json:"integrity_pact_num"`
	IntegrityPactCheck       int        `json:"integrity_pact_check"`
	IntegrityPactCheckDate   *time.Time `json:"integrity_pact_check_date"`
	StatementLetterCheck     int        `json:"statement_letter_check"`
	StatementLetterCheckDate *time.Time `json:"statement_letter_check_date"`
	ContractID               int        `json:"contract_id"`
	OldContractID            int        `json:"old_contract_id"`
	EmployeeStatus           string     `json:"employee_status"`
	Salary                   float64    `json:"salary"`
	ShowContract             int        `json:"show_contract"`
	EmployeeLetterCode       string     `json:"employee_letter_code"`
	BiodataConfirm           int        `json:"biodata_confirm"`
	BiodataConfirmDate       *time.Time `json:"biodata_confirm_date"`
	CurrentAddress           string     `json:"current_address"`
	BankAccountNumber        string     `json:"bank_account_number"`
	RoleID                   int        `json:"role_id"`
	Status                   int        `json:"status"`
	StatementRejected        string     `json:"statement_rejected"`
	IsDaily                  int        `json:"is_daily"`
	IsFlexibleAbsent         int        `json:"is_flexible_absent"`
	IsOvertime               int        `json:"is_overtime"`
	OvertimeLimit            int        `json:"overtime_limit"`
	FlexibleAbsentEndDate    string     `json:"flexible_absent_end_date"`
	DeviceToken              string     `json:"device_token"`
	EmailVerifiedAt          *time.Time `json:"email_verified_at"`
	StatusContractUpdate     int        `json:"status__contract_update"`
}

// TableName overrides the default table name behavior
func (Employee) TableName() string {
	return "employees" // Custom table name
}

type EmployeeGetAllResponse struct {
	ID           uint64 `json:"id"`
	Nik          string `json:"nik"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	JobTitleID   int    `json:"job_title_id"`
	PhoneNumber  string `json:"phone_number"`
	Photo        string `json:"photo"`
	DepartmentID int    `json:"department_id"`
	BranchCode   string `json:"branch_code"`
	Signature          string  `json:"signature"`
	ContractID         int     `json:"contract_id"`
	OldContractID      int     `json:"old_contract_id"`
	EmployeeStatus     string  `json:"employee_status"`
	Salary             float64 `json:"salary"`
	ShowContract       int     `json:"show_contract"`
	EmployeeLetterCode string  `json:"employee_letter_code"`
}

type EmployeeRequest struct {
	Nik                      string     `json:"nik"`
	Fullname                 string     `json:"fullname"`
	Email                    string     `json:"email"`
	JobTitleID               int        `json:"job_title_id"`
	PhoneNumber              string     `json:"phone_number"`
	Photo                    string     `json:"photo"`
	PhotoEmbedding           string     `json:"photo_embedding"`
	DepartmentID             int        `json:"department_id"`
	BranchCode               string     `json:"branch_code"`
	Password                 string     `json:"password"`
	Signature                string     `json:"signature"`
	IntegrityPactNum         string     `json:"integrity_pact_num"`
	IntegrityPactCheck       int        `json:"integrity_pact_check"`
	IntegrityPactCheckDate   *time.Time `json:"integrity_pact_check_date"`
	StatementLetterCheck     int        `json:"statement_letter_check"`
	StatementLetterCheckDate *time.Time `json:"statement_letter_check_date"`
	ContractID               int        `json:"contract_id"`
	OldContractID            int        `json:"old_contract_id"`
	EmployeeStatus           string     `json:"employee_status"`
	Salary                   float64    `json:"salary"`
	ShowContract             int        `json:"show_contract"`
	EmployeeLetterCode       string     `json:"employee_letter_code"`
	BiodataConfirm           int        `json:"biodata_confirm"`
	BiodataConfirmDate       *time.Time `json:"biodata_confirm_date"`
	CurrentAddress           string     `json:"current_address"`
	BankAccountNumber        string     `json:"bank_account_number"`
	RoleID                   int        `json:"role_id"`
	Status                   int        `json:"status"`
	StatementRejected        string     `json:"statement_rejected"`
	IsDaily                  int        `json:"is_daily"`
	IsFlexibleAbsent         int        `json:"is_flexible_absent"`
	IsOvertime               int        `json:"is_overtime"`
	OvertimeLimit            int        `json:"overtime_limit"`
	FlexibleAbsentEndDate    string     `json:"flexible_absent_end_date"`
	DeviceToken              string     `json:"device_token"`
	EmailVerifiedAt          *time.Time `json:"email_verified_at"`
	StatusContractUpdate     int        `json:"status__contract_update"`
}
