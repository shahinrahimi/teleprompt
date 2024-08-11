package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/store"
)

type TelegramBot struct {
	l *log.Logger
	b *tgbotapi.BotAPI
	s store.Storage
}

func NewTelegramBot(logger *log.Logger, store store.Storage, apiKey string) *TelegramBot {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		logger.Fatalf("error creating a new bot api: %v", err)
	}
	return &TelegramBot{
		l: logger,
		b: bot,
		s: store,
	}
}
