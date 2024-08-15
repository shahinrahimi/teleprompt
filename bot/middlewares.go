package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/models"
	"github.com/shahinrahimi/teleprompt/utils"
)

// Logging middleware log the command that received
func (b *Bot) Logging(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		b.l.Printf("Received command: %s", u.Message.Command())
		next(u, ctx)
	}
}

// RequireAuthorization middleware check if user has privilege to do the command (admin user)
func (b *Bot) RequireAuthorization(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		user, _ := b.s.GetUser(u.Message.From.ID)
		if user == nil || !user.IsAdmin {
			b.SendMessage(user.UserID, "You are not allowed")
			return
		}
		next(u, ctx)
	}
}

// RequireAuthentication middleware check if user has already registered to bot (user)
func (b *Bot) RequireAuthentication(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		user, err := b.s.GetUser(u.Message.From.ID)
		if user == nil {
			b.l.Printf("error getting user info: %v", err)
			b.SendMessage(u.Message.From.ID, "You are not registered!\nUsage: 'start'")
			return
		}
		next(u, ctx)
	}
}

// CheckRegistered middleware check if user is not registered
func (b *Bot) CheckRegistered(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		user, err := b.s.GetUser(u.Message.From.ID)
		if user != nil {
			b.l.Printf("error getting user info: %v", err)
			b.SendMessage(user.UserID, "You are already registered")
			return
		}
		next(u, ctx)
	}
}

func (b *Bot) ProvideUser(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		var user models.User
		user.UserID = u.Message.From.ID
		user.Username = u.Message.From.UserName
		c := context.WithValue(ctx, models.KeyUser{}, user)
		next(u, c)
	}
}

func (b *Bot) ProvidePrompt(next Handler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		var prompt models.Prompt
		parts := utils.ParseCommand(u.Message.Text)
		prompt.Title = parts[0]
		prompt.Body = parts[1]
		c := context.WithValue(ctx, models.KeyPrompt{}, prompt)
		next(u, c)
	}
}
