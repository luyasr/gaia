package log

const (
	DefaultCaller       = 4
	DefaultFilterCaller = 5
)

type Logger interface {
	Log(level Level, keyValues ...any) error
}
