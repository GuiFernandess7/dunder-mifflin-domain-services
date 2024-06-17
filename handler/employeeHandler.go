package handler

import (
	"net/http"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func ListEmployeesHandler(ctx *gin.Context){
	db_repository, err := config.DBRepository()
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "bad request")
	}
	employees, err := use_cases.GetEmployees(10, db_repository, ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "bad request 2")
	}
	var responseData []EmployeeResponse

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

    sendSuccess(ctx, "success", responseData)
}