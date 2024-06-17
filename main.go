package main

import (
	"github.com/GuiFernandess7/db_with_sqlc/router"
	_ "github.com/lib/pq"
)

func main(){
	router.Initialize()
	/* db, err := config.InitializeDB()
	if err != nil {
		return
	}

	ctx := context.Background()
	db_repository := database.New(db)

	employee, err := use_cases.SearchEmployeesBySalesQtd(1000, 20000, db_repository, ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range employee {
		fmt.Println(employee[i])
	} */
}


