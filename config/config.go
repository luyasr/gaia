package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/gaia/errors"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type Config interface {
	Read() *config
	Watch() *config
}

type config struct {
	path string
	cfg  any
}

type Option func(*config)

func Load(path string, cfgObj any) Option {
	return func(c *config) {
		c.path = path
		c.cfg = cfgObj
	}
}

func New(opts ...Option) Config {
	options := &config{}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// Load loads the config from the given path.
func (c *config) Read() *config {
	dir, filenameOnly, extension := pathParse(c.path)
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
func (c *config) Watch() *config {
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s \n", e.Name)
		if err := viper.Unmarshal(&c.cfg); err != nil {
			handleError(err, "unmarshal conf failed")
		}
	})

	viper.WatchConfig()

	return c
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
		log.Fatal(errors.Internal(message, err.Error()))
	}
}
