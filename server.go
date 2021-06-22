package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	counter := 0
	e.GET("/", func(c echo.Context) error {
		counter += 1
		return c.String(http.StatusOK, fmt.Sprintf("Hello, World! %d", counter))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
