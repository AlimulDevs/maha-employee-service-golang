package attendance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetAttendanceID godoc
// @Summary Get an Attendance by ID (Attendance))
// @Description Retrieve detailed information of a specific Attendance using its unique ID. If the Attendance with the given ID is found, the Attendance's data is returned. </br>If no Attendance is found for the provided ID, a '404' error response is returned indicating the Attendance was not found.
// @Param id path string true "Attendance ID" - The unique identifier of the Attendance to retrieve.
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Attendance "Attendance data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Attendance matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /attendance/{id} [get]
// @Tags Attendance (Attendance)
func GetAttendanceID(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	id := lib.StrToUInt64(c.Params("id"))

	var data model.Attendance
	result := db.Model(&data).
		Where(db.Where(model.Attendance{
			ID: id,
		})).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
