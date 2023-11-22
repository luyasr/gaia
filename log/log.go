package log

type Logger interface {
	Log(level Level, keyValues ...any) error
}
