package app

import (
	"log/slog"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/meesooqa/storeque/common/config"
	"github.com/meesooqa/storeque/db/db_provider"
)

// App interface defines the methods to access application-level dependencies
type App interface {
	Config() *config.AppConfig
	Logger() *slog.Logger
	LangBundle() *i18n.Bundle
	DBProvider() db_provider.DBProvider
}
