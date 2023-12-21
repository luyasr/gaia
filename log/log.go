package log

import (
	"log/slog"
	"os"
	"sync"
)

const (
	DefaultCaller       = 4
	DefaultFilterCaller = 5
)

// defaultLogger is a singleton instance of Slog.
var (
	defaultLogger Logger
	once          sync.Once
)

func getDefaultLogger() Logger {
	once.Do(func() {
		defaultLogger = NewSlog(slog.New(slog.NewTextHandler(os.Stdout, HandlerOptions())))
	})
	return defaultLogger
}

// Logger is an interface for logging.
type Logger interface {
	Log(level Level, args ...any)
}
