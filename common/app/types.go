package app

import (
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/db/db_provider"
)

type App interface {
	Config() *config.AppConfig
	Logger() *slog.Logger
	LangBundle() *i18n.Bundle
	DBProvider() db_provider.DBProvider
}
