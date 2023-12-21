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

var LevelToSlog = map[Level]slog.Level{
	LevelDebug: slog.Level(-4),
	LevelInfo:  slog.Level(0),
	LevelWarn:  slog.Level(4),
	LevelError: slog.Level(8),
	LevelFatal: slog.Level(12),
}

var SlogLevelToString = map[slog.Leveler]string{
	slog.Level(-4): "DEBUG",
	slog.Level(0):  "INFO",
	slog.Level(4):  "WARN",
	slog.Level(8):  "ERROR",
	slog.Level(12): "FATAL",
}
