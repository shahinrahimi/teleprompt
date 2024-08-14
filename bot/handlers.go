package bot

import (
	"database/sql"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *TelegramBot) handleUpdate(u tgbotapi.Update) {
	switch {
	case u.Message != nil:
		t.handleMessage(u.Message)
	case u.CallbackQuery != nil:
		t.handleButton(u.CallbackQuery)
	}
}

func (t *TelegramBot) handleMessage(m *tgbotapi.Message) {
	userID := m.From.ID
	text := m.Text

	user, err := t.getUser(m.From.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			t.l.Printf("unexpected error in getting user: %v", err)
			return
		}
	}
	// handle newUser
	if strings.HasPrefix(text, "/") {
		switch {
		case user == nil:
			t.handleNewUserCommand(userID, text)
		case !user.IsAdmin:
			t.handleUserCommand(userID, text)
		case user.IsAdmin:
			t.handleAdminCommand(userID, text)
		default:
			t.l.Println("should never happened")
		}
	} else {
		switch {
		case user == nil:
			t.handleSendUsageForNewUser(userID)
		case !user.IsAdmin:
			t.handleSendUsageForUser(userID)
		case user.IsAdmin:
			t.handleSendUsageForAdmin(userID)
		default:
			t.l.Println("should never happened")
		}
	}

	// inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("View Prompts", "view_prompts"),
	// 		tgbotapi.NewInlineKeyboardButtonData("Add Prompt", "add_prompt"),
	// 		tgbotapi.NewInlineKeyboardButtonData("Delete Prompt", "delete_prompt"),
	// 	),
	// )

	// msg := tgbotapi.NewMessage(userID, "Choose an action:")
	// msg.ReplyMarkup = inlineKeyboard

	// t.b.Send(msg)

	// t.l.Printf("user with id: `%d` wrote `%s`", user.ID, text)

}

func (t *TelegramBot) handleButton(q *tgbotapi.CallbackQuery) {
	user := q.From
	text := q.Message

	t.l.Printf("user with id: `%d` wrote `%s`", user.ID, text)

	return

}
