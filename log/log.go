package log

import (
	"log/slog"
	"os"
	"runtime"
	"sync"
)

// defaultLogger is a singleton instance of Slog.
var (
	defaultLogger Logger
	once          sync.Once
)

func getDefaultLogger() Logger {
	once.Do(func() {
		defaultLogger = NewSlog(slog.New(slog.NewTextHandler(os.Stdout, HandlerOptions())))
	})
	return defaultLogger
}

// Logger is an interface for logging.
type Logger interface {
	Log(level Level, args ...any)
}

func CallerDepth(skip int) int {
	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	// skip = 0 means runtime.Callers; skip = 1 means the caller of runtime.Callers; etc.
	pc := make([]uintptr, 10)
	n := runtime.Callers(skip + 1, pc)
	if n == 0 {
		return 0
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	frameCount := 0
	for {
		_, more := frames.Next()
		if !more {
			break
		}
		frameCount++
	}
	return frameCount
}
