package handler

import (
	"net/http"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func GetClientByEmployeeHandler(ctx *gin.Context){
	dbRepository, err := config.DBRepository()
	if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

	employee_name := ctx.Query("name")

	if employee_name == "" {
		sendError(ctx, http.StatusBadRequest, "branch name must be passed")
		return
	}

	clients, err := use_cases.SearchClientsByEmployee(employee_name, dbRepository, ctx)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error getting branches")
	}

	var responseData []ClientNameResponse

	for _, c := range clients {
		response := ClientNameResponse{
			ClientID: int(c.ClientID),
			ClientName: c.ClientName.String,
		}
		responseData = append(responseData, response)
	}
	sendSuccess(ctx, "search employee by name", responseData)
}