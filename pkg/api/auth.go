package api

import (
	"context"
	"log"
	"net/http"

	"github.com/asajaroff/microblog/pkg/db"

	"github.com/labstack/echo/v4"
)

func init() {
	// Init database
	//db = db.GetDB()
}

func LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	//saveSession := ctx.FormValue("save_session")

	var storedPassword string

	// query := "SELECT password FROM users WHERE primaryEmail = $1"

	err := db.DB.QueryRow(context.Background(),
		"SELECT password FROM users WHERE primaryEmail = $1", username).Scan(&storedPassword)

	if err != nil {
		log.Printf("ERR: $v\n", err)
		c.JSON(http.StatusForbidden, "User not found")
		return nil
	}

	if storedPassword != password {
		log.Println("INFO: Missmatching password for user - ", username)
		c.JSON(http.StatusForbidden, "Wrong username or password")
		return nil
	}
	c.JSON(http.StatusOK, "Welcome back :-)")
	return nil
}

func getUser(c echo.Context) error {
	c.String(http.StatusOK, "query made!")
	return nil
}
