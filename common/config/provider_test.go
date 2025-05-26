package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func MockLoad(fileName string) (*AppConfig, error) {
	if fileName == "etc/config.yml" {
		return &AppConfig{
			DB: &DbConfig{
				Host:     "localhost",
				Port:     5432,
				SslMode:  "disable",
				User:     "testuser",
				Password: "testpass",
				DbName:   "testdb",
			},
		}, nil
	} else if fileName == "etc/invalid.yml" {
		return nil, errors.New("failed to read config file")
	} else if fileName == "etc/custom.yml" {
		return &AppConfig{
			DB: &DbConfig{
				Host:     "customhost",
				Port:     5433,
				SslMode:  "prefer",
				User:     "customuser",
				Password: "custompass",
				DbName:   "customdb",
			},
		}, nil
	}

	return nil, errors.New("unknown config file")
}

func TestNewDefaultConfigProvider(t *testing.T) {
	provider := NewDefaultConfigProvider()

	assert.NotNil(t, provider)
	assert.Equal(t, "etc/config.yml", provider.fname)
	assert.NotNil(t, provider.loadFunc)
}

func TestNewConfigProviderWithCustomLoader(t *testing.T) {
	customFilename := "custom/path.yml"
	provider := NewDefaultConfigProviderWithCustomLoader(customFilename, MockLoad)

	assert.NotNil(t, provider)
	assert.Equal(t, customFilename, provider.fname)
	assert.NotNil(t, provider.loadFunc)
}

func TestDefaultConfigProvider_GetConf(t *testing.T) {
	provider := NewDefaultConfigProviderWithCustomLoader("etc/config.yml", MockLoad)

	conf, err := provider.GetAppConfig()

	require.NoError(t, err)
	require.NotNil(t, conf)
	require.NotNil(t, conf.DB)

	assert.Equal(t, "localhost", conf.DB.Host)
	assert.Equal(t, 5432, conf.DB.Port)
	assert.Equal(t, "disable", conf.DB.SslMode)
	assert.Equal(t, "testuser", conf.DB.User)
	assert.Equal(t, "testpass", conf.DB.Password)
	assert.Equal(t, "testdb", conf.DB.DbName)
}

func TestDefaultConfigProvider_GetConf_WithCustomFile(t *testing.T) {
	provider := NewDefaultConfigProviderWithCustomLoader("etc/custom.yml", MockLoad)

	conf, err := provider.GetAppConfig()

	require.NoError(t, err)
	require.NotNil(t, conf)
	require.NotNil(t, conf.DB)

	assert.Equal(t, "customhost", conf.DB.Host)
	assert.Equal(t, 5433, conf.DB.Port)
	assert.Equal(t, "prefer", conf.DB.SslMode)
	assert.Equal(t, "customuser", conf.DB.User)
	assert.Equal(t, "custompass", conf.DB.Password)
	assert.Equal(t, "customdb", conf.DB.DbName)
}

func TestDefaultConfigProvider_GetConf_Error(t *testing.T) {
	provider := NewDefaultConfigProviderWithCustomLoader("etc/invalid.yml", MockLoad)

	conf, err := provider.GetAppConfig()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read config file")
	assert.Nil(t, conf)
}

func TestDefaultConfigProvider_GetConf_UnknownFile(t *testing.T) {
	provider := NewDefaultConfigProviderWithCustomLoader("etc/unknown.yml", MockLoad)

	conf, err := provider.GetAppConfig()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown config file")
	assert.Nil(t, conf)
}
