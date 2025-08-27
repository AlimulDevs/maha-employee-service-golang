const fs = require('fs');
const path = require('path');

const folderName = process.argv[2]; // contoh: facultyTes
if (!folderName) {
  console.error('Harap berikan nama folder. Contoh: node generateHandler.js facultyTes');
  process.exit(1);
}

const tags = "";

function toSnakeCase(str) {
  return str.replace(/([a-z])([A-Z])/g, '$1_$2').toLowerCase();
}

function toPascalCase(str) {
  return str.replace(/(^|_)(\w)/g, (_, __, c) => c.toUpperCase());
}


function toSpacedName(str) {
  return str
    .replace(/([a-z0-9])([A-Z])/g, '$1 $2')  // tambahkan spasi sebelum huruf kapital
    .replace(/^./, s => s.toUpperCase());   // kapitalisasi huruf pertama
}



const entityName = toSpacedName(folderName);
const entitySnake = toSnakeCase(folderName);      // faculty_tes
const entityPascal = toPascalCase(entitySnake);   // FacultyTes
const folderPath = path.join(__dirname, 'app', 'controller', folderName);

// Buat folder jika belum ada
fs.mkdirSync(folderPath, { recursive: true });

const postFileName = `${entitySnake}_post.go`;
const getFileName = `${entitySnake}_get.go`;
const getIdFileName = `${entitySnake}_id_get.go`;
const putFileName = `${entitySnake}_put.go`;
const deleteFileName = `${entitySnake}_delete.go`;

const postTesFileName = `${entitySnake}_post_test.go`;
const getTesFileName = `${entitySnake}_get_test.go`;
const getIdTesFileName = `${entitySnake}_id_get_test.go`;
const putTesFileName = `${entitySnake}_put_test.go`;
const deleteTesFileName = `${entitySnake}_delete_test.go`;

const postContent = `package ${folderName}

import (
\t"api/app/lib"
\t"api/app/model"
\t"api/app/services"
\t"fmt"

\t"github.com/gofiber/fiber/v2"
)

// Post${entityPascal} godoc
// @Summary Create new ${entityPascal} (${entityPascal})
// @Description Create a new ${entityPascal} by providing the required data. If the ${entityPascal} is successfully created, a \`201\` response with the created ${entityPascal} data will be returned. In case of errors, appropriate error messages will be returned, such as invalid input data or conflicts.
// @Param data body model.${entityPascal}Request true "${entityPascal} data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.${entityPascal} "Created ${entityPascal} data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 409 {object} lib.Response "Conflict: The ${entityPascal} is in use or cannot be created due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /${entitySnake.replace(/_/g, '-')} [post]
// @Tags ${entityName} (${entityName})
func Post${entityPascal}(c *fiber.Ctx) error {
\tapi := new(model.${entityPascal}Request)
\tif err := lib.BodyParser(c, api); nil != err {
\t\treturn lib.ErrorBadRequest(c, err)
\t}

\tdb := services.DB.WithContext(c.UserContext())

\tvar data model.${entityPascal}
\tvar totalData int64

\tdb.Model(&model.${entityPascal}{}).Count(&totalData)
\tfmt.Println(totalData)
\ttotalData++
\tdata.Sort = lib.Int64ptr((totalData))

\tlib.Merge(api, &data)
\tdata.CreatorID = lib.GetXUserID(c)

\tif err := db.Create(&data).Error; nil != err {
\t\treturn lib.ErrorConflict(c, err.Error())
\t}

\treturn lib.Created(c, data)
}
`;

