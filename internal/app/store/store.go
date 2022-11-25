package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Open() error {
	fmt.Println(s.config.DatabaseURL)
	db, err := sql.Open("postgres", "host=localhost user=admin password=lillol000 dbname=restipa_dev sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}
func (s *Store) Close() {
	s.db.Close()
}
