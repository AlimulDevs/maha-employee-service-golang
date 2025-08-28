package departmentController

import (
	"api/app/lib"
	"api/app/models/dto/departmentDto"
	"api/app/models/model/departmentModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetDepartment(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []departmentModel.DepartmentModel

	var dto []departmentDto.DepartmentResponse

	err := db.Preload("JobTitle").Find(&model).Error

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
