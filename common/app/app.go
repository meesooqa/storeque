package app

import (
	"log"
	"log/slog"
	"sync"

	"github.com/joho/godotenv"

	"tg-star-shop-bot-001/common/common_log"
	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/common/lang"
	"tg-star-shop-bot-001/db/db_provider"

	// Load lang phrases
	_ "tg-star-shop-bot-001/tg"
)

type App interface {
	ChangeLang(langTag string)
	Config() *config.AppConfig
	Logger() *slog.Logger
	Lang() lang.Localization
	DBProvider() db_provider.DBProvider
}

type appDeps struct {
	config     *config.AppConfig
	logger     *slog.Logger
	lang       lang.Localization
	dbprovider db_provider.DBProvider
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
			lang:       lang.NewLang(logger, conf.System.DefaultLangTag),
			dbprovider: db_provider.NewDefaultDBProvider(),
		}
	})
	return app
}

func (a *appDeps) ChangeLang(langTag string) {
	a.lang = lang.NewLang(a.Logger(), langTag)
}

func (a *appDeps) Config() *config.AppConfig {
	return a.config
}

func (a *appDeps) Logger() *slog.Logger {
	return a.logger
}

func (a *appDeps) Lang() lang.Localization {
	return a.lang
}

func (a *appDeps) DBProvider() db_provider.DBProvider {
	return a.dbprovider
}

/*func NewAppDeps() *AppDeps {
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
}*/
