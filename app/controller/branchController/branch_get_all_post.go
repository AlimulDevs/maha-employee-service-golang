package branchController

import (
	"api/app/lib"
	"api/app/model"
	"api/app/models/dto/branchDto"
	"api/app/models/model/branchModel"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func PostBranchGetAll(c *fiber.Ctx) error {
	api := new(branchDto.BranchRequest)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	var data []branchModel.BranchModel

	err := db.Model(&model.Branch{}).Where("is_active = ?", api.IsActive).Find(&data).Error

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
