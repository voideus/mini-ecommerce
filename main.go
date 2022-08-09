package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/voideus/mini-ecommerce/route"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Application starting")
	loadenv()
	log.Fatal(route.RunAPI(":5000"))
}
