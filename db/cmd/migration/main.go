package main

import (
	"embed"
	"log/slog"

	"tg-star-shop-bot-001/common/app"
	"tg-star-shop-bot-001/db/migration"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	appDeps := app.NewAppDeps()

	migrator := migration.NewMigrator(appDeps.DB)
	err := migrator.Migrate(migrationsFS, "migrations")
	if err != nil {
		appDeps.Logger.Error("migration error", slog.Any("error", err))
	}
}
