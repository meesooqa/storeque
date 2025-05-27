package config

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	c, err := load("testdata/config.yml")

	require.NoError(t, err)

	assert.IsType(t, &SystemConfig{}, c.System)
	assert.Equal(t, "ru", c.System.DefaultLangTag)

	assert.IsType(t, &LogConfig{}, c.Log)
	assert.Equal(t, "debug", c.Log.LevelCode)
	assert.Equal(t, slog.LevelDebug, c.Log.Level)
	assert.Equal(t, "json", c.Log.OutputFormat)
	assert.Equal(t, "var/log/app.log.json", c.Log.Path)
}

func TestLoadConfigNotFoundFile(t *testing.T) {
	r, err := load("/tmp/22a03849-84de-4fa8-8ace-eb7a6777bd71.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "open /tmp/22a03849-84de-4fa8-8ace-eb7a6777bd71.txt: no such file or directory")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	r, err := load("testdata/file.txt")

	assert.Nil(t, r)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Not Yaml` into config.AppConfig")
}
