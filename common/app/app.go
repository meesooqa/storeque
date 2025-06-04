package app

import (
	"log"
	"log/slog"
	"sync"

	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/meesooqa/storeque/common/common_log"
	"github.com/meesooqa/storeque/common/config"
	"github.com/meesooqa/storeque/common/lang"
	"github.com/meesooqa/storeque/db/db_provider"

	// Load lang phrases
	_ "github.com/meesooqa/storeque/tg"
)

type appDeps struct {
	config     *config.AppConfig
	logger     *slog.Logger
	langBundle *i18n.Bundle
	dbProvider db_provider.DBProvider
}

var (
	app  *appDeps
	once sync.Once
)

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

		lp := common_log.NewConsoleLoggerProvider(conf.Log)
		logger, cleanup := lp.GetLogger()
		defer cleanup()

		app = &appDeps{
			config:     conf,
			logger:     logger,
			langBundle: lang.NewBundle(),
			dbProvider: db_provider.NewDefaultDBProvider(),
		}
	})
	return app
}

func (this *appDeps) Config() *config.AppConfig {
	return this.config
}

func (this *appDeps) Logger() *slog.Logger {
	return this.logger
}

func (this *appDeps) LangBundle() *i18n.Bundle {
	return this.langBundle
}

func (this *appDeps) DBProvider() db_provider.DBProvider {
	return this.dbProvider
}
