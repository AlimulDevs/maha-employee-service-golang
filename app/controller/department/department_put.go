package department

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutDepartment godoc
// @Summary Update Department by id (Department)
// @Description Update an existing Department using its unique ID. Provide the new data for the Department, and if the update is successful, the updated Department's data will be returned. </br>If the Department does not exist or there is a conflict during the update, appropriate error responses will be provided.
// @Param id path string true "Department ID"
// @Param data body model.DepartmentRequest true "Department data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Department "Updated Department data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Department matched the provided ID."
// @Failure 409 {object} lib.Response "Conflict: The Department is in use or cannot be updated due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /department/{id} [put]
// @Tags Department (Department)
func PutDepartment(c *fiber.Ctx) error {
	api := new(model.DepartmentRequest)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())
	id := lib.StrToUInt64(c.Params("id"))

	var data model.Department
	result := db.Model(&data).
		Where(db.Where(model.Department{

			ID: id,
		})).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.OK(c, data)
}
