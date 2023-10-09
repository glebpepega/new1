package main

import (
	"log"

	"github.com/glebpepega/new1/internal/apiserver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiserver.New().Start()
}
