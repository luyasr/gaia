package config

import (
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/reflection"
	"github.com/spf13/viper"
)

// Interface defines the methods that a config should have.
type Interface interface {
	Read() (*Config, error)
	Watch() error
}

// Config represents a configuration with a filepath and a target object.
type Config struct {
	filepath string
	target   any
}

// parsedPath represents a parsed file path with directory, filename and extension.
type parsedPath struct {
	dir       string
	filename  string
	extension string
}

type Option func(*Config)

// LoadFile creates an Option that loads the config from the given path into the given object.
func LoadFile(filepath string, configObject any) Option {
	return func(c *Config) {
		c.filepath = filepath
		c.target = configObject
	}
}

// New creates a new config with the given options.
func New(opts ...Option) *Config {
	c := &Config{}

	for _, opt := range opts {
		opt(c)
	}

	_ = c.initConfig()

	return c
}

func (c *Config) initConfig() error {
	if err := reflection.SetUp(c.target); err != nil {
		return err
	}

	p := pathParse(c.filepath)
	viper.AddConfigPath(p.dir)
	viper.SetConfigName(p.filename)
	viper.SetConfigType(p.extension)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	return nil
}

// Read reads the config from the file and unmarshals it into the target object.
func (c *Config) Read() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Internal("failed to read config file", err.Error())
	}

	if err := viper.Unmarshal(&c.target); err != nil {
		return nil, errors.Internal("failed to unmarshal config file", err.Error())
	}

	return c, nil
}

// Watch watches the config file for changes and reloads the config when a change is detected.
func (c *Config) Watch() error {
	errCh := make(chan error, 1)
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&c.target); err != nil {
			errCh <- err
		}
	})

	viper.WatchConfig()

	select {
	case err := <-errCh:
		return errors.Internal("failed to unmarshal config file", err.Error())
	default:
		return nil
	}
}

// pathParse parses the path and returns a parsedPath struct with the directory, filename and extension.
func pathParse(path string) parsedPath {
	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	fileType := strings.TrimPrefix(extension, ".")
	filenameOnly := strings.TrimSuffix(filename, extension)

	return parsedPath{
		dir:       dir,
		filename:  filenameOnly,
		extension: fileType,
	}
}
