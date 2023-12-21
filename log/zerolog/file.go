package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewFileLogger creates a new logger with FileWriter
func NewFileLogger(config Config) zerolog.Logger {
	writer := rotate(getDefaultConfig(config))

	return zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}
