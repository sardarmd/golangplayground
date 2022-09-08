package main

import (
	"log"

	"github.banking/sardarmd/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.Start()
}
