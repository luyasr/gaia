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
	C.Load("../config.yaml", c)

	t.Log(c)
}
