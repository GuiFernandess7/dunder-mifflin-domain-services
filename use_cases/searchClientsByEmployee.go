package use_cases

import (
	"context"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func SearchClientsByEmployee(first_name string, db_repository *database.Queries, ctx context.Context) ([]database.GetClientsByEmployeeRow, error){
	result, err := db_repository.GetClientsByEmployee(ctx, first_name)
	if err != nil {
		log.Fatalf("error getting employees by sales qtd")
		return nil, err
	}
	return result, nil
}