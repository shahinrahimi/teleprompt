package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/store"
)

type TelegramBot struct {
	l *log.Logger
	b *tgbotapi.BotAPI
	s store.Storage
}

func NewTelegramBot(l *log.Logger, s store.Storage, apiKey string) (*TelegramBot, error) {
	b, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		l.Printf("error creating a new bot api: %v", err)
		return nil, err
	}
	return &TelegramBot{
		l: l,
		b: b,
		s: s,
	}, nil
}

func (t *TelegramBot) RunAndListen(ctx context.Context) {
	uc := tgbotapi.NewUpdate(0)
	us := t.b.GetUpdatesChan(uc)
	go t.receiveUpdates(ctx, us)
}

func (t *TelegramBot) receiveUpdates(ctx context.Context, us tgbotapi.UpdatesChannel) {
	for {
		select {
		case <-ctx.Done():
			return
		case u := <-us:
			t.handleUpdate(u)
		}
	}
}

func (t *TelegramBot) sendMessage(userID int64, msgStr string) error {
	msg := tgbotapi.NewMessage(userID, msgStr)
	_, err := t.b.Send(msg)
	return err
}

func (t *TelegramBot) Shutdown() {
	t.l.Println("Bot shutting down...")
	t.b.StopReceivingUpdates()
}
