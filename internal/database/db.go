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
	if err != nil {fmt.Errorf("Problem creating passwords database")}
	_, err = db.database.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			username TEXT NOT NULL,
			otp_secret TEXT NOT NULL,
			enc_key TEXT NOT NULL
			)
		`)
	if err != nil {fmt.Errorf("Problem creating user database")}
	return nil
}

func (db *DB) GetKey() []byte {
	return db.key
}

func (db *DB) AddEntry(pw_name string, enc_pw []byte) error {
	var exists string
	if err := db.database.QueryRow("SELECT name FROM passwords WHERE name = ?", pw_name).Scan(&exists); err == nil {
		fmt.Println("A password under that name already exists")
		return nil
	}
	_, err := db.database.Exec("INSERT INTO passwords (name, encrypted_password) VALUES (?,?)", pw_name, enc_pw)
	if err != nil {return err}
	return nil
}

func (db *DB) GetEntry(pw_name string) error {
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

func (db *DB) ListEntries() error {
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

func (db *DB) DeleteEntry(pw_name string) error {
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

func (db *DB) CheckUserExist() (int, error) {
	var count int
	if err := db.database.QueryRow("SELECT COUNT(*) FROM user").Scan(&count); err != nil {return 0, err}
	return count, nil
}

func (db *DB) AddUser(username, otp_secret string) error {
	enc_key := cl.EncodeEncKey(db.GetKey())
	_, err := db.database.Exec("INSERT INTO user (username, otp_secret, enc_key) VALUES (?,?,?)", username, otp_secret, enc_key)
	if err != nil {return err}
	return nil
}

func (db *DB) GetUser() (string, error) {
	var username string
	err := db.database.QueryRow("SELECT username FROM user").Scan(&username)
	if err != nil {return "", err}
	return username, nil
}

func (db *DB) DeleteUser(username string) error {
	_, err := db.database.Exec("DELETE FROM user WHERE username = ?", username)
	if err != nil {return err}

	fmt.Printf("User \"%s\" has been deleted.\n", username)
	return nil
}

func (db *DB) GetOTPSecret(username string) (string, error) {
	var otp_secret string
	err := db.database.QueryRow("SELECT otp_secret FROM user WHERE username = ?", username).Scan(&otp_secret)
	if err != nil {return "", err}
	return otp_secret, nil
}

func (db *DB) TransferIn() error {
	db.database.Close()

	wd, err := os.Getwd()
	if err != nil {return err}
	src_dir := wd + "/transfer"

	_, err = os.Stat(src_dir+"/passwords.db")
	if os.IsNotExist(err) {
		return fmt.Errorf("Missing password database file")
	} else if err != nil {return err}
	if err := cl.CopyFile(src_dir+"/passwords.db", wd+"/passwords.db"); err != nil {return err}
	
	db.database, err =  sql.Open("sqlite", "./passwords.db")
	if err != nil {return err}

	fmt.Println("Information has been transferred")
	return nil
}

func (db *DB) TransferOut() error {
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
	if err := cl.CopyFile(wd+"/passwords.db", transfer_dir+"/passwords.db"); err != nil {return err}

	db.database, err = sql.Open("sqlite", "./passwords.db")
	if err != nil {return err}

	fmt.Println("Transfer folder has been made")
	return nil
}