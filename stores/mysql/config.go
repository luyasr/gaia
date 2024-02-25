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
	ParseTime       *bool  `json:"parseTime" default:"true"`
	Loc             string `json:"loc" default:"Local"`
	Timeout         int    `json:"timeout" default:"10"`
	MaxIdleConns    int    `json:"maxIdleConns" default:"10"`
	MaxOpenConns    int    `json:"maxOpenConns" default:"100"`
	ConnMaxLifetime int    `jsin:"connMaxLifetime" default:"3600"`
	LogLevel        string `json:"logLevel" default:"silent"`
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

func (c *Config) initConfig() (*Config, error) {
	if c == nil {
		c = &Config{}
	}

	if err := reflection.SetUp(c); err != nil {
		return nil, err
	}

	return c, nil
}
