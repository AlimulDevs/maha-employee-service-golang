package branch

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// PostBranch godoc
// @Summary Create new Branch (Branch)
// @Description Create a new Branch by providing the required data. If the Branch is successfully created, a `201` response with the created Branch data will be returned. In case of errors, appropriate error messages will be returned, such as invalid input data or conflicts.
// @Param data body model.BranchRequest true "Branch data"
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Branch "Created Branch data"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 409 {object} lib.Response "Conflict: The Branch is in use or cannot be created due to a conflict."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /branch [post]
// @Tags Branch (Branch)
func PostBranch(c *fiber.Ctx) error {
	api := new(model.BranchRequest)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	var data model.Branch
	var totalData int64

	db.Model(&model.Branch{}).Count(&totalData)
	fmt.Println(totalData)
	totalData++

	lib.Merge(api, &data)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err.Error())
	}

	return lib.Created(c, data)
}