const getContent = `package ${folderName}

import (
\t"api/app/lib"
\t"api/app/model"
\t"api/app/services"

\t"github.com/gofiber/fiber/v2"
)

// Get${entityPascal} godoc
// @Summary List of ${entityPascal} (${entityPascal})
// @Description Retrieve a paginated list of ${entityPascal} from the database. You can specify the page number, number of records per page, sorting order, and apply custom filters to refine the results. </br>By default, the first page is returned with 10 records per page. You can also specify which fields to include in the response for better performance.
// @Param page query int false "Page number starting from zero. Default is 0."
// @Param size query int false "Number of records per page. Default is 10."
// @Param sort query string false "Sort by a specific field. Prefix with a dash (\`-\`) for descending order, e.g., \`-name\`."
// @Param fields query string false "Comma-separated list of specific fields to include in the response."
// @Param filters query string false "Custom filters for querying data. See [filter format documentation](https://github.com/morkid/paginate#filter-format) for more details."
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.${entityPascal}} "Paginated list of ${entityPascal}"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No ${entityPascal} matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /${entitySnake.replace(/_/g, '-')} [get]
// @Tags ${entityName} (${entityName})
func Get${entityPascal}(c *fiber.Ctx) error {
\tdb := services.DB.WithContext(c.UserContext())
\tpg := services.PG

\tmod := db.Model(&model.${entityPascal}{})

\tpage := pg.With(mod).Request(c.Request()).Response(&[]model.${entityPascal}{})

\treturn lib.OK(c, page)
}
`;
const getIdContent = `package ${folderName}

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get${entityPascal}ID godoc
// @Summary Get an ${entityPascal} by ID (${entityPascal}))
// @Description Retrieve detailed information of a specific ${entityPascal} using its unique ID. If the ${entityPascal} with the given ID is found, the ${entityPascal}'s data is returned. </br>If no ${entityPascal} is found for the provided ID, a '404' error response is returned indicating the ${entityPascal} was not found.
// @Param id path string true "${entityPascal} ID" - The unique identifier of the ${entityPascal} to retrieve.
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.${entityPascal} "${entityPascal} data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No ${entityPascal} matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /${entitySnake.replace(/_/g, '-')}/{id} [get]
// @Tags ${entityName} (${entityName})
func Get${entityPascal}ID(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	id, _ := uuid.Parse(c.Params("id"))

	var data model.${entityPascal}
	result := db.Model(&data).
		Where(db.Where(model.${entityPascal}{
			Base: model.Base{
				ID: &id,
			},
		})).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
`;


const putContent = `
package ${folderName}

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Put${entityPascal} godoc
// @Summary Update ${entityPascal} by id (${entityPascal})
// @Description Update an existing ${entityPascal} using its unique ID. Provide the new data for the ${entityPascal}, and if the update is successful, the updated ${entityPascal}'s data will be returned. </br>If the ${entityPascal} does not exist or there is a conflict during the update, appropriate error responses will be provided.
// @Param id path string true "${entityPascal} ID"
// @Param data body model.${entityPascal}Request true "${entityPascal} data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.${entityPascal} "Updated ${entityPascal} data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No ${entityPascal} matched the provided ID."
// @Failure 409 {object} lib.Response "Conflict: The ${entityPascal} is in use or cannot be updated due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /${entitySnake.replace(/_/g, '-')}/{id} [put]
// @Tags ${entityName} (${entityName})
func Put${entityPascal}(c *fiber.Ctx) error {
	api := new(model.${entityPascal}Request)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())
	id, _ := uuid.Parse(c.Params("id"))

	var data model.${entityPascal}
	result := db.Model(&data).
		Where(db.Where(model.${entityPascal}{
			Base: model.Base{
				ID: &id,
			},
		})).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)
	data.ModifierID = lib.GetXUserID(c)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.OK(c, data)
}
`;


