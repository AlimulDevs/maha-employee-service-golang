package employeeController

import (
	"api/app/lib"
	"api/app/models/dto/employeeDto"
	"api/app/models/model/employeeModel"

	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetEmployee(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())

	var model []employeeModel.EmployeeModel

	var dto []employeeDto.EmployeeResponse

	status := c.Query("status")     // contoh: ?status=1
	fullname := c.Query("fullname") // contoh: ?fullname=Jamal
	project := c.Query("project")
	branch_code := c.Query("branch_code")
	my_id := c.Query("my_id")
	role_id := c.Query("role_id")
	department_code := c.Query("department_code")
	job_title := c.Query("job_title")

	query := db.Model(&employeeModel.EmployeeModel{}).Preload("JobTitle").
		Preload("Department", func(db *gorm.DB) *gorm.DB {
			departmenQuery := db

			// Tambahkan filter 'is_project' jika parameter 'project' tidak kosong.
			if department_code != "" {

				departmenQuery = departmenQuery.Where("department_code = ?", department_code)
			}

			return departmenQuery
		}).
		Preload("Branch", func(db *gorm.DB) *gorm.DB {
			branchQuery := db

			// Tambahkan filter 'is_project' jika parameter 'project' tidak kosong.
			if project != "" {

				branchQuery = branchQuery.Where("is_project = ?", project)
			}

			// Tambahkan filter 'branch_code' jika parameter 'branch_code' tidak kosong.
			if branch_code != "" {
				branchQuery = branchQuery.Where("branch_code = ?", branch_code)
			}

			// Mengembalikan query yang sudah digabungkan.
			return branchQuery
		})

	if my_id != "" {
		query = query.Where("id != ?", my_id)
	}

	if job_title != "" {
		query = query.Where("job_title_id = ?", job_title)

	}

	if role_id != "" {
		query = query.Where("role_id = ?", role_id)
	}

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
		if dt.Branch == nil || dt.JobTitle == nil || dt.Department == nil {
			continue
		}

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
