package use_cases

import (
	"context"
	"database/sql"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func SearchEmployeeByName(first_name string, db_repository *database.Queries, ctx context.Context) ([]database.Employee, error){
	selectedName := sql.NullString{
        String: first_name,
        Valid:  true,
    }
	result, err := db_repository.FindEmployeeByName(ctx, selectedName)
	if err != nil {
		log.Fatalf("failed to find employee: %v", err)
		return nil, err
	}
	return result, err
}