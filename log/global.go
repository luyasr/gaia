package log

import (
	"fmt"
	"os"
	"sync"
)

var global = new(loggerAppliance)

type loggerAppliance struct {
	mu sync.Mutex
	Logger
}

func init() {
	global.SetLogger(getDefaultLogger())
}

func (a *loggerAppliance) SetLogger(logger Logger) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Logger = logger
}

func Debug(msg string, args ...any) {
	global.Log(LevelDebug, append([]any{msg}, args...)...)
}

func Debugf(format string, args ...any) {
	global.Log(LevelDebug, fmt.Sprintf(format, args...))
}

func Debugw(args ...any) {
	global.Log(LevelDebug, args...)
}

func Info(msg string, args ...any) {
	global.Log(LevelInfo, append([]any{msg}, args...)...)
}

func Infof(format string, args ...any) {
	global.Log(LevelInfo, fmt.Sprintf(format, args...))
}

func Infow(args ...any) {
	global.Log(LevelInfo, args...)
}

func Warn(msg string, args ...any) {
	global.Log(LevelWarn, append([]any{msg}, args...)...)
}

func Warnf(format string, args ...any) {
	global.Log(LevelWarn, fmt.Sprintf(format, args...))
}

func Warnw(args ...any) {
	global.Log(LevelWarn, args...)
}

func Error(msg string, args ...any) {
	global.Log(LevelError, append([]any{msg}, args...)...)
}

func Errorf(format string, args ...any) {
	global.Log(LevelError, fmt.Sprintf(format, args...))
}

func Errorw(args ...any) {
	global.Log(LevelError, args...)
}

func Fatal(msg string, args ...any) {
	global.Log(LevelFatal, append([]any{msg}, args...)...)
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	global.Log(LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Fatalw(args ...any) {
	global.Log(LevelFatal, args...)
	os.Exit(1)
}
