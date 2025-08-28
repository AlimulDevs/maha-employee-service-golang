package workHourController

import (
	"api/app/lib"
	"api/app/models/dto/workHourDto"
	"api/app/models/model/workHourModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetWorkHour(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []workHourModel.WorkHourModel

	var dto []workHourDto.WorkHourResponse

	err := db.Find(&model).Error

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
