package branch

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutBranch godoc
// @Summary Update Branch by id (Branch)
// @Description Update an existing Branch using its unique ID. Provide the new data for the Branch, and if the update is successful, the updated Branch's data will be returned. </br>If the Branch does not exist or there is a conflict during the update, appropriate error responses will be provided.
// @Param id path string true "Branch ID"
// @Param data body model.BranchRequest true "Branch data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Branch "Updated Branch data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Branch matched the provided ID."
// @Failure 409 {object} lib.Response "Conflict: The Branch is in use or cannot be updated due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /branch/{id} [put]
// @Tags Branch (Branch)
func PutBranch(c *fiber.Ctx) error {
	api := new(model.BranchRequest)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())
	id := lib.StrToUInt64(c.Params("id"))

	var data model.Branch
	result := db.Model(&data).
		Where(db.Where(model.Branch{

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
