package log

import (
	"os"
	"sync"
)

var global = new(loggerAppliance)

type loggerAppliance struct {
	mu sync.Mutex
	Logger
}

func init() {
	global.SetLogger(defaultLogger)
}

func (a *loggerAppliance) SetLogger(logger Logger) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Logger = logger
}

func Debug(args ...any) {
	global.Log(LevelDebug, args...)
}

func Info(args ...any) {
	global.Log(LevelInfo, args...)
}

func Warn(args ...any) {
	global.Log(LevelWarn, args...)
}

func Error(args ...any) {
	global.Log(LevelError, args...)
}

func Fatal(args ...any) {
	global.Log(LevelError, args...)
	os.Exit(1)
}
