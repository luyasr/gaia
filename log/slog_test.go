package log

import (
	"log/slog"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func (u *User) LogValue() slog.Value {
	return slog.StringValue(u.Name)
}

func TestNewSlog(t *testing.T) {
	user := &User{"alex", 18}
	Info("msg", "user", user)

	Debug("", "user", "user")
	Debugw("", "user", user)
	Debugf("%v", user)
	Info("2")
	Warn("3", "username", "lisi")
}