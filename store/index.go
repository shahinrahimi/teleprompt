package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	logger *log.Logger
	db     *sql.DB
}

type Storage interface {
}

func NewSqliteStore(logger *log.Logger) *SqliteStore {
	// create directory if needed
	if err := os.MkdirAll("db", 0755); err != nil {
		logger.Fatalf("Unable to create a directory for DB: %v", err)
	}
	db, err := sql.Open("sqlite3", "./db/mydb.db")
	if err != nil {
		logger.Fatalf("Unable to connect to DB: %v", err)
	}
	logger.Println("DB Connected!")
	return &SqliteStore{
		logger: logger,
		db:     db,
	}
}

func (s *SqliteStore) Init() error {
	return nil
}

func (s *SqliteStore) CloseDB() error {
	return s.db.Close()
}
