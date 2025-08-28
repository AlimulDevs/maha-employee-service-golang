package employeeDto

import (
	"api/app/lib"
	"api/app/models/model/employeeModel"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type EmployeeResponse struct {
	ID                       int64      `json:"id"`
	Nik                      *string    `json:"nik"`
	Fullname                 string     `json:"fullname"`
	Email                    string     `json:"email"`
	EmailVerifiedAt          time.Time  `json:"email_verified_at"`
	JobTitleID               int64      `json:"job_title_id"`
	PhoneNumber              string     `json:"phone_number"`
	DepartmentID             int64      `json:"department_id"`
	BranchCode               string     `json:"branch_code"`
	IntegrityPactNum         int        `json:"integrity_pact_num"`
	IntegrityPactCheck       int64      `json:"integrity_pact_check"`
	IntegrityPactCheckDate   *time.Time `json:"integrity_pact_check_date"`
	StatementLetterCheck     int64      `json:"statement_letter_check"`
	StatementLetterCheckDate *time.Time `json:"statement_letter_check_date"`
	ContractID               *int64     `json:"contract_id"`
	OldContractID            *string    `json:"old_contract_id"`
	EmployeeStatus           string     `json:"employee_status"`
	Salary                   *string    `json:"salary"`
	PhotoEmbedding           string     `json:"photo_embedding"`
	ShowContract             int64      `json:"show_contract"`
	EmployeeLetterCode       *string    `json:"employee_letter_code"`
	BiodataConfirm           int64      `json:"biodata_confirm"`
	BiodataConfirmDate       *time.Time `json:"biodata_confirm_date"`
	CurrentAddress           string     `json:"current_address"`
	BankAccountNumber        string     `json:"bank_account_number"`
	RoleID                   int        `json:"role_id"`
	Status                   int        `json:"status"`
	StatementRejected        *string    `json:"statement_rejected"`
	IsDaily                  int64      `json:"is_daily"`
	IsFlexibleAbsent         int64      `json:"is_flexible_absent"`
	FlexibleAbsentEndDate    string     `json:"flexible_absent_end_date"`
	IsOvertime               int64      `json:"is_overtime"`
	OvertimeLimit            *int       `json:"overtime_limit"`
	DeviceToken              *string    `json:"device_token"`
	CreatedAt                time.Time  `json:"created_at"`
	DeletedAt                *string    `json:"deleted_at"`
	StartWork                *string    `json:"start_work"`
	PhotoURL                 string     `json:"photo_url"`
	SignatureURL             *string    `json:"signature_url"`
	StatusLabel              string     `json:"status_label"`
	IsBpjs                   bool       `json:"is_bpjs"`
	IsBpjsContribution       bool       `json:"is_bpjs_contribution"`
	BpjsWages                string     `json:"bpjs_wages"`
	BpjsContributionWages    string     `json:"bpjs_contribution_wages"`
	JobTitle                 JobTitle   `json:"job_title"`
	Department               Department `json:"department"`
	Branch                   Branch     `json:"branch"`
}

func ToEmployeeGetAll(model employeeModel.EmployeeModel) EmployeeResponse {
	var photoURL string
	if model.Photo != "" { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", model.Photo)
		photoURL = url
	}
	var signatureURL string
	if model.Signature != "" { // ganti ke field `Photo` kalau ada
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", model.Signature)
		signatureURL = url
	}
	var data EmployeeResponse

	copier.Copy(&data, &model)
	data.PhotoURL = photoURL
	data.SignatureURL = &signatureURL
	salary := lib.FloatToStr(model.Salary)
	data.Salary = &salary
	return data

}

type Branch struct {
	ID               uint64     `json:"id" gorm:"primaryKey"`
	BranchCode       string     `json:"branch_code"`
	BranchLetterCode string     `json:"branch_letter_code"`
	BranchName       string     `json:"branch_name"`
	BranchLocation   string     `json:"branch_location"`
	BranchRadius     int32      `json:"branch_radius"`
	IsProject        int        `json:"is_project"`
	IsSub            int        `json:"is_sub"`
	BranchParentCode string     `json:"branch_parent_code"`
	Meal             int        `json:"meal"`
	IsActive         int        `json:"is_active"`
	IsDeleted        bool       `json:"is_deleted"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type Department struct {
	ID             *int64     `json:"id"`
	DepartmentCode *string    `json:"department_code"`
	DepartmentName *string    `json:"department_name"`
	IsSub          *int64     `json:"is_sub"`
	GmNum          *int64     `json:"gm_num"`
	IsDeleted      *int64     `json:"is_deleted"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type JobTitle struct {
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
