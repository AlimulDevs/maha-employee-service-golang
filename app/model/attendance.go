package model

// Attendance model (entity utama)
type Attendance struct {
	ID                     uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID             uint64  `json:"employee_id" gorm:"column:employee_id;index"`
	AttendanceDate         string  `json:"attendance_date" gorm:"column:attendance_date;not null"`
	AttendanceDateOut      *string `json:"attendance_date_out,omitempty" gorm:"column:attendance_date_out"`
	EntrySchedule          *string `json:"entry_schedule,omitempty" gorm:"column:entry_schedule"`
	HomeSchedule           *string `json:"home_schedule,omitempty" gorm:"column:home_schedule"`
	ClockIn                *string `json:"clock_in,omitempty" gorm:"column:clock_in"`
	ClockOut               *string `json:"clock_out,omitempty" gorm:"column:clock_out"`
	BreakStart             *string `json:"break_start,omitempty" gorm:"column:break_start"`
	BreakFinish            *string `json:"break_finish,omitempty" gorm:"column:break_finish"`
	OvertimeStart          *string `json:"overtime_start,omitempty" gorm:"column:overtime_start"`
	OvertimeFinish         *string `json:"overtime_finish,omitempty" gorm:"column:overtime_finish"`
	PhotoIn                *string `json:"photo_in,omitempty" gorm:"column:photo_in;size:255"`
	PhotoOut               *string `json:"photo_out,omitempty" gorm:"column:photo_out;size:255"`
	LocationIn             *string `json:"location_in,omitempty" gorm:"column:location_in;size:255"`
	LocationOut            *string `json:"location_out,omitempty" gorm:"column:location_out;size:255"`
	OvertimeStartPhoto     *string `json:"overtime_start_photo,omitempty" gorm:"column:overtime_start_photo;size:255"`
	OvertimeFinishPhoto    *string `json:"overtime_finish_photo,omitempty" gorm:"column:overtime_finish_photo;size:255"`
	OvertimeStartLocation  *string `json:"overtime_start_location,omitempty" gorm:"column:overtime_start_location;size:255"`
	OvertimeFinishLocation *string `json:"overtime_finish_location,omitempty" gorm:"column:overtime_finish_location;size:255"`
	WorkHourCode           *string `json:"work_hour_code,omitempty" gorm:"column:work_hour_code;size:255"`
	ClockInType            int     `json:"clock_in_type" gorm:"column:clock_in_type;default:1"`
	ClockOutType           int     `json:"clock_out_type" gorm:"column:clock_out_type;default:1"`
	IsLate                 bool    `json:"is_late" gorm:"column:is_late;default:0"`
	EarlyOut               bool    `json:"early_out" gorm:"column:early_out;default:0"`
	ClockInStatus          int     `json:"clock_in_status" gorm:"column:clock_in_status;default:1"`
	ClockOutStatus         int     `json:"clock_out_status" gorm:"column:clock_out_status;default:1"`
	StatementInRejected    *string `json:"statement_in_rejected,omitempty" gorm:"column:statement_in_rejected;type:text"`
	StatementOutRejected   *string `json:"statement_out_rejected,omitempty" gorm:"column:statement_out_rejected;type:text"`
	ClockInZone            bool    `json:"clock_in_zone" gorm:"column:clock_in_zone;default:1"`
	ClockOutZone           bool    `json:"clock_out_zone" gorm:"column:clock_out_zone;default:1"`
	MealNum                int     `json:"meal_num" gorm:"column:meal_num;default:0"`
	BranchAttendance       *string `json:"branch_attendance,omitempty" gorm:"column:branch_attendance;size:255"`
	AttendanceStatus       *string `json:"attendance_status,omitempty" gorm:"column:attendance_status;size:255"`
	CreateStatus           bool    `json:"create_status" gorm:"column:create_status;default:0"`
	CreatedAt              string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt              string  `json:"updated_at" gorm:"column:updated_at"`
}

// TableName overrides the default table name behavior
func (Attendance) TableName() string {
	return "attendances" // sesuai nama tabel
}

// AttendanceRequest (untuk input)
type AttendanceRequest struct {
	EmployeeID       uint64  `json:"employee_id"`
	AttendanceDate   string  `json:"attendance_date"`
	ClockIn          *string `json:"clock_in,omitempty"`
	ClockOut         *string `json:"clock_out,omitempty"`
	PhotoIn          *string `json:"photo_in,omitempty"`
	PhotoOut         *string `json:"photo_out,omitempty"`
	LocationIn       *string `json:"location_in,omitempty"`
	LocationOut      *string `json:"location_out,omitempty"`
	AttendanceStatus *string `json:"attendance_status,omitempty"`
}

// AttendanceResponse (untuk output API)
type AttendanceResponse struct {
	ID               uint64  `json:"id"`
	EmployeeID       uint64  `json:"employee_id"`
	AttendanceDate   string  `json:"attendance_date"`
	ClockIn          *string `json:"clock_in,omitempty"`
	ClockOut         *string `json:"clock_out,omitempty"`
	IsLate           bool    `json:"is_late"`
	EarlyOut         bool    `json:"early_out"`
	AttendanceStatus *string `json:"attendance_status,omitempty"`
}
