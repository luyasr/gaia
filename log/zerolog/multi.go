package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewMultiLogger creates a new logger with MultiLevelWriter
func NewMultiLogger(c Config) zerolog.Logger {
	if err := c.initConfig(); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
		return zerolog.Logger{}
	}
	writer := rotate(c)
	multi := zerolog.MultiLevelWriter(console(), writer)

	return zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(CallerDepth).Logger()
}
