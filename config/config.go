package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config interface {
	Read() *config
	Watch()
}

type config struct {
	// filepath is the path of the config file.
	filepath string
	// cfg is the config object.
	cfg any
}

type Option func(*config)

// LoadFile loads the config from the given path.
func LoadFile(filepath string, cfgObj any) Option {
	return func(c *config) {
		c.filepath = filepath
		c.cfg = cfgObj
	}
}

// New creates a new config.
func New(opts ...Option) Config {
	options := &config{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// Load loads the config from the given path.
func (c *config) Read() *config {
	dir, filenameOnly, extension := pathParse(c.filepath)
	viper.AddConfigPath(dir)
	viper.SetConfigName(filenameOnly)
	viper.SetConfigType(extension)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		handleError(err, "read conf failed")
	}

	if err := viper.Unmarshal(&c.cfg); err != nil {
		handleError(err, "unmarshal conf failed")
	}

	return c
}

// Watch watches the config file.
func (c *config) Watch() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&c.cfg); err != nil {
			handleError(err, "unmarshal conf failed")
		}
	})

	viper.WatchConfig()
}

// pathParse parses the path and returns the dir, filename and extension.
func pathParse(path string) (string, string, string) {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	fileType := strings.TrimPrefix(extension, ".")
	filenameOnly := strings.TrimSuffix(filename, extension)

	return dir, filenameOnly, fileType
}

// handleError handles the error.
func handleError(err error, message string) {
	if err != nil {
		log.Fatal(errors.Internal(message, err.Error()).Error())
	}
}
