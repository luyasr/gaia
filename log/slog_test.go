package log

import (
	"errors"
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
	Debugf("xxx%vxxx", user)
	Info("2")
	Warn("3", "username", "lisi")
	Fatal("5", "error", errors.New("conn failed"))
	Error("4", "error", errors.New("login failed"))
}
