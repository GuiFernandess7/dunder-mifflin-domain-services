package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync" // Para sincronização segura de concorrência

	database "github.com/GuiFernandess7/db_with_sqlc/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func InitializeDB() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading the .env file:", err)
		}

		connStr := os.Getenv("CONN_STRING")
		secret := os.Getenv("SECRET")

		dbInstance, err = setDBConnection(connStr, secret)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})

	return dbInstance
}

func DBRepository() (*database.Queries, error) {
	db := InitializeDB()
	db_repository := database.New(db)
	return db_repository, nil
}

func setDBConnection(connStr, password string) (*sql.DB, error) {
	if password == "" {
		return nil, fmt.Errorf("password variable is not defined")
	}

	encodedPassword := url.QueryEscape(password)
	connStr = fmt.Sprintf(connStr, encodedPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}
