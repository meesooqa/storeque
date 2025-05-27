package app

import (
	"database/sql"
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
	Config *config.AppConfig
	Logger *slog.Logger
	Lang   lang.Localization
	DB     *sql.DB
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

	loc := lang.NewLang(logger, conf.System.DefaultLangTag)

	dbp := db_provider.NewDefaultDBProvider()
	db, err := dbp.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	return &AppDeps{
		Config: conf,
		Logger: logger,
		Lang:   loc,
		DB:     db,
	}
}
