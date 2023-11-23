package log

import "log/slog"

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	SlogLevelDebug = slog.Level(-4)
	SlogLevelInfo  = slog.Level(0)
	SlogLevelWarn  = slog.Level(4)
	SlogLevelError = slog.Level(8)
	SlogLevelFatal = slog.Level(12)
)

var SlogLevels = map[slog.Leveler]string{
	SlogLevelDebug: "DEBUG",
	SlogLevelInfo:  "INFO",
	SlogLevelWarn:  "WARN",
	SlogLevelError: "ERROR",
	SlogLevelFatal: "FATAL",
}
