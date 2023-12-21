package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewMultiLogger creates a new logger with MultiLevelWriter
func NewMultiLogger(config Config) zerolog.Logger {
	writer := rotate(getDefaultConfig(config))
	multi := zerolog.MultiLevelWriter(console(), writer)

	return zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}
