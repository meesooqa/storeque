package app

import (
	"log/slog"

	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/common/lang"
	"tg-star-shop-bot-001/db/db_provider"
)

type App interface {
	ChangeLang(langTag string)
	Config() *config.AppConfig
	Logger() *slog.Logger
	Lang() lang.Localization
	DBProvider() db_provider.DBProvider
}