const deleteContent = `package ${folderName}
import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// Delete${entityPascal} godoc
// @Summary Delete ${entityPascal} by id (${entityPascal})
// @Description Delete an ${entityPascal} record based on its unique ID. If the ${entityPascal} is found, it will be removed from the database. </br>If the ${entityPascal} does not exist, an error response will be returned. If there is a conflict preventing deletion, an appropriate error message will be provided.
// @Param id path string true "${entityPascal} ID - The unique identifier of the ${entityPascal} to be deleted"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response "Successful deletion of ${entityPascal}"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No ${entityPascal} found with the provided ID."
// @Failure 409 {object} lib.Response "Conflict: The ${entityPascal} is in use or cannot be deleted due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /${entitySnake.replace(/_/g, '-')}/{id} [delete]
// @Tags ${entityName} (${entityName})
func Delete${entityPascal}(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var data model.${entityPascal}
	result := db.Model(&data).Where("id = ?", c.Params("id")).Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Unscoped().Delete(&data)

	return lib.OK(c)
}

`;
const postPath = path.join(folderPath, postFileName);
const getPath = path.join(folderPath, getFileName);
const getIdPath = path.join(folderPath, getIdFileName);
const putPath = path.join(folderPath, putFileName);
const deletePath = path.join(folderPath, deleteFileName);

