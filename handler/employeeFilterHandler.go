package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func ListEmployeesFilterHandler(ctx *gin.Context) {
    dbRepository, err := config.DBRepository()
    if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

    var responseData []EmployeeResponse
    //sex := ctx.Query("sex")
    minSalaryStr := ctx.Query("min")
    maxSalaryStr := ctx.Query("max")

    var minSalary, maxSalary int32
    if minSalaryStr != "" {
        minSalary64, err := strconv.ParseInt(minSalaryStr, 10, 32)
        if err != nil {
            sendError(ctx, http.StatusBadRequest, "Invalid min salary")
            return
        }
        minSalary = int32(minSalary64)
    }
    if maxSalaryStr != "" {
        maxSalary64, err := strconv.ParseInt(maxSalaryStr, 10, 32)
        if err != nil {
            sendError(ctx, http.StatusBadRequest, "Invalid max salary")
            return
        }
        maxSalary = int32(maxSalary64)
    }

    employees, err := use_cases.GetEmployeesBy(minSalary, maxSalary, nil, dbRepository, ctx)
	fmt.Println(employees)
    if err != nil {
        sendError(ctx, http.StatusBadRequest, err.Error())
        return
    }

	if len(employees) != 0 {
		for _, emp := range employees {
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
	} else {
		responseData = []EmployeeResponse{}
	}
    sendSuccess(ctx, "success", responseData)
}
