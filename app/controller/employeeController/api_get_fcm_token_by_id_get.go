package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetEmployeeFcmTokenById(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	id := c.Params("id")

	var model employeeModel.EmployeeFcmTokenModel

	var dto employeeDto.EmployeeFcmTokenResponse

	err := db.Where("employee_id = ?", id).
		First(&model).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}

	copier.Copy(&dto, &model)
	dto.GlobalPassword = "E!!fu!0T--T4~h@7hQ"

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    dto,
	}

	return lib.OK(c, response)
}
