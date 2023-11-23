package log

import (
	"fmt"
	"os"
)

// DefaultMessageKey is the default key for message field.
var DefaultMessageKey = "msg"

// Option Helper is a wrapper of Logger.
type Option func(*Helper)

type Helper struct {
	logger  Logger
	msgKey  string
	sprint  func(a ...any) string
	sprintf func(format string, a ...any) string
}

func WithMessageKey(key string) Option {
	return func(h *Helper) {
		h.msgKey = key
	}
}

func WithSprint(sprint func(a ...any) string) Option {
	return func(h *Helper) {
		h.sprint = sprint
	}
}

func WithSprintf(sprintf func(format string, a ...any) string) Option {
	return func(h *Helper) {
		h.sprintf = sprintf
	}
}

// NewHelper creates a new Helper.
func NewHelper(logger Logger, opts ...Option) *Helper {
	options := &Helper{
		logger:  logger,
		msgKey:  DefaultMessageKey,
		sprint:  fmt.Sprint,
		sprintf: fmt.Sprintf,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func (h *Helper) Log(level Level, args ...any) {
	h.logger.Log(level, args...)
}

func (h *Helper) Debug(a ...any) {
	h.logger.Log(LevelDebug, h.msgKey, h.sprint(a...))
}

func (h *Helper) Debugf(format string, a ...any) {
	h.logger.Log(LevelDebug, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Debugw(args ...any) {
	h.logger.Log(LevelDebug, args...)
}

func (h *Helper) Info(a ...any) {
	h.logger.Log(LevelInfo, h.msgKey, h.sprint(a...))
}

func (h *Helper) Infof(format string, a ...any) {
	h.logger.Log(LevelInfo, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Infow(args ...any) {
	h.logger.Log(LevelInfo, args...)
}

func (h *Helper) Warn(a ...any) {
	h.logger.Log(LevelWarn, h.msgKey, h.sprint(a...))
}

func (h *Helper) Warnf(format string, a ...any) {
	h.logger.Log(LevelWarn, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Warnw(args ...any) {
	h.logger.Log(LevelWarn, args...)
}

func (h *Helper) Error(a ...any) {
	h.logger.Log(LevelError, h.msgKey, h.sprint(a...))
}

func (h *Helper) Errorf(format string, a ...any) {
	h.logger.Log(LevelError, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Errorw(args ...any) {
	h.logger.Log(LevelError, args...)
}

func (h *Helper) Fatal(a ...any) {
	h.logger.Log(LevelFatal, h.msgKey, h.sprint(a...))
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...any) {
	h.logger.Log(LevelFatal, h.msgKey, h.sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(args ...any) {
	h.logger.Log(LevelFatal, args...)
	os.Exit(1)
}
