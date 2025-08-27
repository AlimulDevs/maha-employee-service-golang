package migrations

import "gorm.io/gorm"

var ()

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{}
}

// InitialSeeds seeds the initial data if it doesn't already exist
func InitialSeeds(tx *gorm.DB) error {
	return nil
}
