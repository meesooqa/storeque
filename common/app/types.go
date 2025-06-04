package app

import (
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/meesooqa/tg-star-shop-bot/common/config"
	"github.com/meesooqa/tg-star-shop-bot/db/db_provider"
)

type App interface {
	Config() *config.AppConfig
	Logger() *slog.Logger
	LangBundle() *i18n.Bundle
	DBProvider() db_provider.DBProvider
}
