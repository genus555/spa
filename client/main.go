package main

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
	startup		"github.com/genus555/spa/internal/init"
	cl			"github.com/genus555/spa/internal/clientloop"
	h			"github.com/genus555/spa/internal/handlers"
	database	"github.com/genus555/spa/internal/database"
)

func main() {
	fmt.Println("Welcome to the Simple Password Aggregator")
	db_file, err := sql.Open("sqlite", "./passwords.db")
	if err != nil {
		log.Fatalf("Problem with database: %v", err)
	}
	defer db_file.Close()

	godotenv.Load()
	if err := startup.GenerateEnv(); err != nil {
		log.Fatalf("Problem creating .env: %v", err)
	}
	key, err := startup.DecodeKey(os.Getenv("ENCRYPTION_KEY"))
	if err != nil {
		log.Fatalf("Problem getting encryption key: %v", err)
	}

	db := database.NewDB(db_file, key)
	if err := db.Setup(); err != nil {
		log.Fatalf("Problem connecting with database: %v", err)
	}

	err = h.HandleCheckUser(db)
	if err != nil {
		log.Fatalf("Problem authenticating user: %v", err)
	}

	username, err := h.HandleGetUser(db)
	if err != nil {
		log.Fatalf("Problem getting user from database: %v", err)
	}
	db.Username = username

	cl.PrintCommands()

	for {
		inputs := cl.GetInput()
		if len(inputs) == 0 {
			continue
		}
		switch inputs[0] {
		case "register":
			err := h.HandleRegister(db, inputs)
			if err != nil {fmt.Println(err)}
		case "get":
			err := h.HandleGet(db, inputs)
			if err != nil {fmt.Println(err)}
		case "delete":
			err := h.HandleDelete(db, inputs)
			if err != nil {fmt.Println(err)}
		case "list":
			err := h.HandleList(db)
			if err != nil {fmt.Println(err)}
		case "transfer":
			err := h.HandleTransfer(db, inputs)
			if err != nil {fmt.Println(err)}
		case "deleteuser":
			err := h.HandleDeleteUser(db, inputs)
			if err != nil {fmt.Println(err)} else {return}
		case "help":
			cl.PrintCommands()
		case "quit":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Printf("\"%s\" is not a valid command\n", inputs[0])
		}
	}
}