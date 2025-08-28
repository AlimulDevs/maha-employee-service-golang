package config

// Environment variables
var Environment = map[string]interface{}{
	"app_id":          "6A6EF734-F557-4F7E-9C2A-32CD28E43420",
	"app_version":     "v1.0.0",
	"app_name":        "Master Service",
	"app_description": "",
	"port":            9091,
	"endpoint":        "/api/v1/employee",
	"environment":     "development",
	"db_host":         "mysql-maha",
	"db_port":         3306,
	"db_user":         "root",
	"db_pass":         "alimul130502",
	"db_name":         "hrisapps_employee",
	"db_table_prefix": "",
	"redis_host":      "redis",
	"redis_port":      6379,
	"redis_pass":      "",
	"redis_index":     0,
	"prefork":         false,

	"header_token_key": "x-Token",
	"value_token_key":  "secret123",
}
