package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/models"
)

func (b *Bot) HandleRegisterUser(u *tgbotapi.Update, ctx context.Context) error {
	user := ctx.Value(models.KeyUser{}).(models.User)
	if err := b.s.CreateUser(&user); err != nil {
		return err
	}
	b.SendMessage(user.UserID, "Registered Successfully")
	return nil
}

func (b *Bot) HandleUnregisterUser(u *tgbotapi.Update, ctx context.Context) error {
	user := ctx.Value(models.KeyUser{}).(models.User)
	// TODO delete all prompts for user
	if err := b.s.DeleteUser(user.UserID); err != nil {
		return err
	}
	b.SendMessage(user.UserID, "Unregistered Successfully")
	return nil
}

func (b *Bot) HandleAddPrompt(u *tgbotapi.Update, ctx context.Context) error {
	prompt := ctx.Value(models.KeyPrompt{}).(models.Prompt)
	user := ctx.Value(models.KeyUser{}).(models.User)
	prompt.UserID = user.UserID
	b.l.Printf("prompt title: %s body: %s", prompt.Title, prompt.Body)
	// if err := b.s.CreatePrompt()
	return nil
}

func (b *Bot) HandleDeletePrompt(u *tgbotapi.Update, ctx context.Context) error {
	return nil
}

func (b *Bot) HandleViewUsers(u *tgbotapi.Update, ctx context.Context) error {
	return nil
}

func (b *Bot) MakeHandlerBotFunc(f ErrorHandler) Handler {
	return func(u *tgbotapi.Update, ctx context.Context) {
		if err := f(u, ctx); err != nil {
			b.l.Printf("we have error %v", err)
		}
	}
}
