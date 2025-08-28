package bankcontroller

import (
	"api/app/lib"
	"api/app/models/dto/bankDto"
	"api/app/models/model/bankModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetBank(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []bankModel.BankModel

	var dto []bankDto.BankResponse

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
