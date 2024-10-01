package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	//"github.com/asajaroff/microblog/models"

	"github.com/jackc/pgx/v5"
)

var pool *sql.DB // Database connection pool

func init() {
	dsn, exists := os.LookupEnv("DSN")
	if !exists {
		dsn = "postgres://postgres:password@localhost:5432/users?sslmode=disable"
	}

	// Init database
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	defer db.Close(context.Background())
}

func main() {
	fmt.Println("Aca estamos")
}
