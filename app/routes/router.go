package routes

import (
	"api/app/controller"
	bankcontroller "api/app/controller/bankController"
	"api/app/controller/branchController"

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

	api.Post("branch/get-all", middleware.TokenValidator(), branchController.PostBranchGetAll)
	api.Get("branch", middleware.TokenValidator(), branchController.GetBranchParent)
	api.Get("branch/children", middleware.TokenValidator(), branchController.GetBranchChildren)
	api.Get("branch/children/parent/:parent_code", middleware.TokenValidator(), branchController.GetBranchChildrenByParent)
	api.Get("branch/:branch_code/detail", middleware.TokenValidator(), branchController.GetBranchDetailByCode)

	api.Get("bank", middleware.TokenValidator(), bankcontroller.GetBank)
	api.Get("bank/:id", middleware.TokenValidator(), bankcontroller.GetBankById)

}
