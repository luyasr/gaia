package log

import "log/slog"

const (
	DefaultCaller       = 4
	DefaultFilterCaller = 5
)

var defaultLogger = NewSlog(slog.Default())

//var defaultLogger = NewSlog(slog.New(slog.NewTextHandler(os.Stdout, HandlerOptions())))

type Logger interface {
	Log(level Level, args ...any)
}
