package attendance

import (
	"api/app/lib"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetAttendance godoc
// @Summary List of Attendance (Attendance)
// @Description Retrieve a paginated list of Attendance from the database. You can specify the page number, number of records per page, sorting order, and apply custom filters to refine the results. </br>By default, the first page is returned with 10 records per page. You can also specify which fields to include in the response for better performance.
// @Param page query int false "Page number starting from zero. Default is 0."
// @Param size query int false "Number of records per page. Default is 10."
// @Param sort query string false "Sort by a specific field. Prefix with a dash (`-`) for descending order, e.g., `-name`."
// @Param fields query string false "Comma-separated list of specific fields to include in the response."
// @Param filters query string false "Custom filters for querying data. See [filter format documentation](https://github.com/morkid/paginate#filter-format) for more details."
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Attendance} "Paginated list of Attendance"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Attendance matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /attendance [get]
// @Tags Attendance (Attendance)
type AttendanceFlat struct {
	AttendanceID     uint64  `json:"attendance_id"`
	EmployeeID       uint64  `json:"employee_id"`
	AttendanceDate   string  `json:"attendance_date"`
	EntrySchedule    *string `json:"entry_schedule"`
	ClockIn          *string `json:"clock_in"`
	ClockOut         *string `json:"clock_out"`
	EmployeeFullname string  `json:"employee_fullname"`
	EmployeeEmail    string  `json:"employee_email"`
}

type AttendanceWithEmployee struct {
	AttendanceID   uint64  `json:"id"`
	EmployeeID     uint64  `json:"employee_id"`
	AttendanceDate string  `json:"attendance_date"`
	EntrySchedule  *string `json:"entry_schedule"`
	ClockIn        *string `json:"clock_in"`
	ClockOut       *string `json:"clock_out"`
	EmployeeData   struct {
		EmployeeFullname string `json:"employee_fullname"`
		EmployeeEmail    string `json:"employee_email"`
	} `json:"employee_data"`
}

func GetAttendance(c *fiber.Ctx) error {
	db := services.DB_ATTENDANCE.WithContext(c.UserContext())
	var flatData []AttendanceFlat

	err := db.Raw(`
		SELECT 
			a.id AS attendance_id,
			a.employee_id,
			a.attendance_date,
			a.entry_schedule,
			a.clock_in,
			a.clock_out,
			e.fullname AS employee_fullname,
			e.email AS employee_email
		FROM hrisapps_attendance.attendances a
		JOIN hrisapps_employee.employees e ON a.employee_id = e.id
		LIMIT 1000
	`).Scan(&flatData).Error
	if err != nil {
		return err
	}

	data := make([]AttendanceWithEmployee, len(flatData))
	for i, f := range flatData {
		data[i].AttendanceID = f.AttendanceID
		data[i].EmployeeID = f.EmployeeID
		data[i].AttendanceDate = f.AttendanceDate
		data[i].EntrySchedule = f.EntrySchedule
		data[i].ClockIn = f.ClockIn
		data[i].ClockOut = f.ClockOut
		data[i].EmployeeData.EmployeeFullname = f.EmployeeFullname
		data[i].EmployeeData.EmployeeEmail = f.EmployeeEmail
	}

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    data,
	}

	return lib.OK(c, response)
}
