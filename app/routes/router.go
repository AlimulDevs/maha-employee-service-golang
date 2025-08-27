package routes

import (
	"api/app/controller"
	"api/app/controller/attendance"
	"api/app/controller/branch"
	"api/app/controller/department"
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

	api.Post("department", middleware.TokenValidator(), department.PostDepartment)
	api.Get("department", middleware.TokenValidator(), department.GetDepartment)
	api.Get("department/:id", middleware.TokenValidator(), department.GetDepartmentID)
	api.Put("department/:id", middleware.TokenValidator(), department.PutDepartment)
	api.Delete("department/:id", middleware.TokenValidator(), department.DeleteDepartment)

	api.Post("branch", middleware.TokenValidator(), branch.PostBranch)
	api.Post("branch/get-all", middleware.TokenValidator(), branch.PostBranchGetAll)
	api.Get("branch", middleware.TokenValidator(), branch.GetBranchParent)
	api.Get("branch/children", middleware.TokenValidator(), branch.GetBranchChildren)
	api.Get("branch/children/parent/:parent_code", middleware.TokenValidator(), branch.GetBranchChildrenByParent)

	api.Get("branch/:branch_code/detail", middleware.TokenValidator(), branch.GetBranchDetailByCode)

	api.Get("attendance", middleware.TokenValidator(), attendance.GetAttendance)
	api.Get("attendance/:id", middleware.TokenValidator(), attendance.GetAttendanceID)

}
