package employeeModel

import "api/app/models/model/bankModel"

type EmployeeBankModel struct {
	ID            int64                `json:"id"`
	EmployeeID    int64                `json:"employee_id"`
	BankID        int64                `json:"bank_id"`
	AccountNumber string               `json:"account_number"`
	AccountName   string               `json:"account_name"`
	Bank          *bankModel.BankModel `json:"bank" gorm:"foreignKey:BankID;references:ID"`
}

func (EmployeeBankModel) TableName() string {
	return "employee_banks" // Custom table name
}
