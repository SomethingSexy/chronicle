package main

import (
	"log"

	chronicleService "github.com/SomethingSexy/chronicle/internal/chronicle/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chronicleService.NewService()
	log.Println("Service running")
}
