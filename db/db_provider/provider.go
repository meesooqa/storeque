package db_provider

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

type DBProvider interface {
	OpenDB() (*sql.DB, error)
}

type DefaultDBProvider struct{}

func NewDefaultDBProvider() *DefaultDBProvider {
	return &DefaultDBProvider{}
}

func (o *DefaultDBProvider) OpenDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", o.constructDbUrl())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// constructDbUrl creates a connection string from config
func (o *DefaultDBProvider) constructDbUrl() string {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
}
