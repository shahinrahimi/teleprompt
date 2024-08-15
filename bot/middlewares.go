package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Logging middleware log the command that received
func (b *Bot) Logging(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		b.l.Printf("Received command: %s", u.Message.Command())
		next(u)
	}
}

// RequireAuthorization middleware check if user has privilege to do the command (admin user)
func (b *Bot) RequireAuthorization(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, _ := b.s.GetUser(u.Message.From.ID)
		if user == nil || !user.IsAdmin {
			b.SendMessage(user.UserID, "You are not allowed")
			return
		}
		next(u)
	}
}

// RequireAuthentication middleware check if user has already registered to bot (user)
func (b *Bot) RequireAuthentication(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, err := b.s.GetUser(u.Message.From.ID)
		if user == nil {
			b.l.Printf("error getting user info: %v", err)
			b.SendMessage(u.Message.From.ID, "You are not registered!\nUsage: 'start'")
			return
		}
		next(u)
	}
}

// CheckRegistered middleware check if user is not registered
func (b *Bot) CheckRegistered(next Handler) Handler {
	return func(u *tgbotapi.Update) {
		user, err := b.s.GetUser(u.Message.From.ID)
		if user != nil {
			b.l.Printf("error getting user info: %v", err)
			b.SendMessage(user.UserID, "You are already registered")
			return
		}
		next(u)
	}
}
