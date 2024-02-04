package sqlite

import (
	"sync"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const name = "sqlite"

var once sync.Once

type Sqlite struct {
	Client *gorm.DB
}

type Option func(*Sqlite)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Sqlite{})
}

func (s *Sqlite) Init() error {
	cfg, ok := ioc.Container.GetFieldValueByConfig("Sqlite")
	if !ok {
		return nil
	}

	sqliteCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("sqlite", "Sqlite type assertion failed, expected *Config, got %T", cfg)
	}

	sqlite, err := New(sqliteCfg)
	if err != nil {
		return err
	}
	s.Client = sqlite.Client

	return nil
}

func (s *Sqlite) Name() string {
	return name
}

func New(c *Config, opts ...Option) (*Sqlite, error) {
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	s := &Sqlite{}

	for _, opt := range opts {
		opt(s)
	}

	s, err := new(c, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func new(c *Config, s *Sqlite) (*Sqlite, error) {
	var err error

	once.Do(func() {
		s.Client, err = gorm.Open(sqlite.Open(c.Path), &gorm.Config{})
	})

	return s, err
}

func (s *Sqlite) Close() error {
	db, err := s.Client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
