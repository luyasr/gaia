package log

import (
	"context"
	"log/slog"
	"time"
)

var _ Logger = (*Slog)(nil)

const (
	defaultCallerDepth = 4
)

type Slog struct {
	logger            *slog.Logger
	callerDepth       int
	filterCallerDepth int
}

type SlogOption func(*Slog)

func WithCallerDepth(depth int) SlogOption {
	return func(s *Slog) {
		s.callerDepth = depth
	}
}

func WithFilterCallerDepth(depth int) SlogOption {
	return func(s *Slog) {
		s.filterCallerDepth = depth
	}
}

func NewSlog(logger *slog.Logger, opts ...SlogOption) *Slog {
	slog := &Slog{
		logger:      logger,
		callerDepth: defaultCallerDepth,
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

func (s *Slog) log(level Level, msg string, args ...any) {
	switch level {
	case LevelDebug:
		s.logger.With("caller", Caller(s.callerDepth)).Debug(msg, args...)
	case LevelInfo:
		s.logger.With("caller", Caller(s.callerDepth)).Info(msg, args...)
	case LevelWarn:
		s.logger.With("caller", Caller(s.callerDepth)).Warn(msg, args...)
	case LevelError:
		s.logger.With("caller", Caller(s.callerDepth)).Error(msg, args...)
	case LevelFatal:
		s.logger.With("caller", Caller(s.callerDepth)).Log(context.Background(), LevelToSlog[LevelFatal], msg, args...)
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
