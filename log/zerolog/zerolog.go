package zerolog

import (
	"fmt"
	"github.com/luyasr/gaia/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"strings"
	"time"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *zerolog.Logger
}

func NewLogger(logger *zerolog.Logger) *Logger {
	return &Logger{
		log: logger,
	}
}

func NewConsoleLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return zerolog.New(console()).With().Timestamp().Logger()
}

func console() zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
	}

	return output
}

func (l *Logger) Log(level log.Level, keyValues ...any) error {
	var event *zerolog.Event
	if len(keyValues) == 0 {
		return nil
	}

	if len(keyValues)&1 == 1 {
		keyValues = append(keyValues, "")
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

	for i := 0; i < len(keyValues); i += 2 {
		key, ok := keyValues[i].(string)
		if !ok {
			continue
		}
		event = event.Any(key, keyValues[i+1])
	}

	event.Send()
	return nil
}
