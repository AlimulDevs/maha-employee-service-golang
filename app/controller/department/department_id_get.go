package department

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetDepartmentID godoc
// @Summary Get an Department by ID (Department))
// @Description Retrieve detailed information of a specific Department using its unique ID. If the Department with the given ID is found, the Department's data is returned. </br>If no Department is found for the provided ID, a '404' error response is returned indicating the Department was not found.
// @Param id path string true "Department ID" - The unique identifier of the Department to retrieve.
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Department "Department data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Department matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /department/{id} [get]
// @Tags Department (Department)
func GetDepartmentID(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	id := lib.StrToUInt64(c.Params("id"))

	var data model.Department
	result := db.Model(&data).
		Where("id = ?", id).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
