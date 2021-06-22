package handler

import (
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
	sqlStatement := "INSERT INTO users (email, name, password)VALUES ($1, $2, $3)"
	res, err := db.Query(sqlStatement, u.Email, u.Name, u.Password)

	// Returns
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusCreated, res)
	}
}
