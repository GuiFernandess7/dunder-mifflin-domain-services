package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	database "github.com/GuiFernandess7/db_with_sqlc/db"
	"github.com/GuiFernandess7/db_with_sqlc/use_cases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	db, err := setDBConnection(os.Getenv("CONN_STRING"), os.Getenv("SECRET"))

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	defer db.Close()

	ctx := context.Background()
	db_repository := database.New(db)

	employee, err := use_cases.SearchEmployeesBySalesQtd(1000, 20000, db_repository, ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range employee {
		fmt.Println(employee[i])
	}
}

func setDBConnection(conn_str string, password string) (*sql.DB, error){
	if password == "" {
		log.Fatal("Password variable is not defined")
	}

	encodedPassword := url.QueryEscape(password)
	connStr := fmt.Sprintf(conn_str, encodedPassword)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

