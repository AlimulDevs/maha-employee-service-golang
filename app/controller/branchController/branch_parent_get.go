package branchController

import (
	"api/app/lib"
	"api/app/models/model/branchModel"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetBranchParent(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	var data []branchModel.BranchModel
	err := db.Where("branch_parent_code IS NULL").Find(&data).Error

	if err != nil {
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
