package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.GET("/env", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, os.Environ())
	})

	e.Start(":1323")
}
