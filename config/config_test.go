package config

import (
	"testing"
	"time"
)

type conf struct {
	Http Http `json:"http"`
}

type Http struct {
	Host string `json:"host" default:"localhost"`
	Port int    `json:"port" default:"8080"`
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
	time.Sleep(30 * time.Second)
}
