package employeeModel

type EmployeeFamilyModel struct {
	ID                   int64   `json:"id"`
	EmployeeID           int64   `json:"employee_id"`
	FatherName           string  `json:"father_name"`
	FatherStatus         int64   `json:"father_status"`
	FatherAge            int64   `json:"father_age"`
	FatherLastEducation  string  `json:"father_last_education"`
	FatherLastJobTitle   string  `json:"father_last_job_title"`
	FatherLastJobCompany string  `json:"father_last_job_company"`
	MotherName           string  `json:"mother_name"`
	MotherStatus         int64   `json:"mother_status"`
	MotherAge            int64   `json:"mother_age"`
	MotherLastEducation  string  `json:"mother_last_education"`
	MotherLastJobTitle   string  `json:"mother_last_job_title"`
	MotherLastJobCompany string  `json:"mother_last_job_company"`
	MaritalStatus        string  `json:"marital_status"`
	MaritalYear          *int64  `json:"marital_year"`
	CoupleName           *string `json:"couple_name"`
	CoupleAge            *int16  `json:"couple_age"`
	CoupleLastEducation  *string `json:"couple_last_education"`
	CoupleLastJobTitle   *string `json:"couple_last_job_title"`
	CoupleLastJobCompany *string `json:"couple_last_job_company"`
}

func (EmployeeFamilyModel) TableName() string {
	return "employee_families" // Custom table name
}
