package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	key			[]byte
}

func NewDB(db *sql.DB, key []byte) DB {
	return DB{
		database:	db,
		key:		key,
	}
}