package department

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteDepartment godoc
// @Summary Delete Department by id (Department)
// @Description Delete an Department record based on its unique ID. If the Department is found, it will be removed from the database. </br>If the Department does not exist, an error response will be returned. If there is a conflict preventing deletion, an appropriate error message will be provided.
// @Param id path string true "Department ID - The unique identifier of the Department to be deleted"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response "Successful deletion of Department"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Department found with the provided ID."
// @Failure 409 {object} lib.Response "Conflict: The Department is in use or cannot be deleted due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /department/{id} [delete]
// @Tags Department (Department)
func DeleteDepartment(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var data model.Department
	result := db.Model(&data).Where("id = ?", c.Params("id")).Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Unscoped().Delete(&data)

	return lib.OK(c)
}
