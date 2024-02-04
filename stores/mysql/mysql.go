package mysql

import (
	"sync"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const name = "mysql"

var once sync.Once

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
		return errors.Internal("mysql", "Mysql type assertion failed, expected *Config, got %T", cfg)
	}

	rdb, err := New(mysqlCfg)
	if err != nil {
		return err
	}
	m.Client = rdb.Client

	return nil
}

func (m *Mysql) Name() string {
	return name
}

func New(c *Config, opts ...Option) (*Mysql, error) {
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	m := &Mysql{}

	for _, opt := range opts {
		opt(m)
	}

	m, err := new(c, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func new(c *Config, m *Mysql) (*Mysql, error) {
	var err error
	
	once.Do(func() {
		m.Client, err = gorm.Open(mysql.Open(c.dsn()), &gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(c.logLevel())),
		})
		if err != nil {
			return
		}

		sqlDB, err := m.Client.DB()
		if err != nil {
			return
		}

		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	})

	return m, err
}

func (m *Mysql) Close() error {
	sqlDB, err := m.Client.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
