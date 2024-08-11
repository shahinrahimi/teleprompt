package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/teleprompt/bot"
	"github.com/shahinrahimi/teleprompt/store"
)

func main() {
	fmt.Println("Hello")
	// create custom logger
	logger := log.New(os.Stdout, "[TELEPROMPT] ", log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading .env file: %v", err)
	}

	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		logger.Fatal("wrong environmental variable")
	}

	s := store.NewSqliteStore(logger)
	if err := s.Init(); err != nil {
		logger.Fatalf("error initializing store: %v", err)
	}

	bot.NewTelegramBot(logger, s, apiKey)

}
