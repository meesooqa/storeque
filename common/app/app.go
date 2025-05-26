package app

import (
	"log"
	"log/slog"

	"tg-star-shop-bot-001/common/common_log"
	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/common/lang"
	_ "tg-star-shop-bot-001/tg"
)

type AppDeps struct {
	Config *config.AppConfig
	Logger *slog.Logger
	Lang   lang.Localization
}

func NewAppDeps() *AppDeps {
	cp := config.NewDefaultConfigProvider()
	conf, err := cp.GetAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	lp := common_log.NewConsoleLoggerProvider(conf.Log)
	logger, cleanup := lp.GetLogger()
	defer cleanup()

	loc := lang.NewLang("en")

	return &AppDeps{
		Config: conf,
		Logger: logger,
		Lang:   loc,
	}
}
