package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	database "github.com/GuiFernandess7/db_with_sqlc/db"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func ListEmployeesFilterHandler(ctx *gin.Context) {
    dbRepository, err := config.DBRepository()
    if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

    minSalaryStr := ctx.Query("min")
	maxSalaryStr := ctx.Query("max")
	sex := ctx.Query("sex")

	minSalary, err := parseSalary(ctx, minSalaryStr, "min")
	if err != nil {
		return
	}

	maxSalary, err := parseSalary(ctx, maxSalaryStr, "max")
	if err != nil {
		return
	}

	var employees []database.Employee

    if sex != "" {
        err := validateSexParam(sex)
        if err != nil {
            sendError(ctx, http.StatusBadRequest, err.Error())
            return
        }
    }

	if sex == "" {
		employees, err = use_cases.GetEmployeesBy(minSalary, maxSalary, nil, dbRepository, ctx)
	} else {
        //validateSexParam(sex, ctx)
		employees, err = use_cases.GetEmployeesBy(minSalary, maxSalary, &sex, dbRepository, ctx)
	}

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to fetch employees")
		return
	}

	responseData, err := getResponseData(employees)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to format response data")
        return
	}

	sendSuccess(ctx, "success", responseData)
}

func validateSexParam(sex string) error {
    if sex != "M" && sex != "F" {
        return errors.New("sex parameter has to be 'M' or 'F'")
    }
    return nil
}

func parseSalary(ctx *gin.Context, salaryStr string, salaryType string) (int32, error) {
	if salaryStr == "" {
		return 0, nil
	}

	salary64, err := strconv.ParseInt(salaryStr, 10, 32)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid "+salaryType+" salary")
		return 0, err
	}

	return int32(salary64), nil
}

func getResponseData(employees []database.Employee) ([]EmployeeResponse, error) {
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

	return responseData, nil
}