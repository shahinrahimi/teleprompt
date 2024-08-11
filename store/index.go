package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shahinrahimi/teleprompt/models"
)

type SqliteStore struct {
	l  *log.Logger
	db *sql.DB
}

type Storage interface {
	GetUser(user_id int64) (*models.User, error)
	GetUsers() ([]models.User, error)
	CreateUser(u *models.User) error
	DeleteUser(user_id int64) error

	GetPrompt(id int) (*models.Prompt, error)
	GetPrompts() ([]models.Prompt, error)
	CreatePrompt(p *models.Prompt) error
	DeletePrompt(id int) error
}

func NewSqliteStore(l *log.Logger) (*SqliteStore, error) {
	// create directory if needed
	if err := os.MkdirAll("db", 0755); err != nil {
		l.Printf("Unable to create a directory for DB: %v", err)
		return nil, err
	}
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		l.Printf("Unable to connect to DB: %v", err)
		return nil, err
	}
	l.Println("DB Connected!")
	return &SqliteStore{
		l:  l,
		db: db,
	}, nil
}

func (s *SqliteStore) Init() error {
	if _, err := s.db.Exec(models.CREATE_TABLE_USERS); err != nil {
		s.l.Fatalf("error creating users table: %v", err)
		return err
	}
	if _, err := s.db.Exec(models.CREATE_TABLE_PROMPTS); err != nil {
		s.l.Fatalf("error creating prompts table: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) CloseDB() error {
	return s.db.Close()
}
