package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetEmployee(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []employeeModel.EmployeeModel

	var dto []employeeDto.EmployeeResponse

	status := c.Query("status")     // contoh: ?status=1
	fullname := c.Query("fullname") // contoh: ?fullname=Jamal

	query := db.Model(&employeeModel.EmployeeModel{}).Preload("JobTitle").Preload("Department").Preload("Branch")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if fullname != "" {
		query = query.Where("fullname LIKE ?", "%"+fullname+"%")
	}

	err := query.Find(&model).Error
	if err != nil {
		return lib.ErrorNotFound(c)
	}

	for _, dt := range model {
		dto = append(dto, employeeDto.ToEmployeeGetAll(dt))
	}

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
