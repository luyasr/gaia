package zerolog

import (
	"dario.cat/mergo"
	"github.com/luyasr/gaia/reflection"
)

// Mode defines the log cutting mode.
// 0: size
// 1: time
type Mode int

const (
	ModeSize Mode = iota
	ModeTime
)

// Config defines the configuration for the logger.
type Config struct {
	Mode         Mode
	Filepath     string `default:"."`
	Filename     string `default:"app.log"`
	MaxSize      int    `default:"100"`
	MaxBackups   int    `default:"5"`
	MaxAge       int    `default:"30"`
	Compress     bool   `default:"false"`
	MaxAgeDay    int    `default:"7"`
	RotationTime int    `default:"1"`
}

func getDefaultConfig(config Config) Config {
	var defaultConfig Config
	// use reflection to set tag
	reflection.SetUp(&defaultConfig)
	// merge the default configuration with the configuration passed in
	_ = mergo.Merge(&defaultConfig, config, mergo.WithOverride)

	return defaultConfig
}
