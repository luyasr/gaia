package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewConsoleLogger creates a new logger with ConsoleWriter
func NewConsoleLogger() zerolog.Logger {
	return zerolog.New(console()).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}
