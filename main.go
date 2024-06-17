package main

import (
	"github.com/GuiFernandess7/db_with_sqlc/router"
	_ "github.com/lib/pq"
)

func main(){
	router.Initialize()
	/* db := config.InitializeDB()

	ctx := context.Background()
	db_repository := database.New(db)

	employee, err := use_cases.SearchClientsByBranch("Scranton", db_repository, ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range employee {
		fmt.Println(employee[i])
	} */
}


