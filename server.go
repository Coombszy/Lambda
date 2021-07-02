package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/Coombszy/lambda/handlers"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {

	// Basic HTTP
	// Echo instance
	e := echo.New()

	// -----------
	// Middleware
	// -----------

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			// Skip logging on dev endpoint
			return c.Path() == "/dev"
		},
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/login" || c.Path() == "/signup" || c.Path() == "/dev" {
				return true
			}
			return false
		},
	}))

	// -----------
	// Database
	// -----------

	// Database connection
	connStr := "postgres://lambda:lambda_password@localhost/lambda_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Database migration
	m, err := migrate.New(
		"file://db",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// -----------
	// Routing
	// -----------

	// Dev
	counter := 0
	e.GET("/dev", func(c echo.Context) error {
		counter += 1
		return c.JSON(http.StatusOK, fmt.Sprintf("Hello, World! %d", counter))
	})

	// Users
	e.POST("/signup", h.Signup)
	// e.POST("/login", h.Login)

	// Start Echo and Logger
	e.Logger.Fatal(e.Start(":1323"))
}
