package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func GetEmployeeBySalesQtd(ctx *gin.Context){
	dbRepository, err := config.DBRepository()
	if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

	min := ctx.Query("min")
	max := ctx.Query("max")

	minParsed, err := parseSales(ctx, min, "min")
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error parsing sales value")
		return
	}

	maxParsed, err := parseSales(ctx, max, "max")
	if err != nil {
		log.Fatalf("error parsing sales value: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error parsing sales value")
		return
	}

	result, err := use_cases.SearchEmployeesBySalesQtd(minParsed, maxParsed, dbRepository, ctx)
	if err != nil {
		log.Fatalf("error getting employees by sales qtd: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error getting employees by sales qtd")
		return
	}
	var responseData []EmployeeNameResponse

	for _, emp := range result {
		response := EmployeeNameResponse{
			FirstName: emp.FirstName.String,
			LastName: emp.LastName.String,
		}
		responseData = append(responseData, response)
	}

	sendSuccess(ctx, "getting employees by sales", responseData)
}

func parseSales(ctx *gin.Context, salaryStr string, salaryType string) (int32, error) {
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