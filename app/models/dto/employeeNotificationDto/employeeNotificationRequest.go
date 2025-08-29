package employeeNotificationDto

import "time"

type EmployeeNotificationRequest struct {
	ID                 int64     `json:"id"`
	EmployeeWorkerID   int64     `json:"employee_worker_id"`
	StatusPerson       *int      `json:"status_person"`
	StatusNotification int       `json:"status_notification"`
	StatusLabel        *int      `json:"status_label"`
	Type               string    `json:"type"`
	DateTime           time.Time `json:"date_time"`
	Content            string    `json:"content"`
	IsRead             *int      `json:"is_read"`
	DataID             int64     `json:"data_id"`
}
