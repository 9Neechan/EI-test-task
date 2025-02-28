package db

import (
	"database/sql"

	sqlc "github.com/9Neechan/EI-test-task/stats-server/internal/db/sqlc"
)

type Store interface {
	sqlc.Querier
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

// provides all funcs to exec db queries and transactions
// composition for extending Queries functionality
type SQLStore struct {
	*sqlc.Queries
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}
