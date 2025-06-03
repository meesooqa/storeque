package common_log

import "log/slog"

type LoggerProvider interface {
	GetLogger() (*slog.Logger, func())
}
