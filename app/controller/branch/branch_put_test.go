package branch

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

func TestPutBranch(t *testing.T) {
	db := services.DBConnectTest()
	lib.LoadEnvironment(config.Environment)

	app := fiber.New()
	app.Use(middleware.TokenValidator())

	app.Put("/branch/:id", PutBranch)

	initial := model.Branch{
		BranchCode:       "branch_code",
		BranchName:       "branch_name",
		BranchLetterCode: "branch_letter_code",
	}

	initial2 := model.Branch{
		BranchCode:       "branch_code 2",
		BranchName:       "branch_name 2",
		BranchLetterCode: "branch_letter_code 2",
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/branch/" + fmt.Sprintf("%v", initial.ID)

	payload := `{
		"branch_code": "OJK",
		"branch_name": "OJK",
		"branch_letter_code": "OJK"
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
	uri = "/branch/100"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/branch/" + fmt.Sprintf("%v", initial2.ID)
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
