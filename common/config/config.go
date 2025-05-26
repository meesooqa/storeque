package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

// AppConfig from config yml
type AppConfig struct {
	//System *SystemConfig `yaml:"system"`
	Log *LogConfig `yaml:"log"`
	DB  *DbConfig  `yaml:"db"`
}

// SystemConfig - system parameters
// type SystemConfig struct {}

// LogConfig - log parameters
type LogConfig struct {
	LevelCode    string `yaml:"level_code"`
	Level        slog.Level
	OutputFormat string `yaml:"output_format"`
	Path         string `yaml:"path"`
}

// DbConfig - DB parameters
type DbConfig struct {
	IsDebugMode bool   `yaml:"is_debug_mode"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	SslMode     string `yaml:"sslmode"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DbName      string `yaml:"dbname"`
}

// load config from file
func load(fname string) (res *AppConfig, err error) {
	res = &AppConfig{}
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, res); err != nil {
		return nil, err
	}

	level, ok := logLevelMap[res.Log.LevelCode]
	if ok {
		res.Log.Level = level
	} else {
		res.Log.Level = slog.LevelInfo
	}

	return res, nil
}
