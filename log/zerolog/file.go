package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewFileLogger creates a new logger with FileWriter
func NewFileLogger(c Config) zerolog.Logger {
	if err := c.initConfig(); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
		return zerolog.Logger{}
	}
	writer := rotate(c)

	return zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(CallerDepth).Logger()
}
