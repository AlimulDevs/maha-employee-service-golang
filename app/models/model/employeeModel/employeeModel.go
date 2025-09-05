package employeeModel

import (
	"api/app/models/model/branchModel"
	"api/app/models/model/departmentModel"
	"api/app/models/model/jobTitleModel"
	"time"
)

type EmployeeModel struct {
	ID                       uint64                           `json:"id"`
	Nik                      *string                          `json:"nik"`
	Fullname                 string                           `json:"fullname"`
	Email                    string                           `json:"email"`
	JobTitleID               int                              `json:"job_title_id"`
	PhoneNumber              string                           `json:"phone_number"`
	Photo                    string                           `json:"photo"`
	PhotoEmbedding           string                           `json:"photo_embedding"`
	DepartmentID             int                              `json:"department_id"`
	BranchCode               string                           `json:"branch_code"`
	Password                 string                           `json:"password"`
	Signature                string                           `json:"signature"`
	IntegrityPactNum         string                           `json:"integrity_pact_num"`
	IntegrityPactCheck       int                              `json:"integrity_pact_check"`
	IntegrityPactCheckDate   *time.Time                       `json:"integrity_pact_check_date"`
	StatementLetterCheck     int                              `json:"statement_letter_check"`
	StatementLetterCheckDate *time.Time                       `json:"statement_letter_check_date"`
	ContractID               *int64                           `json:"contract_id"`
	OldContractID            *int64                           `json:"old_contract_id"`
	EmployeeStatus           string                           `json:"employee_status"`
	Salary                   float64                          `json:"salary"`
	ShowContract             int                              `json:"show_contract"`
	EmployeeLetterCode       *string                          `json:"employee_letter_code"`
	BiodataConfirm           int                              `json:"biodata_confirm"`
	BiodataConfirmDate       *time.Time                       `json:"biodata_confirm_date"`
	CurrentAddress           string                           `json:"current_address"`
	BankAccountNumber        string                           `json:"bank_account_number"`
	RoleID                   int                              `json:"role_id"`
	Status                   int                              `json:"status"`
	StatementRejected        *string                          `json:"statement_rejected"`
	IsDaily                  int                              `json:"is_daily"`
	IsFlexibleAbsent         int                              `json:"is_flexible_absent"`
	IsOvertime               int                              `json:"is_overtime"`
	OvertimeLimit            *int                             `json:"overtime_limit"`
	FlexibleAbsentEndDate    string                           `json:"flexible_absent_end_date"`
	DeviceToken              string                           `json:"device_token"`
	EmailVerifiedAt          *time.Time                       `json:"email_verified_at"`
	StatusContractUpdate     int                              `json:"status__contract_update"`
	JobTitle                 *jobTitleModel.JobTitleModel     `json:"job_title" gorm:"foreignKey:JobTitleID;references:ID"`
	Department               *departmentModel.DepartmentModel `json:"department" gorm:"foreignKey:DepartmentID;references:ID"`
	Branch                   *branchModel.BranchModel         `json:"branch" gorm:"foreignKey:BranchCode;references:BranchCode"`
	EmployeeContract         *EmployeeContractModel           `json:"contract" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeEducation        *EmployeeEducationModel          `json:"education" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeSkill            *[]EmployeeSkillModel            `json:"employee_skill" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeBank             *EmployeeBankModel               `json:"employee_bank" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeBiodata          *EmployeeBiodataModel            `json:"biodata" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeFamily           *EmployeeFamilyModel             `json:"family" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeWorkHour         *EmployeeWorkHourModel           `json:"work_hour" gorm:"foreignKey:EmployeeID;references:ID"`
	EmployeeDocument         *EmployeeDocumentModel           `json:"document" gorm:"foreignKey:EmployeeID;references:ID"`
	CreatedAt                *time.Time                       `json:"created_at"`
	UpdatedAt                *time.Time                       `json:"updated_at"`
	DeletedAt                *time.Time                       `json:"deleted_at"`
}

func (EmployeeModel) TableName() string {
	return "employees" // Custom table name
}
