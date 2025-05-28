package app

import (
	"log"
	"log/slog"

	"github.com/joho/godotenv"

	"tg-star-shop-bot-001/common/common_log"
	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/common/lang"
	"tg-star-shop-bot-001/db/db_provider"

	// Lang
	_ "tg-star-shop-bot-001/tg"
)

type AppDeps struct {
	Config     *config.AppConfig
	Logger     *slog.Logger
	Lang       lang.Localization
	DBProvider db_provider.DBProvider
}

func NewAppDeps() *AppDeps {
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

	return &AppDeps{
		Config:     conf,
		Logger:     logger,
		Lang:       lang.NewLang(logger, conf.System.DefaultLangTag),
		DBProvider: db_provider.NewDefaultDBProvider(),
	}
}
