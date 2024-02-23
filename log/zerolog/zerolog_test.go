package zerolog

import (
	"encoding/json"
	"testing"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/log"
)

type Prosen struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Prosen) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

func TestNew(t *testing.T) {
	p := &Prosen{"alex", 19}
	logger := New(NewConsoleLogger())
	helper := log.NewHelper(logger)
	helper.Infof("hello world %s", p)
	str := "??"
	helper.Debugf("%s", str)
	helper.Debugw("msg", 12345, "error", errors.Internal("login failed", "incorrect account name or password").Error())
	filterLogger := New(NewConsoleLogger().With().CallerWithSkipFrameCount(FilterCallerDepth).Logger())
	filterHelper := log.NewHelper(log.NewFilter(filterLogger, log.FilterKey("password")))
	filterHelper.Error("hello world")
	filterHelper.Infow("password", "12345")

	// l, _ := NewFileLogger(Config{Mode: ModeSize})
	// fileLogger := New(l)
	// fileHelper := log.NewHelper(fileLogger)
	// fileHelper.Debug("hello world")
}
