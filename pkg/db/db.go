package db

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO pool and cleanup of db/DB
var db *pgxpool.Pool

func InitDB() (*pgxpool.Pool, error) {
	dsn, exists := os.LookupEnv("DSN")
	if !exists {
		dsn = "postgres://postgres:password@localhost:5432/microblog?sslmode=disable"
	}

	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, err
	}
	config.MaxConns = 10                               // Set maximum number of
	config.MinConns = 2                                // Set minimum number of connections
	config.MaxConnIdleTime = 5 * time.Minute           // Idle connection timeout
	config.ConnConfig.ConnectTimeout = 3 * time.Second // Connection timeout

	// Create the connection pool with the customized config
	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDB() *pgxpool.Pool {
	return db
}
