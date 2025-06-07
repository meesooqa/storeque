package applog

import "log/slog"

// LoggerProvider is an interface for providing a logger
type LoggerProvider interface {
	GetLogger() (*slog.Logger, func())
}
