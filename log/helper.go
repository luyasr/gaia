package log

import (
	"fmt"
	"os"
)

// DefaultMessageKey is the default key for message field.
var DefaultMessageKey = "msg"

// Helper is a wrapper of Logger.
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

func (h *Helper) Log(level Level, keyValues ...any) error {
	return h.logger.Log(level, keyValues...)
}

func (h *Helper) Debug(a ...any) {
	_ = h.logger.Log(DebugLevel, h.msgKey, h.sprint(a...))
}

func (h *Helper) Debugf(format string, a ...any) {
	_ = h.logger.Log(DebugLevel, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Debugw(keyValues ...any) {
	_ = h.logger.Log(DebugLevel, keyValues...)
}

func (h *Helper) Info(a ...any) {
	_ = h.logger.Log(InfoLevel, h.msgKey, h.sprint(a...))
}

func (h *Helper) Infof(format string, a ...any) {
	_ = h.logger.Log(InfoLevel, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Infow(keyValues ...any) {
	_ = h.logger.Log(InfoLevel, keyValues...)
}

func (h *Helper) Warn(a ...any) {
	_ = h.logger.Log(WarnLevel, h.msgKey, h.sprint(a...))
}

func (h *Helper) Warnf(format string, a ...any) {
	_ = h.logger.Log(WarnLevel, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Warnw(keyValues ...any) {
	_ = h.logger.Log(WarnLevel, keyValues...)
}

func (h *Helper) Error(a ...any) {
	_ = h.logger.Log(ErrorLevel, h.msgKey, h.sprint(a...))
}

func (h *Helper) Errorf(format string, a ...any) {
	_ = h.logger.Log(ErrorLevel, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Errorw(keyValues ...any) {
	_ = h.logger.Log(ErrorLevel, keyValues...)
}

func (h *Helper) Fatal(a ...any) {
	_ = h.logger.Log(FatalLevel, h.msgKey, h.sprint(a...))
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...any) {
	_ = h.logger.Log(FatalLevel, h.msgKey, h.sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(keyValues ...any) {
	_ = h.logger.Log(FatalLevel, keyValues...)
	os.Exit(1)
}
