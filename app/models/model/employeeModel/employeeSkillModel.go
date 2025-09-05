package employeeModel

type EmployeeSkillModel struct {
	ID         int64  `json:"id"`
	EmployeeID int64  `json:"employee_id"`
	Skill      string `json:"skill"`
}

func (EmployeeSkillModel) TableName() string {
	return "employee_skills" // Custom table name
}
