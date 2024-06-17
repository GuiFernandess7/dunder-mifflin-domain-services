package handler

import (
	"net/http"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func SearchEmployeeHandler(ctx *gin.Context){
	dbRepository, err := config.DBRepository()
	var responseData []EmployeeResponse
    if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

	name := ctx.Query("name")
	emp, err := use_cases.SearchEmployeeByName(name, dbRepository, ctx)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to fetch employees")
		return
	}

	for _, emp := range emp {
		response := EmployeeResponse{
			EmpID:    int(emp.EmpID),
			FirstName: emp.FirstName.String,
			LastName: emp.LastName.String,
			BirthDay: &emp.BirthDay.Time,
			Sex:      emp.Sex.String,
			Salary:   int(emp.Salary.Int32),
		}
		responseData = append(responseData, response)
	}
	sendSuccess(ctx, "search employee by name", responseData)
}