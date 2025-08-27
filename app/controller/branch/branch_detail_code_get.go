package branch

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetBranchID godoc
// @Summary Get an Branch by ID (Branch))
// @Description Retrieve detailed information of a specific Branch using its unique ID. If the Branch with the given ID is found, the Branch's data is returned. </br>If no Branch is found for the provided ID, a '404' error response is returned indicating the Branch was not found.
// @Param id path string true "Branch ID" - The unique identifier of the Branch to retrieve.
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Branch "Branch data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Branch matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /branch/{id} [get]
// @Tags Branch (Branch)
func GetBranchDetailByCode(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	branch_code := c.Params("branch_code")

	var data model.Branch
	result := db.Model(&data).
		Where("branch_code = ?", branch_code).
		Take(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}
	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    data,
	}

	return lib.OK(c, response)
}
