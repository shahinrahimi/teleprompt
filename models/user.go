package models

import "time"

type User struct {
	ID        int       `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

type KeyUser struct{}

const (
	CREATE_TABLE_USERS string = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			user_id INTEGER UNIQUE,
			username TEXT,
			is_admin bool,
			created_at TIMESTAMP NOT NULL
		)
	`
	SELECT_COUNT_USERS     string = `SELECT COUNT(*) FROM users`
	SELECT_USER            string = `SELECT id, user_id, username, is_admin, created_at FROM users WHERE id = ?`
	SELECT_USER_BY_USER_ID string = `SELECT id, user_id, username, is_admin, created_at FROM users WHERE user_id = ?`
	SELECT_USERS           string = `SELECT id, user_id, username, is_admin, created_at FROM users`
	INSERT_USER            string = `INSERT INTO users (user_id, username, is_admin, created_at) VALUES (?, ?, ?, ?)`
	DELETE_USER            string = `DELETE FROM users WHERE id = ?`
	DELETE_USER_BY_USER_ID string = `DELETE FROM users WHERE user_id = ?`
)

// ToArgs returns user_id, username, is_admin, created_at as value
// use for inserting to DB
func (u *User) ToArgs() []interface{} {
	return []interface{}{u.UserID, u.Username, u.IsAdmin, u.CreatedAt}
}

// ToFields returns id, user_id, username, is_admin, created_at as reference
// use for scanning from DB
func (u *User) ToFields() []interface{} {
	return []interface{}{&u.ID, &u.UserID, &u.Username, &u.IsAdmin, &u.CreatedAt}
}
