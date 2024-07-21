package main

import (
	"log"

	chronicleService "github.com/SomethingSexy/chronicle/internal/chronicle/service"
)

func main() {

	chronicleService.NewService()
	log.Println("verify running")
}
