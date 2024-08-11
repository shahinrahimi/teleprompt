package models

import "time"

type Prompt struct {
	ID        int       `json:"id"`
	UserID    float64   `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

const (
	CREATE_TABLE_PROMPTS string = `
		CREATE TABLE IF NOT EXISTS prompts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
)
