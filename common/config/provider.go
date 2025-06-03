package config

type LoadConfigFunc func(string) (*AppConfig, error)

type DefaultConfigProvider struct {
	fname    string
	loadFunc LoadConfigFunc
}

func NewDefaultConfigProvider() *DefaultConfigProvider {
	return NewDefaultConfigProviderWithCustomLoader("etc/config.yml", load)
}

func NewDefaultConfigProviderWithCustomLoader(fname string, loader LoadConfigFunc) *DefaultConfigProvider {
	return &DefaultConfigProvider{
		fname:    fname,
		loadFunc: loader,
	}
}

// GetAppConfig provides AppConfig from default config file
func (o *DefaultConfigProvider) GetAppConfig() (res *AppConfig, err error) {
	return o.loadFunc(o.fname)
}
