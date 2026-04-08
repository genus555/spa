package database

import (
	"database/sql"
	"fmt"

	_	"modernc.org/sqlite"
	cl	"github.com/genus555/spa/internal/clientloop"
)

func (db *DB) Setup() error {
	_, err := db.database.Exec(`
		CREATE TABLE IF NOT EXISTS passwords (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			encrypted_password BLOB NOT NULL
			)
		`)
		return err
}

func (db *DB) addEntry (pw_name string, enc_pw []byte) error {
	var exists string
	if err := db.database.QueryRow("SELECT name FROM passwords WHERE name = ?", pw_name).Scan(&exists); err == nil {
		fmt.Println("A password under that name already exists")
		return nil
	}
	_, err := db.database.Exec("INSERT INTO passwords (name, encrypted_password) VALUES (?,?)", pw_name, enc_pw)
	if err != nil {return err}
	return nil
}

func (db *DB) getEntry (pw_name string) error {
	var enc_pw []byte
	err := db.database.QueryRow("SELECT encrypted_password FROM passwords WHERE name = ?", pw_name).Scan(&enc_pw)
	if err == sql.ErrNoRows {
		fmt.Printf("No password saved under the name \"%s\"\n", pw_name)
		return nil
	} else if err != nil {return err}
	
	pw, err := cl.DecryptPW(db.key, enc_pw)
	fmt.Println(pw)
	return nil
}

func (db *DB) listEntries () error {
	var count int
	if err := db.database.QueryRow("SELECT COUNT(*) FROM passwords").Scan(&count); err != nil {return err}
	if count == 0 {
		fmt.Println("No passwords saved")
		return nil
	}

	names, err := db.database.Query("SELECT name FROM passwords")
	if err != nil {return err}
	defer names.Close()

	for names.Next() {
		var name string
		if err := names.Scan(&name); err != nil {return err}
		fmt.Println(name)
	}
	return nil
}

func (db *DB) deleteEntry (pw_name string) error {
	var exists string
	if err := db.database.QueryRow("SELECT name FROM passwords WHERE name = ?", pw_name).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Password \"%s\" doesn't exist.\n", pw_name)
			return nil
		}
		return err
	}
	_, err := db.database.Exec("DELETE FROM passwords WHERE name = ?", pw_name)
	if err != nil {return err}
	
	fmt.Printf("Password \"%s\" has been deleted.\n", pw_name)
	return nil
}

func (db *DB) transferIn () error {
	fmt.Println("Transferring in")
	return nil
}

func (db *DB) transferOut() error {
	fmt.Println("Transferring out")
	return nil
}