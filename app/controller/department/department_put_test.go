package department

import (
	"api/app/config"
	"api/app/lib"
	"api/app/middleware"
	"api/app/model"
	"api/app/services"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/spf13/viper"
)

func TestPutDepartment(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Put("/department/:id", PutDepartment)

	initial := model.Department{
		DepartmentRequest: model.DepartmentRequest{
			DepartmentCode: "test 2",
			DepartmentName: "test 2",
			IsSub:          1,
			GmNum:          1,
		},
	}

	initial2 := model.Department{
		DepartmentRequest: model.DepartmentRequest{
			DepartmentCode: "test",
			DepartmentName: "test",
			IsSub:          1,
			GmNum:          1,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/department/" + fmt.Sprintf("%v", initial.ID)

	payload := `{
		 "department_code": "test 2",
            "department_name": "test",
            "is_sub": 0,
            "gm_num": 0
	}`

	headers := map[string]string{
		"Content-Type":                      "application/json",
		viper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
	}

	response, body, err := lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")

	// test invalid json body
	response, _, err = lib.PutTest(app, uri, headers, "invalid json format")
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, response.StatusCode, "getting response code")

	// test update with non existing id
	uri = "/department/1344"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/department/" + fmt.Sprintf("%v", initial2.ID)
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
