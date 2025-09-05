package employeeModel

type EmployeeContractModel struct {
	ID                    int64                  `json:"id"`
	EmployeeID            int64                  `json:"employee_id"`
	LetterNumber          string                 `json:"letter_number"`
	LetterID              *uint16                `json:"letter_id"`
	JobTitleID            int64                  `json:"job_title_id"`
	DepartmentID          int64                  `json:"department_id"`
	BranchCode            string                 `json:"branch_code"`
	ContractStatus        string                 `json:"contract_status"`
	Salary                string                 `json:"salary"`
	Project               *string                `json:"project"`
	ContractLengthNum     int64                  `json:"contract_length_num"`
	ContractLengthTime    string                 `json:"contract_length_time"`
	StartContract         string                 `json:"start_contract"`
	EndContract           string                 `json:"end_contract"`
	JobdeskContent        *string                `json:"jobdesk_content"`
	CheckContract         int64                  `json:"check_contract"`
	CheckContractDatetime *string                `json:"check_contract_datetime"`
	ApproverID            *uint64                `json:"approver_id"`
	ApproverJobTitle      *uint64                `json:"approver_job_title"`
	ConfirmContract       int64                  `json:"confirm_contract"`
	ConfirmContractDate   *string                `json:"confirm_contract_date"`
	ContractFile          *string                `json:"contract_file"`
	Status                int64                  `json:"status"`
	StatusParaf           bool                   `json:"status_paraf"`
	SalaryText            string                 `json:"salary_text"`
	ContractJobdesk       []ContractJobdeskModel `json:"jobdesks" gorm:"foreignKey:ContractID;references:ID"`
}

type ContractJobdeskModel struct {
	ID         uint64 `json:"id"`
	ContractID int64  `json:"contract_id"`
	Jobdesk    string `json:"jobdesk"`
}

func (EmployeeContractModel) TableName() string {
	return "employee_contracts" // Custom table name
}
func (ContractJobdeskModel) TableName() string {
	return "contract_jobdesks" // Custom table name
}
