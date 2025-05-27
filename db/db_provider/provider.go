package db_provider

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tg-star-shop-bot-001/common/config"
	"tg-star-shop-bot-001/db/db_types"
)

type postgresGormOpener struct{}

func (o *postgresGormOpener) Open(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), config)
}

type DefaultDBProvider struct {
	configProvider config.ConfigProvider
	gormOpener     db_types.GormOpener
}

func NewDefaultDBProvider() *DefaultDBProvider {
	return NewDefaultDBProviderWithCustomOpener(
		config.NewDefaultConfigProvider(),
		&postgresGormOpener{},
	)
}

func NewDefaultDBProviderWithCustomOpener(configProvider config.ConfigProvider, gormOpener db_types.GormOpener) *DefaultDBProvider {
	return &DefaultDBProvider{
		configProvider: configProvider,
		gormOpener:     gormOpener,
	}
}

func (o *DefaultDBProvider) GetDB(ctx context.Context) (*gorm.DB, error) {
	conf, err := o.configProvider.GetAppConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	dsn := o.constructDSN(conf)
	db, err := o.gormOpener.Open(dsn, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	// TODO "gorm.io/gorm/logger"
	//gormLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		LogLevel: logger.Info,
	//	},
	//)
	//&gorm.Config{Logger: gormLogger})
	if conf.DB.IsDebugMode {
		db = db.Debug()
	}
	return db.WithContext(ctx), nil
}

// constructDSN creates a connection string from config
func (o *DefaultDBProvider) constructDSN(conf *config.AppConfig) string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.SslMode, conf.DB.User, conf.DB.Password, conf.DB.DbName)
}
