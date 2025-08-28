package workHourDto

import "time"

type WorkHourRequest struct {
	ID             int64      `json:"id"`
	WorkHourCode   string     `json:"work_hour_code"`
	WorkHourName   string     `json:"work_hour_name"`
	StartEntryHour string     `json:"start_entry_hour"`
	EntryHour      string     `json:"entry_hour"`
	EndEntryHour   string     `json:"end_entry_hour"`
	HomeHour       string     `json:"home_hour"`
	DifferentDay   int64      `json:"different_day"`
	IsDeleted      bool       `json:"is_deleted"`
	DeletedAt      *time.Time `json:"deleted_at"`
}
