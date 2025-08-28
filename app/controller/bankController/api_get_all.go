package bankcontroller

import (
	"api/app/lib"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetBank(c *fiber.Ctx) error {
	services.DB.WithContext(c.UserContext())

	// var model []bankmodel.BankModel

	// var dto []bankdto.BankResponse

	// err := db.Find(&model).Error

	// if err != nil {
	// 	return lib.ErrorNotFound(c)
	// }
	// copier.Copy(&dto, &model)

	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    nil,
	}

	return lib.OK(c, response)
}
