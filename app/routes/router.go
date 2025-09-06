package routes

import (
	"api/app/controller"
	bankcontroller "api/app/controller/bankController"
	"api/app/controller/branchController"
	"api/app/controller/departmentController"
	"api/app/controller/employeeController"
	"api/app/controller/employeeNotifcationController"
	"api/app/controller/jobTitleController"
	"api/app/controller/workHourController"

	"api/app/middleware"

	"api/app/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			lib.PrintStackTrace(e)
		},
	}))

	api := app.Group(viper.GetString("ENDPOINT"))

	api.Static("/swagger", "docs/swagger.json")
	api.Get("/", controller.GetAPIIndex)
	api.Get("/info.json", controller.GetAPIInfo)

	api.Post("branch/get-all", middleware.VerifyTokenJwt(), branchController.PostBranchGetAll)
	api.Get("branch", middleware.VerifyTokenJwt(), branchController.GetBranchParent)
	api.Get("branch/children", middleware.VerifyTokenJwt(), branchController.GetBranchChildren)
	api.Get("branch/children/parent/:parent_code", middleware.VerifyTokenJwt(), branchController.GetBranchChildrenByParent)
	api.Get("branch/:branch_code/detail", middleware.VerifyTokenJwt(), branchController.GetBranchDetailByCode)

	api.Get("bank", middleware.VerifyTokenJwt(), bankcontroller.GetBank)
	api.Get("bank/:id", middleware.VerifyTokenJwt(), bankcontroller.GetBankById)

	api.Get("department", middleware.VerifyTokenJwt(), departmentController.GetDepartment)
	api.Get("department/:code", middleware.VerifyTokenJwt(), departmentController.GetDepartmentByCode)

	api.Get("job-title", middleware.VerifyTokenJwt(), jobTitleController.GetJobTitle)
	api.Get("job-title/:id", middleware.VerifyTokenJwt(), jobTitleController.GetJobTitleById)
	api.Get("job-title/department/:department_id", middleware.VerifyTokenJwt(), jobTitleController.GetJobTitleByDepartment)

	api.Get("work-hour", middleware.VerifyTokenJwt(), workHourController.GetWorkHour)
	api.Get("work-hour/:code", middleware.VerifyTokenJwt(), workHourController.GetWorkHourByCode)

	api.Get("employee", middleware.VerifyTokenJwt(), employeeController.GetEmployee)
	api.Get("employee/:id", middleware.VerifyTokenJwt(), employeeController.GetEmployeeById)
	api.Get("employee/employee-signature/:id", middleware.VerifyTokenJwt(), employeeController.GetEmployeeSignatureById)
	api.Get("employee/employee-fcm-token/:id", middleware.VerifyTokenJwt(), employeeController.GetEmployeeFcmTokenById)
	api.Get("employee/employee-document/:id", middleware.VerifyTokenJwt(), employeeController.GetEmployeeDocumentById)

	api.Post("employee/employee-notification/get-by-employee-id", middleware.VerifyTokenJwt(), employeeNotifcationController.GetByEmployeeId)

}
