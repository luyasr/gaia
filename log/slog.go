package log

import (
	"context"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
)

const (
	defaultSlogCaller = 3
)

var _ Logger = (*Slog)(nil)

type Slog struct {
	logger *slog.Logger
}

type SlogOption func(*Slog)

func NewSlog(logger *slog.Logger, opts ...SlogOption) *Slog {
	options := &Slog{
		logger: logger,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func HandlerOptions() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: SlogLevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := SlogLevels[level]
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

func (s *Slog) Log(level Level, args ...any) {
	switch level {
	case LevelDebug:
		s.logger.With("caller", caller(defaultSlogCaller)).Debug(args[0].(string), args[1:]...)
	case LevelInfo:
		s.logger.With("caller", caller(defaultSlogCaller)).Info(args[0].(string), args[1:]...)
	case LevelWarn:
		s.logger.With("caller", caller(defaultSlogCaller)).Warn(args[0].(string), args[1:]...)
	case LevelError:
		s.logger.With("caller", caller(defaultSlogCaller)).Error(args[0].(string), args[1:]...)
	case LevelFatal:
		s.logger.With("caller", caller(defaultSlogCaller)).Log(context.Background(), SlogLevelFatal, args[0].(string), args[1:]...)
	}
}
