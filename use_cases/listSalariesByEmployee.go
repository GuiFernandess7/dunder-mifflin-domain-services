package use_cases

import (
	"context"
	"errors"
	"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)


func ListSalariesByEmployee(mode string, db_repository *database.Queries, ctx context.Context) ([]database.Employee, error){
	if mode == "ASC" {
		result, err := db_repository.ListSalariesASC(ctx)
		if err != nil {
			log.Fatalf("failed to list salaries in asc order: %v", err)
			return nil, err
		}
		return result, nil

	} else if mode == "DESC" {
		result, err := db_repository.ListSalariesDESC(ctx)
		if err != nil {
			log.Fatalf("failed to list salaries in desc order: %v", err)
			return nil, err
		}
		return result, nil

	} else {
		return nil, errors.New("order invalid")
	}
}