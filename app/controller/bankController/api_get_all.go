package bankcontroller

import (
	"api/app/lib"
	bankdto "api/app/models/dto/bankDto"
	bankmodel "api/app/models/model/bankModel"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetBank(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []bankmodel.BankModel

	var dto []bankdto.BankResponse

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
