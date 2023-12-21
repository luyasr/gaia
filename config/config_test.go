package config

import (
	"sync"
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
	c := new(conf)
	cf := New(LoadFile("config.yaml", c))
	cf.Read().Watch()
	t.Log(c)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Second)
	}()

	wg.Wait()
	t.Log(c)
}
