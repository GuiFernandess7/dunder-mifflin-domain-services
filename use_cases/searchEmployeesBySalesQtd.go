package use_cases

import (
	"context"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func SearchEmployeesBySalesQtd(min int32, max int32, db_repository *database.Queries, ctx context.Context) ([]database.GetEmployeeBySalesQtdRow, error){
	params := database.GetEmployeeBySalesQtdParams{
		Min: min,
		Max: max,
	}
	result, err := db_repository.GetEmployeeBySalesQtd(ctx, params)
	if err != nil {
		log.Fatalf("error getting employees by sales qtd")
		return nil, err
	}
	return result, nil
}