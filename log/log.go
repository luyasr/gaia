package log

import (
	"log/slog"
	"os"
)

const (
	DefaultCaller       = 4
	DefaultFilterCaller = 5
)

var defaultLogger = NewSlog(slog.New(slog.NewTextHandler(os.Stdout, HandlerOptions())))

type Logger interface {
	Log(level Level, args ...any)
}
