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

func TestGetDepartmentID(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Get("/department/:id", GetDepartmentID)

	initial := model.Department{
		DepartmentRequest: model.DepartmentRequest{
			DepartmentCode: "test",
			DepartmentName: "test",
			IsSub:          1,
			GmNum:          1,
		},
	}

	db.Create(&initial)

	headers := map[string]string{
		viper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
	}

	uri := "/department/" + fmt.Sprintf("%v", initial.ID)
	response, body, err := lib.GetTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, fmt.Sprintf("%v", initial.ID), fmt.Sprintf("%v", body["id"]), "getting response body")

	// test get non existing id
	uri = "/department/non-existing-id"
	response, _, err = lib.GetTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test invalid token
	response, _, err = lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 401, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
