package config

import (
	"testing"
)

type conf struct {
	Http Http `json:"http"`
	Mode string `json:"mode" default:"debug"`
}

type Http struct {
	Host string `json:"host" default:"localhost"`
	Port int    `json:"port"`
}

func TestConfig_Load(t *testing.T) {
	cfg := new(conf)
	conf, err := New(LoadFile("config.yaml", cfg))
	if err != nil {
		t.Fatal(err)
	}
	err = conf.Read()
	if err != nil {
		t.Fatal(err)
	}
	err = conf.Watch()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg)
}