if (!fs.existsSync(postPath)) {
  fs.writeFileSync(postPath, postContent);
  console.log(`✅ File '${postFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${postFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(getPath)) {
  fs.writeFileSync(getPath, getContent);
  console.log(`✅ File '${getFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${getFileName}' sudah ada, tidak diubah.`);
}
if (!fs.existsSync(getIdPath)) {
  fs.writeFileSync(getIdPath, getIdContent);
  console.log(`✅ File '${getIdFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${getIdFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(putPath)) {
  fs.writeFileSync(putPath, putContent);
  console.log(`✅ File '${putFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${putFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(deletePath)) {
  fs.writeFileSync(deletePath, deleteContent);
  console.log(`✅ File '${deleteFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${deleteFileName}' sudah ada, tidak diubah.`);
}


// create testing file



const postTesContent = `package ${folderName}

import (
    \t"api/app/config"
    \t"api/app/lib"
    \t"api/app/middleware"
    \t"api/app/services"
    \t"testing"
    
    \t"github.com/gofiber/fiber/v2"
    \t"github.com/gofiber/fiber/v2/utils"
    \t"github.com/spf13/viper"
    )
    
    func TestPost${entityPascal}(t *testing.T) {
    \tdb := services.DBConnectTest()
    \tlib.LoadEnvironment(config.Environment)
    
    \tapp := fiber.New()
    \tapp.Use(middleware.TokenValidator())
    
    \tapp.Post("/${entitySnake.replace(/_/g, '-')}", Post${entityPascal})
    
    \turi := "/${entitySnake.replace(/_/g, '-')}"
    
    \tpayload := \`{
    \t\t"name": "Contoh"
    \t}\`
    
    \theaders := map[string]string{
    \t\t"Content-Type":                      "application/json",
    \t\tviper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
    \t}
    
    \tresponse, body, err := lib.PostTest(app, uri, headers, payload)
    \tutils.AssertEqual(t, nil, err, "sending request")
    \tutils.AssertEqual(t, 201, response.StatusCode, "getting response code")
    \tutils.AssertEqual(t, false, nil == body, "validate response body")
    
    \t// test invalid json format
    \tresponse, _, err = lib.PostTest(app, uri, headers, "invalid json format")
    \tutils.AssertEqual(t, nil, err, "sending request")
    \tutils.AssertEqual(t, 400, response.StatusCode, "getting response code")
    
    \t// test duplicate data
    \tresponse, _, err = lib.PostTest(app, uri, headers, payload)
    \tutils.AssertEqual(t, nil, err, "sending request")
    \tutils.AssertEqual(t, 409, response.StatusCode, "getting response code")
    
    \tsqlDB, _ := db.DB()
    \tsqlDB.Close()
    }
`;


const getTestContent = `package ${folderName}

import (
\t"api/app/config"
\t"api/app/lib"
\t"api/app/middleware"
\t"api/app/model"
\t"api/app/services"
\t"testing"

\t"github.com/gofiber/fiber/v2"
\t"github.com/gofiber/fiber/v2/utils"
\t"github.com/spf13/viper"
)

func TestGet${entityPascal}(t *testing.T) {
\tdb := services.DBConnectTest()
\tlib.LoadEnvironment(config.Environment)

\tapp := fiber.New()
\tapp.Use(middleware.TokenValidator())

\tapp.Get("/${entitySnake.replace(/_/g, '-')}", Get${entityPascal})

\tinitial := model.${entityPascal}{

\t}

\tdb.Create(&initial)

\theaders := map[string]string{
\t\tviper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
\t}

\turi := "/${entitySnake.replace(/_/g, '-')}?page=0&size=1"
\tresponse, body, err := lib.GetTest(app, uri, headers)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 200, response.StatusCode, "getting response code")
\tutils.AssertEqual(t, false, nil == body, "validate response body")
\tutils.AssertEqual(t, float64(1), body["total"], "getting response body")

\t// test invalid token
\tresponse, _, err = lib.GetTest(app, uri, nil)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 401, response.StatusCode, "getting response code")

\tsqlDB, _ := db.DB()
\tsqlDB.Close()
}
`;


const getIdTesContent = `package ${folderName}

import (
\t"api/app/config"
\t"api/app/lib"
\t"api/app/middleware"
\t"api/app/model"
\t"api/app/services"
\t"testing"

\t"github.com/gofiber/fiber/v2"
\t"github.com/gofiber/fiber/v2/utils"
\t"github.com/spf13/viper"
)

func TestGet${entityPascal}ID(t *testing.T) {
\tdb := services.DBConnectTest()
\tlib.LoadEnvironment(config.Environment)

\tapp := fiber.New()
\tapp.Use(middleware.TokenValidator())

\tapp.Get("/${entitySnake.replace(/_/g, '-')}/:id", Get${entityPascal}ID)

\tinitial := model.${entityPascal}{
\t\t${entityPascal}Request: model.${entityPascal}Request{
\t\t\tName: lib.Strptr("${entityName} Sample"),
\t\t},
\t}

\tdb.Create(&initial)

\theaders := map[string]string{
\t\tviper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
\t}

\turi := "/${entitySnake.replace(/_/g, '-')}/" + initial.ID.String()
\tresponse, body, err := lib.GetTest(app, uri, headers)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 200, response.StatusCode, "getting response code")
\tutils.AssertEqual(t, false, nil == body, "validate response body")
\tutils.AssertEqual(t, initial.ID.String(), body["id"], "getting response body")

\t// test get non existing id
\turi = "/${entitySnake.replace(/_/g, '-')}/non-existing-id"
\tresponse, _, err = lib.GetTest(app, uri, headers)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 404, response.StatusCode, "getting response code")

\t// test invalid token
\tresponse, _, err = lib.GetTest(app, uri, nil)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 401, response.StatusCode, "getting response code")

\tsqlDB, _ := db.DB()
\tsqlDB.Close()
}
`;




const putTesContent = `package ${folderName}

import (
\t"api/app/config"
\t"api/app/lib"
\t"api/app/middleware"
\t"api/app/model"
\t"api/app/services"
\t"testing"
\t"fmt"

\t"github.com/gofiber/fiber/v2"
\t"github.com/gofiber/fiber/v2/utils"
\t"github.com/spf13/viper"
)

func TestPut${entityPascal}(t *testing.T) {
\tdb := services.DBConnectTest()
\tlib.LoadEnvironment(config.Environment)

\tapp := fiber.New()
\tapp.Use(middleware.TokenValidator())

\tapp.Put("/${entitySnake.replace(/_/g, '-')}/:id", Put${entityPascal})

\tinitial := model.${entityPascal}{
\t\t${entityPascal}Request: model.${entityPascal}Request{
\t\t\tName:          lib.Strptr("${entityPascal}) Essay 1")
\t\t},
\t}

\tinitial2 := model.${entityPascal}{
\t\t${entityPascal}Request: model.${entityPascal}Request{
\t\t\tName:          lib.Strptr("${entityPascal}) Essay 2")
\t\t},
\t}

\tdb.Create(&initial)
\tdb.Create(&initial2)

\turi := "/${entitySnake.replace(/_/g, '-')}/" + initial.ID.String()

\tpayload := \`{
\t\t"name": "${entityPascal}) Essay 1"
\t}\`

\theaders := map[string]string{
\t\t"Content-Type":                      "application/json",
\t\tviper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
\t}

\tresponse, body, err := lib.PutTest(app, uri, headers, payload)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 200, response.StatusCode, "getting response code")
\tutils.AssertEqual(t, false, nil == body, "validate response body")

\t// test invalid json body
\tresponse, _, err = lib.PutTest(app, uri, headers, "invalid json format")
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 400, response.StatusCode, "getting response code")

\t// test update with non existing id
\turi = "/${entitySnake.replace(/_/g, '-')}/non-existing-id"
\tresponse, _, err = lib.PutTest(app, uri, headers, payload)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 404, response.StatusCode, "getting response code")

\t// test duplicate data
\turi = "/${entitySnake.replace(/_/g, '-')}/" + initial2.ID.String()
\tresponse, _, err = lib.PutTest(app, uri, headers, payload)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 409, response.StatusCode, "getting response code")

\tsqlDB, _ := db.DB()
\tsqlDB.Close()
}
`;

const deleteTestContent = `package ${folderName}

import (
\t"api/app/config"
\t"api/app/lib"
\t"api/app/middleware"
\t"api/app/model"
\t"api/app/services"
\t"testing"
\t"fmt"

\t"github.com/gofiber/fiber/v2"
\t"github.com/gofiber/fiber/v2/utils"
\t"github.com/spf13/viper"
)

func TestDelete${entityPascal}(t *testing.T) {
\tdb := services.DBConnectTest()
\tlib.LoadEnvironment(config.Environment)

\tapp := fiber.New()
\tapp.Use(middleware.TokenValidator())

\tapp.Delete("/${entitySnake.replace(/_/g, '-')}/:id", Delete${entityPascal})

\tinitial := model.${entityPascal}{

\t}

\tdb.Create(&initial)

\theaders := map[string]string{
\t\tviper.GetString("HEADER_TOKEN_KEY"): viper.GetString("VALUE_TOKEN_KEY"),
\t}

\turi := "/${entitySnake.replace(/_/g, '-')}/" + fmt.Sprintf("%v", initial.ID)
\tresponse, _, err := lib.DeleteTest(app, uri, headers)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 200, response.StatusCode, "getting response code")

\t// test delete with non existing id
\tresponse, _, err = lib.DeleteTest(app, uri, headers)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 404, response.StatusCode, "getting response code")

\t// test invalid token
\tresponse, _, err = lib.DeleteTest(app, uri, nil)
\tutils.AssertEqual(t, nil, err, "sending request")
\tutils.AssertEqual(t, 401, response.StatusCode, "getting response code")

\tsqlDB, _ := db.DB()
\tsqlDB.Close()
}
`;


const postTesPath = path.join(folderPath, postTesFileName);
const getTesPath = path.join(folderPath, getTesFileName);
const getIdTesPath = path.join(folderPath, getIdTesFileName);
const putTesPath = path.join(folderPath, putTesFileName);
const deleteTesPath = path.join(folderPath, deleteTesFileName);

if (!fs.existsSync(postTesPath)) {
  fs.writeFileSync(postTesPath, postTesContent);
  console.log(`✅ File '${postTesFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${postTesFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(getTesPath)) {
  fs.writeFileSync(getTesPath, getTestContent);
  console.log(`✅ File '${getTesFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${getTesFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(getIdTesPath)) {
  fs.writeFileSync(getIdTesPath, getIdTesContent);
  console.log(`✅ File '${getIdTesFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${getIdTesFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(putTesPath)) {
  fs.writeFileSync(putTesPath, putTesContent);
  console.log(`✅ File '${putTesFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${putTesFileName}' sudah ada, tidak diubah.`);
}

if (!fs.existsSync(deleteTesPath)) {
  fs.writeFileSync(deleteTesPath, deleteTestContent);
  console.log(`✅ File '${deleteTesFileName}' berhasil dibuat.`);
} else {
  console.log(`⚠️  File '${deleteTesFileName}' sudah ada, tidak diubah.`);
}
