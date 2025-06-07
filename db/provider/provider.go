package provider

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// DefaultDBProvider is a default implementation of the DBProvider interface
type DefaultDBProvider struct {
	db *sql.DB
}

// NewDefaultDBProvider creates a new instance of DefaultDBProvider
func NewDefaultDBProvider() *DefaultDBProvider {
	return &DefaultDBProvider{}
}

// Connect establishes a connection to the PostgreSQL database using environment variables
func (o *DefaultDBProvider) Connect() (*sql.DB, error) {
	var err error
	if o.db != nil {
		o.db.Close() // nolint
		o.db = nil
	}
	o.db, err = sql.Open("postgres", o.constructDsn())
	if err != nil {
		return nil, err
	}
	return o.db, nil
}

// constructDbUrl creates a connection string from config
func (o *DefaultDBProvider) constructDsn() string {
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
