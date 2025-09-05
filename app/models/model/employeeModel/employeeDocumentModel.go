package employeeModel

type EmployeeDocumentModel struct {
	ID                  int64   `json:"id"`
	EmployeeID          int64   `json:"employee_id"`
	Photo               *string `json:"photo"`
	Ktp                 *string `json:"ktp"`
	Kk                  *string `json:"kk"`
	Certificate         *string `json:"certificate"`
	GradeTranscript     *string `json:"grade_transcript"`
	CertificateSkillURL *string `json:"certificate_skill"`
	BankAccount         *string `json:"bank_account"`
	Npwp                *string `json:"npwp"`
	BpjsKtn             *string `json:"bpjs_ktn"`
	BpjsKes             *string `json:"bpjs_kes"`
}

func (EmployeeDocumentModel) TableName() string {
	return "employee_documents" // Custom table name
}
