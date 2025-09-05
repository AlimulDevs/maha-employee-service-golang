package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetEmployeeById(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	id := c.Params("id")

	var model employeeModel.EmployeeModel

	var dto employeeDto.EmployeeGetByIDResponse

	err := db.Where("id = ?", id).
		Preload("JobTitle").
		Preload("Branch").
		Preload("Department").
		Preload("EmployeeEducation").
		Preload("EmployeeSkill").
		Preload("EmployeeBank.Bank").
		Preload("EmployeeContract.ContractJobdesk").
		Preload("EmployeeBiodata").
		Preload("EmployeeFamily").
		Preload("EmployeeWorkHour").
		Preload("EmployeeDocument").
		First(&model).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}
	copier.Copy(&dto, employeeDto.ToEmployeeGetById(model))

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
