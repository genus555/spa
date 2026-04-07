package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	startup		"github.com/genus555/spa/internal/init"
	cl			"github.com/genus555/spa/internal/clientloop"
)

func main() {
	fmt.Println("Welcome to the Simple Password Aggregator")
	godotenv.Load()
	if err := startup.GenerateEnv(); err != nil {
		log.Fatalf("Problem creating .env: %v", err)
	}
	key, err := startup.DecodeKey(os.Getenv("ENCRYPTION_KEY"))
	if err != nil {
		log.Fatalf("Problem getting encryption key: %v", err)
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
			err := cl.TestEncryptPW(key, inputs)
			if err != nil {fmt.Println(err)}
		case "register":
			err := cl.HandleRegister(key, inputs)
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