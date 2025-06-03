package common_log

import (
	"io"
	"log"
	"log/slog"
	"os"

	"tg-star-shop-bot-001/common/config"
)

type ConsoleLoggerProvider struct {
	conf *config.LogConfig
}

func NewConsoleLoggerProvider(conf *config.LogConfig) *ConsoleLoggerProvider {
	return &ConsoleLoggerProvider{
		conf: conf,
	}
}

// GetLogger returns logger writing to Stdout
func (o *ConsoleLoggerProvider) GetLogger() (*slog.Logger, func()) {
	noop := func() {
		// skip any actions
	}
	return slog.New(getLogHandler(o.conf, os.Stdout, &slog.HandlerOptions{Level: o.conf.Level})), noop
}

type FileLoggerProvider struct {
	conf *config.LogConfig
}

func NewFileLoggerProvider(conf *config.LogConfig) *FileLoggerProvider {
	return &FileLoggerProvider{
		conf: conf,
	}
}

// GetLogger returns logger writing to file and cleanup func.
// logger, cleanup := loggerProvider.GetLogger()
// defer cleanup()
func (o *FileLoggerProvider) GetLogger() (*slog.Logger, func()) {
	file, err := os.OpenFile(o.conf.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	logger := slog.New(getLogHandler(o.conf, file, &slog.HandlerOptions{Level: o.conf.Level}))
	cleanup := func() {
		file.Close()
	}
	return logger, cleanup
}

func getLogHandler(conf *config.LogConfig, w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch conf.OutputFormat {
	case "json":
		return slog.NewJSONHandler(w, opts)
	default:
		return slog.NewTextHandler(w, opts)
	}
}
