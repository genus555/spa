package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	key			[]byte
	Username	string
}

func NewDB(db *sql.DB, key []byte) *DB {
	database := DB{
		database:	db,
		key:		key,
		Username:	"",
	}
	return &database
}