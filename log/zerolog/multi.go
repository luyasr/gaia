package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewMultiLogger creates a new logger with MultiLevelWriter
func NewMultiLogger(c Config) (zerolog.Logger, error) {
	if err := c.initConfig(); err != nil {
		return zerolog.Logger{}, err
	}
	writer := rotate(c)
	multi := zerolog.MultiLevelWriter(console(), writer)

	return zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger(), nil
}
