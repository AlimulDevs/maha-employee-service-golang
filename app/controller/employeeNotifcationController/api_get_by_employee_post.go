package employeeNotifcationController

import (
	"api/app/lib"
	"api/app/models/dto/employeeNotificationDto"
	"api/app/models/model/employeeNotificationModel"
	"fmt"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetByEmployeeId(c *fiber.Ctx) error {
	api := new(employeeNotificationDto.EmployeeNotificationRequest)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB.WithContext(c.UserContext())

	var model []employeeNotificationModel.EmployeeNotificationModel

	var dto []employeeNotificationDto.EmployeeNotificationResponse

	query := db.Model(&employeeNotificationModel.EmployeeNotificationModel{}).Where("employee_worker_id = ?", api.EmployeeWorkerID)

	if api.IsRead != nil && *api.IsRead != 2 {
		fmt.Println(*api.IsRead)
		query = query.Where("is_read = ?", *api.IsRead)
	}
	if api.StatusPerson != nil {
		query = query.Where("status_person = ?", *api.StatusPerson)
	}
	if api.StatusLabel != nil {
		query = query.Where("status_person = ?", *api.StatusPerson)
	}

	err := query.Find(&model).Error

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
