// The TestDeleteDepartment function tests the deletion of a department using Fiber framework in Go.
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

func TestDeleteDepartment(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Delete("/department/:id", DeleteDepartment)

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
	response, _, err := lib.DeleteTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")

	// test delete with non existing id
	response, _, err = lib.DeleteTest(app, uri, headers)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test invalid token
	response, _, err = lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 401, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
