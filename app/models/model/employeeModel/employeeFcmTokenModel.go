package employeeModel

import "time"

type EmployeeFcmTokenModel struct {
	ID            int64     `json:"id"`
	EmployeeID    int64     `json:"employee_id"`
	StatusAccount int64     `json:"status_account"`
	Token         string    `json:"token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (EmployeeFcmTokenModel) TableName() string {
	return "employee_fcm_tokens" // Custom table name
}
