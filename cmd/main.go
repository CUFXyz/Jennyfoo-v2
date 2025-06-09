package main

import (
	"log"
	server "v2/internal/Server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	server.JFServerSetup(":8080").Run()
}
