package use_cases

import (
	"context"
	//"log"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
)

func SearchClientsByBranch(branch string, db_repository *database.Queries, ctx context.Context) ([]database.GetClientbyBranchRow, error){
	result, _ := db_repository.GetClientbyBranch(ctx, branch)
	println(result)
	return result, nil
}

