package api

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type Categories struct {
	Db *sql.DB
}

func (c Categories) Get() echo.HandlerFunc {
	return func(context echo.Context) error {
		id := context.Param("id")
		// TODO: Implement db calls
		return context.String(http.StatusNotImplemented, fmt.Sprintf("Requested id: %s", id))
	}
}

func (c Categories) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.String(http.StatusNotImplemented, "Will return all categories")
	}
}
