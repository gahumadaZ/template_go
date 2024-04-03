package main

import (
	"log"

	"github.com/joho/godotenv"
	"zapping.com/my_project/pkg/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.New()
	log.Println(cfg)
}
