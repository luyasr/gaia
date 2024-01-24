package zerolog

import (
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
)

// NewFileLogger creates a new logger with FileWriter
func NewFileLogger(c Config) (zerolog.Logger, error) {
	if err := c.initConfig(); err != nil {
		return zerolog.Logger{}, err
	}
	writer := rotate(c)

	return zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger(), nil
}
