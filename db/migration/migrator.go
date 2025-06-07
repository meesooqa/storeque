package migration

import (
	"database/sql"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Migrator is responsible for applying database migrations
type Migrator struct {
	db *sql.DB
}

// NewMigrator creates a new instance of Migrator
func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

// Migrate applies the migrations from the embedded filesystem to the PostgreSQL database
func (o *Migrator) Migrate(fs embed.FS, path string) error {
	driver, err := postgres.WithInstance(o.db, &postgres.Config{})
	if err != nil {
		return err
	}

	src, err := iofs.New(fs, path)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
