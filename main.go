package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/shahinrahimi/teleprompt/bot"
	"github.com/shahinrahimi/teleprompt/store"
)

func main() {
	fmt.Println("Hello")
	// create custom logger
	l := log.New(os.Stdout, "[TELEPROMPT] ", log.LstdFlags)

	// check .env file
	if err := godotenv.Load(); err != nil {
		l.Fatalf("error loading .env file: %v", err)
	}

	// check environmental variable
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		l.Fatal("wrong environmental variable")
	}

	// create a store
	s, err := store.NewSqliteStore(l)
	if err != nil {
		l.Fatalf("error creating new SqliteStore: %v", err)
	}
	defer s.CloseDB()

	// create tables if not exists
	if err := s.Init(); err != nil {
		l.Fatalf("error initializing store: %v", err)
	}

	// create a bot instance
	t, err := bot.NewTelegramBot(l, s, apiKey)
	if err != nil {
		l.Fatalf("error creating a new telegram bot: %v", err)
	}

	// create context bot to received updates and gracefully shutdown
	ctx := context.WithoutCancel(context.Background())
	go func() {
		l.Println("Bot started and running and listen for updates.")
		t.RunAndListen(ctx)
	}()

	// create signal
	c := make(chan os.Signal, 1)
	// filter all other signal
	signal.Notify(c, os.Interrupt)

	// block until a signal received
	rc := <-c
	l.Println("go signal", rc)

	// gracefully shutdown bot, waiting max 30 secs
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.Shutdown()

}
