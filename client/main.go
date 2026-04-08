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

	fmt.Println("Insert 2 factor here")

	cl.PrintCommands()

	for {
		inputs := cl.GetInput()
		if len(inputs) == 0 {
			continue
		}
		switch inputs[0] {
		case "test":
			fmt.Println("no tests atm")
		case "register":
			err := db.HandleRegister(inputs)
			if err != nil {fmt.Println(err)}
		case "get":
			err := db.HandleGet(inputs)
			if err != nil {fmt.Println(err)}
		case "delete":
			err := db.HandleDelete(inputs)
			if err != nil {fmt.Println(err)}
		case "list":
			err := db.HandleList()
			if err != nil {fmt.Println(err)}
		case "transfer":
			err := db.HandleTransfer(inputs)
			if err != nil {fmt.Println(err)}
		case "help":
			cl.PrintCommands()
		case "quit":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Printf("\"%s\" is not a valid command\n", inputs[0])
			continue
		}
	}
}