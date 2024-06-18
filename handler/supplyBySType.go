package handler

import (
	"net/http"

	"github.com/GuiFernandess7/db_with_sqlc/config"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/gin-gonic/gin"
)

func GetSupplierBySupplyType(ctx *gin.Context){
	dbRepository, err := config.DBRepository()
	var responseData []SupplierResponse
	if err != nil {
        sendError(ctx, http.StatusBadRequest, "Failed to connect to database")
        return
    }

	supply_type := ctx.Query("type")

	if supply_type == "" {
		sendError(ctx, http.StatusInternalServerError, "supply type must be provided.")
		return
	}
	sups, err := use_cases.SearchTypeBySupplier(supply_type, dbRepository, ctx)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to fetch suppliers")
		return
	}

	for _, s := range sups {
		response := SupplierResponse{
			SupplierName: s.SupplierName,
			BranchName: s.BranchName.String,
			SupplyType: s.SupplyType.String,
		}
		responseData = append(responseData, response)
	}
	sendSuccess(ctx, "get suppliers and branches by supply type", responseData)

}