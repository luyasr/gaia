package mysql

import (
	"fmt"

	"github.com/luyasr/gaia/reflection"
)

type Config struct {
	Host            string `json:"host" default:"localhost"`
	Port            int    `json:"port" default:"3306"`
	Username        string `json:"username" default:"root"`
	Password        string `json:"password"`
	DataBase        string `json:"database"`
	Charset         string `json:"charset" default:"utf8mb4"`
	ParseTime       *bool  `json:"parse_time" default:"true"`
	Loc             string `json:"loc" default:"Local"`
	Timeout         int    `json:"timeout" default:"10"`
	MaxIdleConns    int    `json:"max_idle_conns" default:"10"`
	MaxOpenConns    int    `json:"max_open_conns" default:"100"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" default:"3600"`
	LogLevel        string `json:"log_level" default:"silent"`
}

func (c *Config) address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) dsn() string {
	parseTime := "False"
	if c.ParseTime != nil && *c.ParseTime {
		parseTime = "True"
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%ds",
		c.Username,
		c.Password,
		c.address(),
		c.DataBase,
		c.Charset,
		parseTime,
		c.Loc,
		c.Timeout,
	)
}

func (c *Config) logLevel() int {
	switch c.LogLevel {
	case "silent":
		return 1
	case "error":
		return 2
	case "warn":
		return 3
	case "info":
		return 4
	default:
		return 1
	}
}

func (c *Config) initConfig() error {
	return reflection.SetUp(c)
}
