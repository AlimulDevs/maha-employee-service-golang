package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetEmployee(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []employeeModel.EmployeeModel

	var dto []employeeDto.EmployeeResponse

	err := db.Preload("JobTitle").Preload("Department").Preload("Branch").Find(&model).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}
	copier.Copy(&dto, &model)

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
