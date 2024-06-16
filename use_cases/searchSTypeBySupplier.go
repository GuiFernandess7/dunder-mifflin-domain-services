package use_cases

import (
	"context"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func SearchTypeBySupplier(t string, db_repository *database.Queries, ctx context.Context) ([]database.GetSupplierByTypeRow, error){
	result, err := db_repository.GetSupplierByType(ctx, t)
	if err != nil {
		log.Fatalf("error getting supply type by supplier")
		return nil, err
	}
	return result, nil
}