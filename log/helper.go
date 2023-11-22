package log

import "fmt"

var DefaultMessageKey = "msg"

type Option func(*Helper)

type Helper struct {
	logger  Logger
	msgKey  string
	sprint  func(a ...any) string
	sprintf func(format string, a ...any) string
}

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
