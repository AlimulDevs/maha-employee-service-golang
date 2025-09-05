package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetEmployeeSignatureById(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	id := c.Params("id")

	var model employeeModel.EmployeeModel

	var dto employeeDto.EmployeeSignatureResponse

	err := db.Where("id = ?", id).
		Select("id", "signature").
		First(&model).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}
	dto.EmployeeID = model.ID
	dto.SignatureURL = lib.GenerateEmployeeFileURL(&model.Signature)

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
