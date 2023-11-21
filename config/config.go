package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/gaia/errors"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

var C = new(config)

type Config interface {
	Load(string, any) *config
	Watch() *config
}

type config struct {
	cfg any
}

// Load loads the config from the given path.
func (c *config) Load(path string, cfgObj any) *config {
	c.cfg = cfgObj

	dir, filenameOnly, extension := pathParse(path)
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
