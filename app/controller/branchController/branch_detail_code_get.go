package branchController

import (
	"api/app/lib"
	"api/app/models/model/branchModel"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetBranchDetailByCode(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	branch_code := c.Params("branch_code")

	var data branchModel.BranchModel
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
