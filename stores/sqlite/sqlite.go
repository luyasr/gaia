package sqlite

import (
	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Name = "sqlite"

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
		return errors.Internal("sqlite config not found", "expected *Config, got %T", cfg)
	}

	sqliteCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("sqlite config type assertion failed", "expected *Config, got %T", cfg)
	}

	sqlite, err := New(sqliteCfg)
	if err != nil {
		return err
	}
	s.Client = sqlite.Client

	return nil
}

func (s *Sqlite) Name() string {
	return Name
}

func New(c *Config, opts ...Option) (*Sqlite, error) {
	cfg, err := c.initConfig()
	if err != nil {
		return nil, err
	}

	s := &Sqlite{}

	for _, opt := range opts {
		opt(s)
	}

	return new(cfg, s)
}

func new(c *Config, s *Sqlite) (*Sqlite, error) {
	var err error

	s.Client, err = gorm.Open(sqlite.Open(c.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(c.logLevel())),
	})

	return s, err
}

func (s *Sqlite) Close() error {
	if s.Client == nil {
		return nil
	}

	db, err := s.Client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
