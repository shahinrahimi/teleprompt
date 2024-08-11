package models

import "time"

type User struct {
	ID        int       `json:"id"`
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	CREATE_TABLE_USERS string = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			username TEXT,
			is_admin bool,
			created_at TIMESTAMP NOT NULL,
		)
	`
)
