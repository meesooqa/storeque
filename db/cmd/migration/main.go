package main

import (
	"embed"
	"log/slog"

	"github.com/meesooqa/tg-star-shop-bot/common/app"
	"github.com/meesooqa/tg-star-shop-bot/db/migration"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	appDeps := app.GetInstance()

	db, err := appDeps.DBProvider().Connect()
	if err != nil {
		appDeps.Logger().Error("db connection error", slog.Any("error", err))
	}
	defer db.Close()

	migrator := migration.NewMigrator(db)
	err = migrator.Migrate(migrationsFS, "migrations")
	if err != nil {
		appDeps.Logger().Error("migration error", slog.Any("error", err))
	}
}
