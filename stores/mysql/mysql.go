package mysql

import (
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Name = "mysql"

type Mysql struct {
	Client *gorm.DB
}

type Option func(*Mysql)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Mysql{})
}

func (m *Mysql) Init() error {
	cfg, ok := ioc.Container.GetFieldValueByConfig("Mysql")
	if !ok {
		return nil
	}

	mysqlCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("mysql config type assertion failed", "expected *Config, got %T", cfg)
	}

	rdb, err := New(mysqlCfg)
	if err != nil {
		return err
	}
	m.Client = rdb.Client

	return nil
}

func (m *Mysql) Name() string {
	return Name
}

func New(c *Config, opts ...Option) (*Mysql, error) {
	cfg, err := c.initConfig()
	if err != nil {
		return nil, err
	}

	m := &Mysql{}

	for _, opt := range opts {
		opt(m)
	}

	return new(cfg, m)
}

func new(c *Config, m *Mysql) (*Mysql, error) {
	var err error

	m.Client, err = gorm.Open(mysql.Open(c.dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(c.logLevel())),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := m.Client.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))

	return m, nil
}

func (m *Mysql) Close() error {
	if m.Client == nil {
		return nil
	}
	sqlDB, err := m.Client.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
