package zerolog

import (
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
	"testing"
)

func TestNew(t *testing.T) {
	logger := New(ConsoleLogger)
	helper := log.NewHelper(logger)
	helper.Debug("")
	str := "??"
	helper.Debugf("%s", str)
	helper.Debugw("msg", 12345, "error", errors.Internal("login failed", "incorrect account name or password").Error())

	filterLogger := New(NewConsoleLogger().With().CallerWithSkipFrameCount(log.DefaultFilterCaller).Logger())
	filterHelper := log.NewHelper(log.NewFilter(filterLogger, log.FilterKey("password")))
	filterHelper.Error("hello world")
	filterHelper.Infow("password", "12345")

	New(NewFileLogger(Config{Filename: "acc.log"}))

	ConsoleLogger.Info().Str("error", "error").Send()
}
