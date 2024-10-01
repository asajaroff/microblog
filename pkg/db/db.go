package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

// TODO pool and cleanup of db/DB
var db *pgx.Conn
var DB *pgx.Conn

func InitDB() (*pgx.Conn, error) {
	dsn, exists := os.LookupEnv("DSN")
	if !exists {
		dsn = "postgres://postgres:password@localhost:5432/users_system?sslmode=disable"
	}

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close(context.Background())
	DB = db
	return db, nil
}

func GetDB() *pgx.Conn {
	return db
}
