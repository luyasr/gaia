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

func NewLogger(logger zerolog.Logger, opts ...Option) *Logger {
	options := &Logger{
		log: logger,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func NewConsoleLogger() zerolog.Logger {
	return zerolog.New(console()).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}

func NewFileLogger(config Config) zerolog.Logger {
	var defaultConfig Config
	// use reflection to set tag
	reflection.SetUp(&defaultConfig)
	// merge the default configuration with the configuration passed in
	_ = mergo.Merge(&defaultConfig, config, mergo.WithOverride)
	writer := rotate(defaultConfig)

	return zerolog.New(writer).With().Timestamp().CallerWithSkipFrameCount(log.DefaultCaller).Logger()
}

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

func (l *Logger) Log(level log.Level, kvs ...any) error {
	var event *zerolog.Event
	if len(kvs) == 0 {
		return nil
	}

	if len(kvs)&1 == 1 {
		kvs = append(kvs, "")
	}

	switch level {
	case log.DebugLevel:
		event = l.log.Debug()
	case log.InfoLevel:
		event = l.log.Info()
	case log.WarnLevel:
		event = l.log.Warn()
	case log.ErrorLevel:
		event = l.log.Error()
	case log.FatalLevel:
		event = l.log.Fatal()
	default:
		event = l.log.Info()
	}

	for i := 0; i < len(kvs); i += 2 {
		key, ok := kvs[i].(string)
		if !ok {
			continue
		}
		event = event.Any(key, kvs[i+1])
	}

	event.Send()
	return nil
}
