package employeeModel

type EmployeeBiodataModel struct {
	ID                   int64   `json:"id"`
	EmployeeID           int64   `json:"employee_id"`
	Fullname             string  `json:"fullname"`
	Nickname             string  `json:"nickname"`
	Nik                  string  `json:"nik"`
	IdentityProvince     int64   `json:"identity_province"`
	IdentityRegency      int64   `json:"identity_regency"`
	IdentityDistrict     int64   `json:"identity_district"`
	IdentityVillage      int64   `json:"identity_village"`
	IdentityPostalCode   int64   `json:"identity_postal_code"`
	IdentityAddress      string  `json:"identity_address"`
	CurrentProvince      int64   `json:"current_province"`
	CurrentRegency       int64   `json:"current_regency"`
	CurrentDistrict      int64   `json:"current_district"`
	CurrentVillage       int64   `json:"current_village"`
	CurrentPostalCode    int64   `json:"current_postal_code"`
	CurrentAddress       string  `json:"current_address"`
	ResidenceStatus      string  `json:"residence_status"`
	PhoneNumber          string  `json:"phone_number"`
	EmergencyPhoneNumber string  `json:"emergency_phone_number"`
	StartWork            string  `json:"start_work"`
	Gender               string  `json:"gender"`
	BirthPlace           string  `json:"birth_place"`
	BirthDate            string  `json:"birth_date"`
	Religion             string  `json:"religion"`
	BloodType            *string `json:"blood_type"`
	Weight               *int16  `json:"weight"`
	Height               *int16  `json:"height"`
	CurrentFullAddress   string  `json:"current_full_address"`
	IdentityFullAddress  string  `json:"identity_full_address"`
	Age                  string  `json:"age"`
}

func (EmployeeBiodataModel) TableName() string {
	return "employee_biodatas" // Custom table name
}
