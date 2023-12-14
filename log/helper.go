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
func NewHelper(logger Logger, opt ...Option) *Helper {
	options := &Helper{
		logger:  logger,
		msgKey:  DefaultMessageKey,
		sprint:  fmt.Sprint,
		sprintf: fmt.Sprintf,
	}

	for _, o := range opt {
		o(options)
	}

	return options
}

func (h *Helper) Log(level Level, a ...any) {
	h.logger.Log(level, a...)
}

func (h *Helper) Debug(a ...any) {
	h.logger.Log(LevelDebug, h.msgKey, h.sprint(a...))
}

func (h *Helper) Debugf(format string, a ...any) {
	h.logger.Log(LevelDebug, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Debugw(a ...any) {
	h.logger.Log(LevelDebug, a...)
}

func (h *Helper) Info(a ...any) {
	h.logger.Log(LevelInfo, h.msgKey, h.sprint(a...))
}

func (h *Helper) Infof(format string, a ...any) {
	h.logger.Log(LevelInfo, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Infow(a ...any) {
	h.logger.Log(LevelInfo, a...)
}

func (h *Helper) Warn(a ...any) {
	h.logger.Log(LevelWarn, h.msgKey, h.sprint(a...))
}

func (h *Helper) Warnf(format string, a ...any) {
	h.logger.Log(LevelWarn, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Warnw(a ...any) {
	h.logger.Log(LevelWarn, a...)
}

func (h *Helper) Error(a ...any) {
	h.logger.Log(LevelError, h.msgKey, h.sprint(a...))
}

func (h *Helper) Errorf(format string, a ...any) {
	h.logger.Log(LevelError, h.msgKey, h.sprintf(format, a...))
}

func (h *Helper) Errorw(a ...any) {
	h.logger.Log(LevelError, a...)
}

func (h *Helper) Fatal(a ...any) {
	h.logger.Log(LevelFatal, h.msgKey, h.sprint(a...))
	os.Exit(1)
}

func (h *Helper) Fatalf(format string, a ...any) {
	h.logger.Log(LevelFatal, h.msgKey, h.sprintf(format, a...))
	os.Exit(1)
}

func (h *Helper) Fatalw(a ...any) {
	h.logger.Log(LevelFatal, a...)
	os.Exit(1)
}
