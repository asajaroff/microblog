package main

import (
	"context"
	//"database/sql"
	//"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/asajaroff/microblog/pkg/models"
	//"github.com/asajaroff/microblog/pkg/api"
	"github.com/asajaroff/microblog/pkg/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var pool *pgxpool.Pool

func init() {
	// Initialize database
	var err error
	pool, err = db.InitDB()
	if err != nil {
		log.Fatalf("FATAL: Cannot connect to database %v", err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "microblog v0.1")
	})

	e.GET("/ping", func(c echo.Context) error {
		if err := pool.Ping(context.Background()); err != nil {
			return c.String(http.StatusServiceUnavailable, "Cannot ping database")
		}
		log.Println("Pong!")
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/query", func(c echo.Context) error {
		var row string
		err := pool.QueryRow(context.Background(), "SELECT name from users LIMIT 1").Scan(&row)
		if err != nil {
			log.Printf("ERR: %s\n", err)
		}
		log.Println(row)
		return c.String(http.StatusOK, "query made!")
	})

	e.POST("/login", func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		var encrypted_password string
		err := pool.QueryRow(context.Background(),
			"SELECT encrypted_password FROM users WHERE email = $1", email).Scan(&encrypted_password)
		if err != nil {
			log.Println("ERR: %s\n", err)
			os.Exit(1)
			return c.String(http.StatusUnauthorized, "Cannot validate credentials")
		}

		if password == encrypted_password {
			log.Printf("INFO: User logged in: %s\n", email)
			return c.String(http.StatusOK, "Welcome back!")
		}
		return c.String(http.StatusUnauthorized, "Cannot validate credentials")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
