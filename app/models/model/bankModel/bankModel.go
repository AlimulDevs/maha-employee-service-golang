package bankModel

import "time"

type BankModel struct {
	ID        int64      `json:"id"`
	Jenis     string     `json:"jenis"`
	Kode      string     `json:"kode"`
	Nama      string     `json:"nama"`
	Alamat    string     `json:"alamat"`
	Telp      string     `json:"telp"`
	Website   string     `json:"website"`
	Fax       string     `json:"fax"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (BankModel) TableName() string {
	return "banks" // Custom table name
}
