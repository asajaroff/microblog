package main

import (
	"context"
	//"database/sql"
	//"fmt"
	"log"
	"net/http"

	//"github.com/asajaroff/microblog/pkg/models"
	"github.com/asajaroff/microblog/pkg/api"
	"github.com/asajaroff/microblog/pkg/db"

	//"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func init() {
	// Initialize database
	_, err := db.InitDB()
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
		if err := db.DB.Ping(context.Background()); err != nil {
			return c.String(http.StatusServiceUnavailable, "Cannot ping database")
		}
		log.Println("Pong!")
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/query", func(c echo.Context) error {
		var row string
		err := db.DB.QueryRow(context.Background(), "SELECT name from users LIMIT 1").Scan(&row)
		if err != nil {
			log.Printf("ERR: %v\n", err)
		}
		log.Println(row)
		return c.String(http.StatusOK, "query made!")
	})

	e.POST("/login", api.LoginUser)
	// Authentication
	// e.POST("/v1/login", loginUser)
	e.Logger.Fatal(e.Start(":1323"))
}
