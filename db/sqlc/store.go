package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}

}
