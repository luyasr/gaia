package config

import (
	"testing"
)

type conf struct {
	Http Http `json:"http"`
}

type Http struct {
	Host string `json:"host" default:"localhost"`
	Port int    `json:"port" default:"8080"`
}

func TestConfig_Load(t *testing.T) {
	c := new(conf)
	cf, err := New(LoadFile("config.yaml", c)).Read()
	if err != nil {
		t.Fatal(err)
	}
	cf.Watch()
	t.Log(c)
}
