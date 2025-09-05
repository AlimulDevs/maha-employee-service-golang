package employeeDto

type EmployeeSignatureResponse struct {
	EmployeeID   uint64  `json:"employee_id"`
	SignatureURL *string `json:"signature_url"`
}
