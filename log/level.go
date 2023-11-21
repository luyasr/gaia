package log

import (
	"github.com/luyasr/gaia/errors"
	"strconv"
	"strings"
)

type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	NoLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	}

	return strconv.Itoa(int(l))
}

func LevelFieldMarshalFunc(l Level) string {
	return l.String()
}

func ParseLevel(levelStr string) (Level, error) {
	switch {
	case strings.EqualFold(levelStr, LevelFieldMarshalFunc(DebugLevel)):
		return DebugLevel, nil
	case strings.EqualFold(levelStr, LevelFieldMarshalFunc(InfoLevel)):
		return InfoLevel, nil
	case strings.EqualFold(levelStr, LevelFieldMarshalFunc(WarnLevel)):
		return WarnLevel, nil
	case strings.EqualFold(levelStr, LevelFieldMarshalFunc(ErrorLevel)):
		return ErrorLevel, nil
	case strings.EqualFold(levelStr, LevelFieldMarshalFunc(FatalLevel)):
		return FatalLevel, nil
	}

	i, err := strconv.Atoi(levelStr)
	if err != nil {
		return NoLevel, errors.Errorf(500, "Unknown level",
			"Unknown Level String: %s, defaulting to NoLevel", levelStr)
	}

	return Level(i), nil
}
