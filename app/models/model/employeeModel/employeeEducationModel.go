package employeeModel

type EmployeeEducationModel struct {
	ID                  uint64  `json:"id"`
	EmployeeID          uint64  `json:"employee_id"`
	LastEducation       string  `json:"last_education"`
	PrimarySchool       *string `json:"primary_school"`
	PSStartYear         *string `json:"ps_start_year"`
	PSEndYear           *string `json:"ps_end_year"`
	PSCertificate       *string `json:"ps_certificate"`
	PSGpa               *string `json:"ps_gpa"`
	JuniorHighSchool    *string `json:"junior_high_school"`
	JhsStartYear        *string `json:"jhs_start_year"`
	JhsEndYear          *string `json:"jhs_end_year"`
	JhsCertificate      *string `json:"jhs_certificate"`
	JhsGpa              *string `json:"jhs_gpa"`
	SeniorHighSchool    string  `json:"senior_high_school"`
	ShsStartYear        string  `json:"shs_start_year"`
	ShsEndYear          string  `json:"shs_end_year"`
	ShsCertificate      *string `json:"shs_certificate"`
	ShsGpa              *string `json:"shs_gpa"`
	BachelorUniversity  string  `json:"bachelor_university"`
	BachelorMajor       string  `json:"bachelor_major"`
	BachelorStartYear   string  `json:"bachelor_start_year"`
	BachelorEndYear     string  `json:"bachelor_end_year"`
	BachelorCertificate *string `json:"bachelor_certificate"`
	BachelorGpa         string  `json:"bachelor_gpa"`
	BachelorDegree      string  `json:"bachelor_degree"`
	MasterUniversity    *string `json:"master_university"`
	MasterMajor         *string `json:"master_major"`
	MasterStartYear     *string `json:"master_start_year"`
	MasterEndYear       *string `json:"master_end_year"`
	MasterCertificate   *string `json:"master_certificate"`
	MasterGpa           *string `json:"master_gpa"`
	MasterDegree        *string `json:"master_degree"`
	DoctoralUniversity  *string `json:"doctoral_university"`
	DoctoralMajor       *string `json:"doctoral_major"`
	DoctoralStartYear   *string `json:"doctoral_start_year"`
	DoctoralEndYear     *string `json:"doctoral_end_year"`
	DoctoralCertificate *string `json:"doctoral_certificate"`
	DoctoralGpa         *string `json:"doctoral_gpa"`
	DoctoralDegree      *string `json:"doctoral_degree"`
	LastEducationMajor  string  `json:"last_education_major"`
}

func (EmployeeEducationModel) TableName() string {
	return "employee_educations" // Custom table name
}
