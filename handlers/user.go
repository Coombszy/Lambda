package handler

import (
	"fmt"
	"net/http"

	model "github.com/Coombszy/lambda/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Get DB
	db := h.DB

	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	sqlStatement := "INSERT INTO users (name, email, creation, password)VALUES ($1, $2, NOW(), $3)"
	res, err := db.Query(sqlStatement, u.Email, u.Name, u.Password)

	// Returns
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusCreated, res)
	}
}

func (h *Handler) ListUsers(c echo.Context) (err error) {
	// Get DB
	db := h.DB

	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Fetch all users
	sqlStatement := "SELECT name, email, password FROM users"
	rows, err := db.Query(sqlStatement)
	// Return if error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Bind SQL to user object
	var user model.User
	rows.Next()
	err = rows.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("scan error : " + err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Returns
	return c.JSON(http.StatusAccepted, user)
}
