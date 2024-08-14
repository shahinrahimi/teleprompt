package bot

import (
	"github.com/shahinrahimi/teleprompt/models"
)

// getUser get user from DB will return sql.ErrNoRows if the user is not found in DB
func (t *TelegramBot) getUser(userID int64) (*models.User, error) {
	user, err := t.s.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, err
}

// registerUser register user in DB with unique userID
func (t *TelegramBot) registerUser(userID int64, username string) error {
	var user models.User
	user.UserID = userID
	user.Username = username
	if err := t.s.CreateUser(&user); err != nil {
		return err
	}
	return nil
}
