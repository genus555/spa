package database

import (
	"database/sql"
	"fmt"
	"os"

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
	db.database.Close()

	wd, err := os.Getwd()
	if err != nil {return err}
	src_dir := wd + "/transfer"

	_, err = os.Stat(src_dir+"/.env")
	if os.IsNotExist(err) {
		return fmt.Errorf("Missing env file")
	} else if err != nil {return err}
	_, err = os.Stat(src_dir+"/passwords.db")
	if os.IsNotExist(err) {
		return fmt.Errorf("Missing password database file")
	} else if err != nil {return err}
	if err := cl.CopyFile(src_dir+"/.env", wd+"/.env"); err != nil {return err}
	if err := cl.CopyFile(src_dir+"/passwords.db", wd+"/passwords.db"); err != nil {return err}
	
	db.database, err =  sql.Open("sqlite", "./passwords.db")
	if err != nil {return err}

	fmt.Println("Information has been transferred")
	return nil
}

func (db *DB) transferOut() error {
	db.database.Close()

	wd, err := os.Getwd()
	if err != nil {return err}
	transfer_dir := wd + "/transfer"

	path_info, err := os.Stat(transfer_dir)
	if os.IsNotExist(err) {
		err = os.Mkdir("transfer", 0755)
		if err != nil {return err}
	} else if err != nil {return err} else if !path_info.IsDir() {
		return fmt.Errorf("Path doesn't lead to a directory")
	}
	if err := cl.CopyFile(wd+"/.env", transfer_dir+"/.env"); err != nil {return err}
	if err := cl.CopyFile(wd+"/passwords.db", transfer_dir+"/passwords.db"); err != nil {return err}

	db.database, err = sql.Open("sqlite", "./passwords.db")
	if err != nil {return err}

	fmt.Println("Transfer folder has been made")
	return nil
}