package db

import (
	"database/sql"

	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

// Store interface defines the methods for querying the database
type Store interface {
	sqlc.Querier
}

// NewStore creates a new instance of the Store interface
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

// SQLStore provides all functions to execute database queries and transactions
// It uses composition to extend the functionality of the Queries interface
type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

// NewSQLStore creates a new instance of the SQLStore
func NewSQLStore(db *sql.DB) *SQLStore {
	store := &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}

	return store
}
