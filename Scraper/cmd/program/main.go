package main

import (
	"Scraper/internal/bot"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	err = bot.InitBot(token)
	if err != nil {
		log.Printf("Error initializing bot: %v", err)
		log.Fatal(err)
	}
}
