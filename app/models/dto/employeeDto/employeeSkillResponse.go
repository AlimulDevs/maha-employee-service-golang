package employeeDto

type EmployeeSkillResponse struct {
	ID         int64  `json:"id"`
	EmployeeID int64  `json:"employee_id"`
	Skill      string `json:"skill"`
}
