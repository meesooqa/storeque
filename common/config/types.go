package config

type ConfigProvider interface {
	GetAppConfig() (res *AppConfig, err error)
}
