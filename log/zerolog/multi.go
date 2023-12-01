package zerolog

import (
	"dario.cat/mergo"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/reflection"
	"github.com/rs/zerolog"
)

// NewMultiLogger creates a new logger with MultiLevelWriter
func NewMultiLogger(config Config) zerolog.Logger {
	var defaultConfig Config
	// use reflection to set tag
	reflection.SetUp(&defaultConfig)
	// merge the default configuration with the configuration passed in
	_ = mergo.Merge(&defaultConfig, config, mergo.WithOverride)
	writer := rotate(defaultConfig)
	multi := zerolog.MultiLevelWriter(console(), writer)

	return zerolog.New(multi).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}
