package config

import (
	"testing"
)

type cfg struct {
	Http Http `json:"http"`
}

type Http struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func TestConfig_Load(t *testing.T) {
	c := new(cfg)
	cf := New(LoadFile("config.yaml", c))
	cf.Read().Watch()
	t.Log(c)
}
