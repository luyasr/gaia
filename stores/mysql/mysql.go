package mysql

import (
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var once sync.Once

type Mysql struct {
	Client *gorm.DB
}

type Option func(*Mysql)

func NewMysql(c Config, opts ...Option) (*Mysql, error) {
	err := c.initConfig()
	if err != nil {
		return nil, err
	}

	m, err := newMysql(c)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(m)
	}

	return m, nil
}

func newMysql(c Config) (*Mysql, error) {
	m := &Mysql{}

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
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mysql) Close() error {
	sqlDB, err := m.Client.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
