package use_cases

import (
	"context"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func GetEmployees(limit int32, db_repository *database.Queries, ctx context.Context)([]database.ListAllEmployeesRow, error){
	result, err := db_repository.ListAllEmployees(ctx)
	if err != nil {
		log.Fatalf("error getting all employees: %v", err)
		return nil, err
	}
	return result, nil
}

