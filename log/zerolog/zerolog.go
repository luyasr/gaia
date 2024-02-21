package zerolog

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	CallerDepth       = 4
	FilterCallerDepth = 5
)

var _ log.Logger = (*Logger)(nil)

// DefaultLogger default console logger
var (
	DefaultLogger = zerolog.New(console()).With().Timestamp().Logger()
)

type Logger struct {
	log zerolog.Logger
}

type Option func(*Logger)

func init() {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

// New creates a new logger with Logger
func New(logger zerolog.Logger) *Logger {
	return &Logger{
		log: logger,
	}
}

// console format the output
func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	return output
}

// rotate the log by size or time
func rotate(config Config) io.Writer {
	var writer io.Writer
	var err error
	file := path.Join(config.Filepath, config.Filename)
	switch config.Mode {
	case ModeSize:
		writer = &lumberjack.Logger{
			Filename:   file,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}
	case ModeTime:
		writer, err = rotatelogs.New(
			file+".%Y-%m-%d_%H",
			rotatelogs.WithLinkName(file),
			rotatelogs.WithMaxAge(time.Duration(config.MaxAgeDay)*24*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(config.RotationTime)*time.Hour),
		)
		if err != nil {
			log.Errorf("failed to create rotatelogs: %s", err)
			return nil
		}
	}

	return writer
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
