package log

import (
	"errors"
	"testing"
)

func TestNewSlog(t *testing.T) {
	Debug("1")
	Info("2")
	Warn("3", "username", "lisi")
	Fatal("5", "error", errors.New("conn failed"))
	Error("4", "error", errors.New("login failed"))
}
