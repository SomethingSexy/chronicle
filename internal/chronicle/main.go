package main

import (
	"log"
	"os"

	chronicleService "github.com/SomethingSexy/chronicle/internal/chronicle/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = chronicleService.NewService()
	if err != nil {
		log.Println("Error starting service", err)
		os.Exit(1)
	}
	log.Println("Service running")
}
