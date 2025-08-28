package jobTitleController

import (
	"api/app/lib"
	"api/app/models/dto/jobTitleDto"
	"api/app/models/model/jobTitleModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetJobTitle(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []jobTitleModel.JobTitleModel

	var dto []jobTitleDto.JobTitleResponse

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
