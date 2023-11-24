package zerolog

import (
	"dario.cat/mergo"
	"fmt"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/reflection"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log zerolog.Logger
}

type Option func(*Logger)

func init() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

// NewLogger creates a new logger with Logger
func NewLogger(logger zerolog.Logger, opts ...Option) *Logger {
	options := &Logger{
		log: logger,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

// NewConsoleLogger creates a new logger with ConsoleWriter
func NewConsoleLogger() zerolog.Logger {
	return zerolog.New(console()).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}

// NewFileLogger creates a new logger with FileWriter
func NewFileLogger(config Config) zerolog.Logger {
	var defaultConfig Config
	// use reflection to set tag
	reflection.SetUp(&defaultConfig)
	// merge the default configuration with the configuration passed in
	_ = mergo.Merge(&defaultConfig, config, mergo.WithOverride)
	writer := rotate(defaultConfig)

	return zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}

// NewMultiLogger creates a new logger with MultiWriter
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

// console format the output
func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	return output
}

// rotate the log by size
func rotate(config Config) io.Writer {
	sizeRotate := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	return sizeRotate
}

func (l *Logger) Log(level log.Level, args ...any) {
	var event *zerolog.Event

	if len(args) == 0 {
		return
	}

	if len(args)&1 == 1 {
		args = append(args, "")
	}

	switch level {
	case log.LevelDebug:
		event = l.log.Debug()
	case log.LevelInfo:
		event = l.log.Info()
	case log.LevelWarn:
		event = l.log.Warn()
	case log.LevelError:
		event = l.log.Error().Stack()
	case log.LevelFatal:
		event = l.log.Fatal().Stack()
	}

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		event = event.Any(key, args[i+1])
	}

	event.Send()
}
