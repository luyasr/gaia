package config

import (
	"path/filepath"
	"strings"

	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/reflection"

	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
	"github.com/spf13/viper"
)

// Interface defines the methods that a config should have.
type Interface interface {
	Read() error
	Watch() error
}

// Config represents a configuration with a filepath and a target object.
type Config struct {
	filepath string
	target   any
	log      *log.Helper
	parsedPath
}

// parsedPath represents a parsed file path with directory, filename and extension.
type parsedPath struct {
	dir       string
	filename  string
	extension string
}

type Option func(*Config)

// WithLogger creates an Option that sets the logger for the config.
func WithLogger(logger log.Logger) Option {
	return func(cfg *Config) {
		cfg.log = log.NewHelper(logger)
	}
}

// LoadFile creates an Option that loads the config from the given path into the given object.
func LoadFile(filepath string, configObject any) Option {
	return func(cfg *Config) {
		cfg.filepath = filepath
		cfg.target = configObject
		cfg.parsedPath = pathParse(filepath)
	}
}

// New creates a new config with the given options.
func New(opts ...Option) (*Config, error) {
	cfg := &Config{
		log: log.NewHelper(zerolog.DefaultLogger),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if err := cfg.initConfig(); err != nil {
		return cfg, err
	}

	cfg.configureViper()

	return cfg, nil
}

func (cfg *Config) initConfig() error {
	return reflection.SetUp(cfg.target)
}

func (cfg *Config) configureViper() {
	viper.AddConfigPath(cfg.dir)
	viper.SetConfigName(cfg.filename)
	viper.SetConfigType(cfg.extension)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
}

// Read reads the config from the file and unmarshals it into the target object.
func (cfg *Config) Read() error {
	if err := viper.ReadInConfig(); err != nil {
		return errors.Internal("failed to read config file", err.Error())
	}

	if err := viper.Unmarshal(&cfg.target); err != nil {
		return errors.Internal("failed to unmarshal config file", err.Error())
	}

	return nil
}

// Watch watches the config file for changes and reloads the config when a change is detected.
func (cfg *Config) Watch() error {
	errCh := make(chan error, 1)
	viper.OnConfigChange(func(e fsnotify.Event) {
		cfg.log.Infof("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&cfg.target); err != nil {
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
