package employeeModel

type EmployeeWorkHourModel struct {
	ID         int64  `json:"id"`
	EmployeeID int64  `json:"employee_id"`
	Sunday     string `json:"sunday"`
	Monday     string `json:"monday"`
	Tuesday    string `json:"tuesday"`
	Wednesday  string `json:"wednesday"`
	Thursday   string `json:"thursday"`
	Friday     string `json:"friday"`
	Saturday   string `json:"saturday"`
}

func (EmployeeWorkHourModel) TableName() string {
	return "employee_work_hours" // Custom table name
}
