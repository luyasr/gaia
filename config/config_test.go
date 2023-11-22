package config

import (
	"testing"
)

func TestConfig_Load(t *testing.T) {
	type config struct {
		Http struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		} `json:"http"`
	}

	c := new(config)
	cf := New(Load("config.yaml", c))
	cf.Read().Watch()
	t.Log(c)
}
