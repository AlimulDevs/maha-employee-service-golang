package employeeDto

type EmployeeBankResponse struct {
	ID            int64  `json:"id"`
	EmployeeID    int64  `json:"employee_id"`
	BankID        int64  `json:"bank_id"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}
