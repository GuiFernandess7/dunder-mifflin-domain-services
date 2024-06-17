package main

import (
	"github.com/GuiFernandess7/db_with_sqlc/router"
	_ "github.com/lib/pq"
)

func main(){
	router.Initialize()
	//db := config.InitializeDB()

	//ctx := context.Background()
	//db_repository := database.New(db)

	//sex := "M"

	//employee, err := use_cases.GetEmployees(10, db_repository, ctx)

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//for i := range employee {
	//	fmt.Println(employee[i])
//}
}


