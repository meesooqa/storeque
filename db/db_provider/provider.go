package db_provider

import (
	"database/sql"
	"fmt"
	"os"
)

type DBProvider interface {
	GetDB() (*sql.DB, error)
}

type DefaultDBProvider struct{}

func NewDefaultDBProvider() *DefaultDBProvider {
	return &DefaultDBProvider{}
}

func (o *DefaultDBProvider) GetDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", o.constructDbUrl())
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}

// constructDbUrl creates a connection string from config
func (o *DefaultDBProvider) constructDbUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}
