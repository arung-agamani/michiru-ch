package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}