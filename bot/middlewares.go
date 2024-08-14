package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) Logging(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		b.l.Printf("Received command: %s", u.Message.Command())
		next(u)
	}
}

func (b *Bot) CheckUser(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, err := b.s.GetUser(u.Message.From.ID)
		if user != nil {
			b.l.Printf("Received err: %v", err)
			b.SendMessage(user.UserID, "You are already registered")
			return
		}
		next(u)
	}
}
