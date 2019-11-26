package main

import (
	"database/sql"
	"emergency-storage-manager/api"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	// Heroku port from env variable
	port := os.Getenv("PORT")

	// -----------------------
	// Database connection config
	// -----------------------

	// Heroku Postgres connection and ping
	dbConn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErrorAndPanic(err)

	defer dbConn.Close()

	err = dbConn.Ping()
	checkErrorAndPanic(err)
	fmt.Println("Correctly pinged DB")

	// -----------------------
	// Echo server definition
	// -----------------------

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))

	// -----------------------
	// Routes definition
	// -----------------------

	// Static endpoint to serve API doc
	// TODO: Write API doc
	e.Static("/docs", "static/docs")
	e.File("/favicon.ico", "static/favicon.ico")

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})

	// -----------------------
	// Categories
	// -----------------------

	var categories = api.Categories{Db: dbConn}

	// GET a single category based on ID
	e.GET("/categories/:id", categories.Get())
	// GET all categories
	e.GET("/categories", categories.GetAll())

	// -----------------------
	// Server Start
	// -----------------------

	e.Logger.Fatal(e.Start(":" + port))
}

func checkErrorAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}
