package redis

import (
	"context"
	"runtime"
	"sync"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/redis/go-redis/v9"
)

const name = "redis"

var once sync.Once

var (
	cfgByIoc   any
	cfgByIocOk bool
)

type Redis struct {
	Client *redis.Client
}

type Option func(*Redis)

func init() {
	cfgByIoc, cfgByIocOk = ioc.Container.GetFieldValueByConfig("Redis")
	if cfgByIocOk {
		ioc.Container.Registry(ioc.DbNamespace, &Redis{})
	}
}

func (r *Redis) Init() error {
	if !cfgByIocOk {
		return nil
	}

	redisCfg, ok := cfgByIoc.(*Config)
	if !ok {
		return errors.Internal("redis", "Redis type assertion failed, expected *Config, got %T", cfgByIoc)
	}

	rdb, err := New(redisCfg)
	if err != nil {
		return err
	}
	r.Client = rdb.Client

	return nil
}

func (r *Redis) Name() string {
	return name
}

func New(c *Config, opts ...Option) (*Redis, error) {
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	if c.PoolSize == 0 {
		c.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	r := &Redis{}

	for _, opt := range opts {
		opt(r)
	}

	r, err := new(c, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func new(c *Config, r *Redis) (*Redis, error) {
	var err error
	once.Do(func() {
		r.Client = redis.NewClient(&redis.Options{
			Addr:     c.address(),
			Password: c.Password,
			DB:       c.DB,
			Protocol: c.Protocol,
			PoolSize: c.PoolSize,
		})

		_, err = r.Client.Ping(context.Background()).Result()
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
