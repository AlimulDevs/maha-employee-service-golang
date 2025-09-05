package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetEmployeeDocumentById(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	id := c.Params("id")

	var model employeeModel.EmployeeDocumentModel

	var dto employeeDto.EmployeeDocumentResponse

	err := db.Where("employee_id = ?", id).
		First(&model).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}

	copier.Copy(&dto, employeeDto.ToEmployeeDocumentResponse(model))

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
