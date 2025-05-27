package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	dbURL := "postgres://user:password@localhost:5432/ssbot_db?sslmode=disable"
	runMigrations(dbURL, migrationsFS)
}

func runMigrations(dbURL string, fs embed.FS) {
	// 1. Подключаемся к БД
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}
	defer db.Close()

	// 2. Создаём драйвер для миграций
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migrate driver: %v", err)
	}

	// 3. Инициализируем источник миграций из embed.FS через iofs
	src, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatalf("failed to create iofs source: %v", err)
	}

	// 4. Создаём объект migrate
	m, err := migrate.NewWithInstance(
		"iofs", src,
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrate: %v", err)
	}

	// 5. Применяем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migrate up failed: %v", err)
	}

	fmt.Println("✔️ Миграции успешно применены")
}
