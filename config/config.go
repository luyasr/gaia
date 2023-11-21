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
	Load(string) *config
	Watch()
	Merge(any) error
}

type config struct {
}

func pathParse(path string) (string, string, string) {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	fileType := strings.TrimPrefix(extension, ".")
	filenameOnly := strings.TrimSuffix(filename, extension)

	return dir, filenameOnly, fileType
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatal(errors.Internal(message, err.Error()))
	}
}

func (c *config) Load(path string) *config {
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

	if err := viper.Unmarshal(C); err != nil {
		handleError(err, "unmarshal conf failed")
	}

	return C
}

func (c *config) Watch() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s \n", e.Name)
		if err := viper.Unmarshal(C); err != nil {
			handleError(err, "unmarshal conf failed")
		}
	})

	viper.WatchConfig()
}

func (c *config) Merge(cfg any) error {
	return nil
}
