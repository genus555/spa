package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	startup "github.com/genus555/spa/internal/init"
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
	fmt.Println(key)

	fmt.Println("Insert 2 factor here")
}