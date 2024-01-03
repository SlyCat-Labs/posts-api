package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Register PostgreSQL driver
)

// Connect establishes a connection to the database
func Connect(ctx context.Context, url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check connection health
	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
