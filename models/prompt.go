package models

import "time"

type Prompt struct {
	ID        int       `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type KeyPrompt struct{}

const (
	CREATE_TABLE_PROMPTS string = `
		CREATE TABLE IF NOT EXISTS prompts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	SELECT_COUNT_PROMPTS            string = `SELECT COUNT(*) FROM prompts`
	SELECT_COUNT_PROMPTS_BY_USER_ID string = `SELECT COUNT(*) FROM prompts WHERE user_id = ?`
	SELECT_PROMPT                   string = `SELECT id, user_id, title, body, created_at FROM prompts WHERE id = ?`
	SELECT_PROMPTS                  string = `SELECT id, user_id, title, body, created_at FROM prompts`
	SELECT_PROMPTS_BY_USER_ID       string = `SELECT id, user_id, title, body, created_at FROM prompts WHERE user_id = ?`
	INSERT_PROMPT                   string = `INSERT INTO prompts (user_id, title, body, created_at) VALUES (?, ?, ?, ?)`
	DELETE_PROMPT                   string = `DELETE FROM prompts WHERE id = ?`
	DELETE_PROMPTS_BY_USER_ID       string = `DELETE FROM prompts WHERE user_id = ?`
)

// ToArgs returns user_id, title, body, created_at as value
// use for inserting to DB
func (p *Prompt) ToArgs() []interface{} {
	return []interface{}{p.UserID, p.Title, p.Body, p.CreatedAt}
}

// ToFields returns id, user_id, username, is_admin, created_at as reference
// use for scanning from DB
func (p *Prompt) ToFields() []interface{} {
	return []interface{}{&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt}
}
