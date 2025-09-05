package lib

import "fmt"

// GenUUIDString func
func GenerateEmployeeFileURL(path *string) *string {
	if path != nil && *path != "" {
		url := fmt.Sprintf("https://employee-service.mahasejahtera.com/public/storage/%s", *path)
		return &url
	}
	return nil
}
