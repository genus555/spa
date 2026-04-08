package database

import (
	"fmt"

	cl		"github.com/genus555/spa/internal/clientloop"
)

func (db *DB) HandleRegister(inputs []string) error {
	if len(inputs) != 3 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Register Usage: register [password_name] [password]")
		return nil
	}

	pw_id := inputs[1]
	pw := inputs[2]

	enc_pw, err := cl.EncryptPW(db.key, pw)
	if err != nil {return err}
	if err := db.addEntry(pw_id, enc_pw); err != nil {return err}

	return nil
}

func (db *DB) HandleGet(inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Get Usage: get [password_name]")
		return nil
	}
	pw_name := inputs[1]
	if err := db.getEntry(pw_name); err != nil {return err}
	return nil
}

func (db *DB) HandleList() error {
	err := db.listEntries()
	if err != nil {return err}
	return nil
}

func (db *DB) HandleDelete(inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Delete Usage: delete [password_name]")
		return nil
	}
	pw_name := inputs[1]
	if err := db.deleteEntry(pw_name); err != nil {return err}
	return nil
}

func (db *DB) HandleTransfer(inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Transfer Usage: transfer [in/out]")
		return nil
	}
	switch inputs[1] {
	case "in":
		err := db.transferIn()
		if err != nil {return err}
	case "out":
		err := db.transferOut()
		if err != nil {return err}
	default:
		return fmt.Errorf("Not a valid transfer.\nTransfer Usage: transfer [in/out]")
	}
	return nil
}