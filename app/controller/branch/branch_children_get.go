package branch

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetBranch godoc
// @Summary List of Branch (Branch)
// @Description Retrieve a paginated list of Branch from the database. You can specify the page number, number of records per page, sorting order, and apply custom filters to refine the results. </br>By default, the first page is returned with 10 records per page. You can also specify which fields to include in the response for better performance.
// @Param page query int false "Page number starting from zero. Default is 0."
// @Param size query int false "Number of records per page. Default is 10."
// @Param sort query string false "Sort by a specific field. Prefix with a dash (`-`) for descending order, e.g., `-name`."
// @Param fields query string false "Comma-separated list of specific fields to include in the response."
// @Param filters query string false "Custom filters for querying data. See [filter format documentation](https://github.com/morkid/paginate#filter-format) for more details."
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Branch} "Paginated list of Branch"
// @Failure 400 {object} lib.Response "Bad Request: Invalid parameters provided."
// @Failure 404 {object} lib.Response "Not Found: No Branch matched the query."
// @Failure 500 {object} lib.Response "Internal Server Error: Unexpected error occurred."
// @Failure default {object} lib.Response "Unexpected error response."
// @Security TokenKey
// @Router /branch [get]
// @Tags Branch (Branch)
func GetBranchChildren(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	var data []model.Branch
	err := db.Where("branch_parent_code IS NOT NULL").Find(&data).Error

	if err != nil {
		return lib.ErrorNotFound(c)
	}
	response := lib.BaseResponse{
		Status:  "success",
		Code:    200,
		Message: "OK",
		Data:    data,
	}

	return lib.OK(c, response)
}
