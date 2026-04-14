package handlers

import (
	"fmt"

	cl			"github.com/genus555/spa/internal/clientloop"
	database	"github.com/genus555/spa/internal/database"
	auth		"github.com/genus555/spa/internal/auth"
)

func HandleRegister(db *database.DB, inputs []string) error {
	if len(inputs) != 3 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Register Usage: register [password_name] [password]")
		return nil
	}

	pw_id := inputs[1]
	pw := inputs[2]

	enc_pw, err := cl.EncryptPW(db.GetKey(), pw)
	if err != nil {return err}
	if err := db.AddEntry(pw_id, enc_pw); err != nil {return err}

	return nil
}

func HandleGet(db *database.DB, inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Get Usage: get [password_name]")
		return nil
	}
	pw_name := inputs[1]
	if err := db.GetEntry(pw_name); err != nil {return err}
	return nil
}

func HandleList(db *database.DB) error {
	err := db.ListEntries()
	if err != nil {return err}
	return nil
}

func HandleDelete(db *database.DB, inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Delete Usage: delete [password_name]")
		return nil
	}
	pw_name := inputs[1]
	if err := db.DeleteEntry(pw_name); err != nil {return err}
	return nil
}

func HandleTransfer(db *database.DB, inputs []string) error {
	if len(inputs) != 2 {
		fmt.Println("Too many/few arguments")
		fmt.Println("Transfer Usage: transfer [in/out]")
		return nil
	}
	switch inputs[1] {
	case "in":
		err := db.TransferIn()
		if err != nil {return err}
	case "out":
		err := db.TransferOut()
		if err != nil {return err}
	default:
		return fmt.Errorf("Not a valid transfer.\nTransfer Usage: transfer [in/out]")
	}
	return nil
}

func HandleAddUser(db *database.DB, username, otp_secret string) error {
	if err := db.AddUser(username, otp_secret); err != nil {return err}
	return nil
}

func HandleCheckUser(db *database.DB) error {
	count, err := db.CheckUserExist()
	if err != nil {return err}
	var username string
	var otp_secret string
	if count == 0 {
		fmt.Println("No saved user creating user profile...")
		username, otp_secret, err = auth.NewUser(db)
		if err != nil {return err}
		err = HandleAddUser(db, username, otp_secret)
		if err != nil {return err}
	} else {
		existing_user, err := HandleGetUser(db)
		if err != nil {return err}
		for {
			username = cl.GetUsername()
			if username != existing_user {
				fmt.Println("Incorrect username")
			} else {
				break
			}
		}
	}

	otp_secret, err = handleGetOTPSecret(db, username)
	if err != nil {return err}
	err = auth.CheckValid(db, username, otp_secret)
	if err != nil {return err}
	return nil
}

func HandleGetUser(db *database.DB) (string, error) {
	username, err := db.GetUser()
	if err != nil {return "", err}
	return username, nil
}

func HandleDeleteUser(db *database.DB, inputs []string) error {
	if len(inputs) != 3 {
		return fmt.Errorf("Too many/few arguments. Usage: deleteuser [username] [passcode]")
	}
	username := inputs[1]
	passcode := inputs[2]
	otp_secret, err := handleGetOTPSecret(db, username)
	if err != nil {return err}
	valid, err := auth.Valid(db, username, passcode, otp_secret)
	if err != nil {return err} else if !valid {return fmt.Errorf("Wrong username or passcode")}

	err = db.DeleteUser(username)
	if err != nil {return err}
	return nil
}

func handleGetOTPSecret(db *database.DB, username string) (string, error) {
	otp_secret, err := db.GetOTPSecret(username)
	if err != nil {return "", err}
	return otp_secret, nil
}