package app

import (
	"log"
	"log/slog"
	"sync"

	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/meesooqa/storeque/common/applog"
	"github.com/meesooqa/storeque/common/config"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/db/provider"

	// Load lang phrases
	_ "github.com/meesooqa/storeque/tg"
)

type appDeps struct {
	config     *config.AppConfig
	logger     *slog.Logger
	langBundle *i18n.Bundle
	dbProvider provider.DBProvider
}

var (
	app  *appDeps
	once sync.Once
)

// GetInstance returns a singleton instance of App
func GetInstance() App {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		cp := config.NewDefaultConfigProvider()
		conf, err := cp.GetAppConfig()
		if err != nil {
			log.Fatal(err)
		}

		lp := applog.NewConsoleLoggerProvider(conf.Log)
		logger, cleanup := lp.GetLogger()
		defer cleanup()

		app = &appDeps{
			config:     conf,
			logger:     logger,
			langBundle: lang.NewBundle(),
			dbProvider: provider.NewDefaultDBProvider(),
		}
	})
	return app
}

func (o *appDeps) Config() *config.AppConfig {
	return o.config
}

func (o *appDeps) Logger() *slog.Logger {
	return o.logger
}

func (o *appDeps) LangBundle() *i18n.Bundle {
	return o.langBundle
}

func (o *appDeps) DBProvider() provider.DBProvider {
	return o.dbProvider
}
