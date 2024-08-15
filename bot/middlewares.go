package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) Logging(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		b.l.Printf("Received command: %s", u.Message.Command())
		next(u)
	}
}

func (b *Bot) RequireAdmin(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, _ := b.s.GetUser(u.Message.From.ID)
		if user == nil || !user.IsAdmin {
			b.SendMessage(user.UserID, "You are not allowed")
			return
		}
		next(u)
	}
}

func (b *Bot) RequireAuth(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, _ := b.s.GetUser(u.Message.From.ID)
		if user == nil {
			b.SendMessage(user.UserID, "You are not registered!\nUsage: 'start'")
			return
		}
		next(u)
	}
}

func (b *Bot) CheckRegistered(next Handler) Handler {
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
