package bankDto

type BankRequest struct {
	ID      int64  `json:"id"`
	Jenis   string `json:"jenis"`
	Kode    string `json:"kode"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telp    string `json:"telp"`
	Website string `json:"website"`
	Fax     string `json:"fax"`
}
