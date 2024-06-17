package handler

import (
	"net/http"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func GetClientByBranchHandler(ctx *gin.Context){
	dbRepository, err := config.DBRepository()
	if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

	branch := ctx.Query("branch")

	if branch == "" {
		sendError(ctx, http.StatusBadRequest, "branch name must be passed")
		return
	}

	data, err := use_cases.SearchClientsByBranch(branch, dbRepository, ctx)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error getting branches")
	}

	var responseData []ClientResponse

	for _, client := range data {
		response := ClientResponse{
			ClientID: int(client.ClientID),
			ClientName: client.ClientName.String,
			Branch: client.BranchName.String,
		}
		responseData = append(responseData, response)
	}
	sendSuccess(ctx, "get client by branch", responseData)
}