package zerolog

import (
	"errors"
	"github.com/luyasr/gaia/log"
	"testing"
)

func TestNewLogger(t *testing.T) {
	zLogger := NewConsoleLogger()
	logger := NewLogger(&zLogger)
	helper := log.NewHelper(logger)
	helper.Debug(1)
	str := "??"
	helper.Debugf("%s", str)
	helper.Debugw("msg", 1, "error", errors.New("1"))
}
