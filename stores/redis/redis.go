package redis

import (
	"context"
	"runtime"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"github.com/redis/go-redis/v9"
)

const Name = "redis"

type Redis struct {
	Client *redis.Client
}

type Option func(*Redis)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Redis{})
}

func (r *Redis) Init() error {
	cfg, ok := ioc.Container.GetFieldValueByConfig("Redis")
	if !ok {
		return nil
	}

	redisCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("redis config type assertion failed", "expected *Config, got %T", cfg)
	}

	rdb, err := New(redisCfg)
	if err != nil {
		return err
	}
	r.Client = rdb.Client

	return nil
}

func (r *Redis) Name() string {
	return Name
}

func New(c *Config, opts ...Option) (*Redis, error) {
	cfg, err := c.initConfig()
	if err != nil {
		return nil, err
	}

	if cfg.PoolSize == 0 {
		cfg.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	r := &Redis{}

	for _, opt := range opts {
		opt(r)
	}

	return new(cfg, r)
}

func new(c *Config, r *Redis) (*Redis, error) {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     c.address(),
		Password: c.Password,
		DB:       c.DB,
		Protocol: c.Protocol,
		PoolSize: c.PoolSize,
	})

	_, err := r.Client.Ping(context.Background()).Result()

	return r, err
}

func (r *Redis) Close() error {
	if r.Client == nil {
		return nil
	}

	return r.Client.Close()
}
