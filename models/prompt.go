package models

import "time"

type Prompt struct {
	ID        string    `json:"id"`
	UserID    float64   `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
