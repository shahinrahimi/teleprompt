package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/teleprompt/models"
)

func (b *Bot) HandleRegisterUser(u *tgbotapi.Update) error {
	var user models.User
	user.UserID = u.Message.From.ID
	user.Username = u.Message.From.UserName
	if err := b.s.CreateUser(&user); err != nil {
		return err
	}
	b.SendMessage(user.UserID, "Registered Successfully")
	return nil
}

func (b *Bot) MakeHandlerBotFunc(f BotHandler) Handler {
	return func(u *tgbotapi.Update) {
		if err := f(u); err != nil {
			b.l.Printf("we have error %v", err)
		}
	}
}
