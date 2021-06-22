package handler

import "database/sql"

type (
	Handler struct {
		DB *sql.DB
	}
)

const (
	// Key (This should come from an config file)
	Key = "SuperSecretSecret"
)
