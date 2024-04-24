package log

import (
	"context"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var _ Logger = (*Slog)(nil)

type Slog struct {
	logger *slog.Logger
}

type SlogOption func(*Slog)

func NewSlog(logger *slog.Logger, opts ...SlogOption) *Slog {
	slog := &Slog{
		logger: logger,
	}

	for _, opt := range opts {
		opt(slog)
	}

	return slog
}

func HandlerOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: LevelToSlog[LevelDebug],
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format(time.DateTime))
			}

			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := SlogLevelToString[level]
				if !exists {
					levelLabel = level.String()
				}

				a.Value = slog.StringValue(levelLabel)
			}

			return a
		},
	}
}

// Caller returns the caller of the function that called it.
func caller(depth int) string {
	_, file, line, _ := runtime.Caller(depth)
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
	idx = strings.LastIndexByte(file[:idx], '/')
	return file[idx+1:] + ":" + strconv.Itoa(line)
}

func (s *Slog) log(level Level, msg string, args ...any) {
	callerDepth := CallerDepth(2)
	switch level {
	case LevelDebug:
		s.logger.With("caller", caller(callerDepth)).Debug(msg, args...)
	case LevelInfo:
		s.logger.With("caller", caller(callerDepth)).Info(msg, args...)
	case LevelWarn:
		s.logger.With("caller", caller(callerDepth)).Warn(msg, args...)
	case LevelError:
		s.logger.With("caller", caller(callerDepth)).Error(msg, args...)
	case LevelFatal:
		s.logger.With("caller", caller(callerDepth)).Log(context.Background(), LevelToSlog[LevelFatal], msg, args...)
	}
}

func (s *Slog) Log(level Level, args ...any) {
	if len(args) == 0 {
		return
	}

	msg, ok := args[0].(string)
	if !ok {
		s.logger.Error("First argument to Log must be a string")
		return
	}

	args = args[1:]
	s.log(level, msg, args...)
}
