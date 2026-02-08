package logging

import (
	"github.com/go-hexagonal-practice/internal/adapters/config"
	"github.com/go-hexagonal-practice/internal/core/ports"
)

func NewLogger(cfg *config.LoggingConfig) ports.Logger {
	switch cfg.Adapter() {
	case "loki":
		return NewLokiLogger(
			cfg.LokiURL(),
			cfg.LokiLabels(),
		)
	case "multi":
		stdoutLogger := NewStdoutLogger()
		lokiLogger := NewLokiLogger(
			cfg.LokiURL(),
			cfg.LokiLabels(),
		)
		loggers := []ports.Logger{stdoutLogger, lokiLogger}
		return NewMultiLogger(loggers...)
	default:
		return NewStdoutLogger()
	}
}
